package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/upload", func(c *gin.Context) {
		// ---- 1. Read the uploaded file ----
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer file.Close()

		// ---- 2. Prepare multipart request ----
		var b bytes.Buffer
		writer := multipart.NewWriter(&b)

		// Add the file to the multipart writer
		part, err := writer.CreateFormFile("file", fileHeader.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		_, err = io.Copy(part, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Close writer to finalize the multipart body
		writer.Close()

		// ---- 3. Do the request ----
		req, err := http.NewRequest("POST", "http://localhost:8000/transcribe", &b)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		req.Header.Set("Content-Type", writer.FormDataContentType())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer resp.Body.Close()

		// ---- 4. Read response ----
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ---- 5. Return the JSON response ----
		c.Data(resp.StatusCode, "application/json", bodyBytes)
		fmt.Println(string(bodyBytes))
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}
