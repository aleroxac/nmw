from fastapi import FastAPI, UploadFile
from faster_whisper import WhisperModel
import shutil
import os
import requests
import json

app = FastAPI()

# Carrega o modelo na RAM apenas uma vez (use "tiny" para ser rápido no teste)
model = WhisperModel("tiny", device="cpu", compute_type="int8")

# model = WhisperModel("tiny", device="cuda", compute_type="int8")
# Unable to load any of {libcudnn_ops.so.9.1.0, libcudnn_ops.so.9.1, libcudnn_ops.so.9, libcudnn_ops.so} 
# Invalid handle. Cannot load symbol cudnnCreateTensorDescriptor


@app.post("/transcribe")
def transcribe(file: UploadFile):
    # 1. transcribe audio
    temp_filename = f"temp_{file.filename}"
    with open(temp_filename, "wb") as buffer:
        shutil.copyfileobj(file.file, buffer)

    segments, info = model.transcribe(temp_filename, beam_size=5)
    transcription = " ".join([segment.text for segment in segments])
    os.remove(temp_filename)

    # 2. get corrections
    ollama_url = "http://localhost:11434/api/generate"
    corrections_prompt = f"""
Você é um revisor gramatical. Responda SOMENTE com JSON válido com os valores preenchidos com base na sua análise gramatical, sem markdown, rótulos, comentários ou texto extra.
Inicie a resposta com {{ e termine com }}. Use exatamente este schema como base para depois popular com os resultados da análise:
{{
  "original": "{transcription}",
  "reviewed": "",
  "grammar_errors": [],
  "suggestions": []
}}
Regras:
- "reviewed": texto revisado completo.
- "grammar_errors" e "suggestions": listas de strings; se não houver, deixe vazias.
- Não repita instruções, não inclua explicações fora do JSON.

Texto para revisar:
{transcription}
""".strip()

    payload = {
        "model": "mistral:7b-instruct",
        "prompt": corrections_prompt,
        "stream": False,
        "format": "json"
    }
    response = requests.post(ollama_url, json=payload)
    print(f"\n\n\n==response==\n{response}\n\n\n")
    transcription_with_corrections = json.loads(response.json()["response"])

    return {
        "language": info.language,
        "probability": info.language_probability,
        "result": transcription_with_corrections
    }
