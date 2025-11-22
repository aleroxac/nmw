# nmw - Not Merely Words
A language learning platform to where is possible to learn "Not Merely Words".

## Features
- [x] Speech to Text transcription with AI analysis

## Components
- frontend (Typescript): Provide an UI, allowing the user record an audio and than get a transcription with grammar analysis.
- backend (Golang): Orchestrate the communication between the frontend and the backend
- worker-ai (Python): Execute a Speech to Text(STT) transcription and call an AI model to provide a grammar analysis based on an audio file.

## Tech Stack
- frontend: Next.js, Tailwindcss, Shadcn
- worker-ai: FastAPI, Uvicorn, faster-whisper
- backend: gin-gonic
- infra: PostgreSQL
