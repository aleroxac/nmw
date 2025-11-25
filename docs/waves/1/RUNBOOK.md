# Runbook

## Worker - Python
``` shell
mkdir worker
cd worker

python -m virtualenv .venv
source .venv/bin/activate
pip install fastapi python-multipart uvicorn faster-whisper
pip freeze > requirements.txt

touch main.py

sudo systemctl start ollama.service
ollama pull mistral:7b-instruct
nohup ollama serve &

uvicorn main:app --reload

curl -s -X POST \
  -F "file=@../assets/audio.mp3" \
  http://localhost:8000/transcribe | jq
```

## Backend - Golang
``` shell
mkdir backend
cd backend

go mod init github.com/aleroxac/nmw/backend
touch main.go
go mod tidy

curl -s -X POST -F "file=@../assets/audio.mp3" http://localhost:8080/upload
```

## Frontend - Javascript
``` shell
touch index.html
touch style.css
touch script.js
```
