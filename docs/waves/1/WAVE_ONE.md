# 🌊 NMW Roadmap: Wave 1 (The Foundation)

**Status:** Planejado
**Foco:** MVP Funcional, Estrutura de Código, Fluxo Síncrono.

## 🎯 Objetivo Principal
Criar o "esqueleto" da aplicação (Walking Skeleton). O usuário grava um áudio, o sistema processa (transcreve + corrige gramática básica) e devolve o resultado. O foco é fazer as peças se conversarem, não em performance.

## 🏗️ Arquitetura da Wave
* **Tipo:** Monólito Modular Distribuído (comunicação direta HTTP).
* **Fluxo:** `Frontend` -> (HTTP POST) -> `Backend Go` -> (HTTP POST) -> `Worker Python` -> (Response JSON) -> `Backend Go` -> (Response JSON) -> `Frontend`.

## 🛠️ Stack Tecnológica
* **Frontend:** Next.js 14 (App Router), TypeScript, TailwindCSS, Shadcn/ui.
* **Backend:** Golang (Framework `Chi` ou `Fiber`), Postgres (Driver `pgx/v5`).
* **AI/Worker:** Python 3.10+, FastAPI, Faster-Whisper (STT), Ollama (LLM local).
* **Infra Local:** Docker (básico) ou rodando processos no terminal.

## 📋 Tarefas Detalhadas

### 1. Frontend (Interface Básica)
- [ ] Setup do Next.js com TypeScript e Shadcn.
- [ ] Criar componente `AudioRecorder` (usando Web Audio API básica).
- [ ] Criar tela de "Loading" enquanto espera a resposta do backend.
- [ ] Exibir resultado simples: Texto Transcrito + Texto Corrigido.

### 2. Backend (Golang Core)
- [ ] Criar estrutura de pastas Clean/Hexagonal (`cmd`, `internal`, `pkg`).
- [ ] Implementar `POST /upload`: Recebe `multipart/form-data`.
- [ ] Implementar **Interface** `AIService`:
    * *Nota:* Isso é crucial para a Wave 2. Defina a interface agora para facilitar a mudança depois.
- [ ] Persistência: Salvar metadados da tentativa no Postgres.

### 3. AI Worker (Python Brain)
- [x] Setup FastAPI.
- [x] Endpoint `/transcript`: Recebe áudio.
- [x] STT: Implementar `faster-whisper` para transcrever.
- [x] LLM: Conectar no Ollama local e pedir correção gramatical simples.

### 4. Integração
- [ ] Fazer o "Hello World" do áudio percorrer todo o caminho e voltar como texto.

## ⚠️ Definição de "Pronto" (DoD)
Consigo subir o projeto com um comando (ou scripts), gravar "I has a car", e receber de volta "I have a car" na tela em menos de 10 segundos.
