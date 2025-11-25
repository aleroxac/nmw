# nmw - Not Merely Words
A language learning platform where is possible to learn "Not Merely Words".

## Features
- [x] Speech to Text transcription with AI analysis

## Components
- frontend (`Typescript`): Provide an UI, allowing the user record an audio and then get a transcription with grammar analysis.
- backend (`Golang`): Orchestrate the communication between the frontend and the backend
- worker (`Python`): Execute a Speech to Text(STT) transcription and call an AI model to provide a grammar analysis based on an audio file.

## Tech Stack
- frontend: Next.js, Tailwindcss, Shadcn
- backend: gin-gonic
- worker: FastAPI, Uvicorn, faster-whisper, Ollama
- infra: PostgreSQL
