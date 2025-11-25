# 🌊 NMW Roadmap: Wave 1 (The Foundation)

**Status:** Planejado
**Foco:** MVP Funcional, Estrutura de Código, Fluxo Síncrono.

## 🎯 Objetivo Principal
Criar o "esqueleto" da aplicação (Walking Skeleton). O usuário grava um áudio, o sistema processa (transcreve + corrige gramática básica) e devolve o resultado. O foco é fazer as peças se conversarem, não em performance.

## 🏗️ Arquitetura da Wave
* **Tipo:** Monólito Modular Distribuído (comunicação direta HTTP).
* **Fluxo:** `Frontend` -> (HTTP POST) -> `Backend Go` -> (HTTP POST) -> `Worker Python` -> (Response JSON) -> `Backend Go` -> (Response JSON) -> `Frontend`.

## 🛠️ Stack Tecnológica
* **Frontend:** HTML, CSS, Javascript
* **Backend:** Golang (Framework `Chi`)
* **AI/Worker:** Python, FastAPI, Faster-Whisper (STT), Ollama (LLM local).

## 📋 Tarefas Detalhadas

### 1. Frontend (Interface Básica)
- [x] Criar página HTML com botão de start/stop de gravação, botão de análise, link para baixar audio, cards para transcrição, erros gramaticais e sugestões.
- [x] Criar arquivo CSS com estilização da página
- [x] Criar arquivo de script Javascript para gerar o arquivo de áudio, chamar o backend, fazer parse do response e mostrar os dados corretamente na tela.

### 2. Backend (Golang Core)
- [x] Criar endpoint `/upload`
    - [x] Recebe `POST` em `/upload`, com arquivo de audio usando `multipart/form-data`
    - [x] Envia `POST` em `/transcribe` do worker, com arquivo de audio usando `multipart/form-data`

### 3. AI Worker (Python Brain)
- [x] Setup FastAPI.
- [x] Endpoint `/transcript`: Recebe áudio.
- [x] STT: Implementar `faster-whisper` para transcrever.
- [x] LLM: Conectar no Ollama local e pedir correção gramatical simples.

### 4. Integração
- [ ] Fazer o "Hello World" do áudio percorrer todo o caminho e voltar como texto.

## ⚠️ Definição de "Pronto" (DoD)
Consigo subir o projeto com um comando (ou scripts), gravar "I has a car", e receber de volta "I have a car" na tela em menos de 10 segundos.
