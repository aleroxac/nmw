let mediaRecorder;
let audioChunks = [];
let audioBlob = null;

const recordBtn = document.getElementById("recordBtn");
const analyzeBtn = document.getElementById("analyzeBtn");
const audioPlayer = document.getElementById("audioPlayer");
const downloadLink = document.getElementById("downloadLink");

const errorBox = document.getElementById("errorBox");

const resultOriginal = document.getElementById("resultOriginal");
const resultErrors = document.getElementById("resultErrors");
const resultSuggestions = document.getElementById("resultSuggestions");

const originalText = document.getElementById("originalText");
const grammarErrors = document.getElementById("grammarErrors");
const suggestions = document.getElementById("suggestions");

/* ----------------------------------------------------
   START / STOP RECORDING
---------------------------------------------------- */
recordBtn.addEventListener("click", async () => {
    errorBox.style.display = "none";

    if (!mediaRecorder || mediaRecorder.state === "inactive") {
        audioChunks = [];
        const stream = await navigator.mediaDevices.getUserMedia({ audio: true });

        mediaRecorder = new MediaRecorder(stream);
        mediaRecorder.ondataavailable = e => audioChunks.push(e.data);

        mediaRecorder.onstop = () => {
            audioBlob = new Blob(audioChunks, { type: "audio/wav" });
            const url = URL.createObjectURL(audioBlob);

            audioPlayer.src = url;
            downloadLink.href = url;

            analyzeBtn.disabled = false;
        };

        mediaRecorder.start();

        // UI: recording mode
        recordBtn.textContent = "⏹️ Stop Recording";
        recordBtn.classList.add("btn-recording");
        analyzeBtn.disabled = true;

    } else {
        mediaRecorder.stop();

        // UI: back to normal
        recordBtn.textContent = "🎙️ Start Recording";
        recordBtn.classList.remove("btn-recording");
    }
});

/* ----------------------------------------------------
   ANALYZE AUDIO
---------------------------------------------------- */
analyzeBtn.addEventListener("click", async () => {
    console.log("DEBUG: analyze button clicked");
    if (!audioBlob) {
        alert("You need to record audio first!");
        return;
    }

    const formData = new FormData();
    formData.append("file", audioBlob, "audio.wav");

    // Reset results
    resultOriginal.style.display = "none";
    resultErrors.style.display = "none";
    resultSuggestions.style.display = "none";

    // UI: analyzing mode
    analyzeBtn.disabled = true;
    analyzeBtn.classList.add("btn-loading");
    analyzeBtn.textContent = "Analyzing...";

    try {
        const resp = await fetch("http://localhost:8080/upload", {
            method: "POST",
            body: formData
        });

        const text = await resp.text();
        let json;

        try {
            json = JSON.parse(text);
        } catch (e) {
            throw new Error("Invalid JSON: " + text);
        }

        const original = json.result.original;
        const errors = json.result.grammar_errors || [];
        const suggs = json.result.suggestions || [];

        originalText.textContent = original;
        resultOriginal.style.display = "block";

        grammarErrors.innerHTML = errors.length
            ? errors.map(e => `• ${e}`).join("<br>")
            : "<i>No errors found</i>";
        resultErrors.style.display = "block";

        suggestions.innerHTML = suggs.length
            ? suggs.map(s => `• ${s}`).join("<br>")
            : "<i>No suggestions</i>";
        resultSuggestions.style.display = "block";

    } catch (err) {
        errorBox.textContent = err;
        errorBox.style.display = "block";
    }

    // UI: restore state
    analyzeBtn.classList.remove("btn-loading");
    analyzeBtn.textContent = "🔍 Analyze Audio";
    analyzeBtn.disabled = false;
});
