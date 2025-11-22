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

uvicorn main:app --reload

curl -X POST \
  -F "file=@audio.mp3" \
  http://localhost:8000/transcribe
```

## Backend - Golang
``` shell
mkdir backend
cd backend

go mod init github.com/aleroxac/nmw/backend
```
