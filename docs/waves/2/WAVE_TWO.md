# 🌊 NMW Roadmap: Wave 2 (Scalability & DSP)

**Status:** Bloqueado (Aguardando Wave 1)
**Foco:** Assincronismo, Docker Compose, Análise de Sinal (DSP).

## 🎯 Objetivo Principal
Resolver o problema de latência da Wave 1 e adicionar análises que não dependem de texto (ritmo, entonação). Transformar a arquitetura para suportar escala usando filas e containers orquestrados.

## 🏗️ Mudanças na Arquitetura
* **De:** HTTP Síncrono (Bloqueante).
* **Para:** Event-Driven (Assíncrono).
* **Novos Componentes:** RabbitMQ (Broker), Redis (Cache/PubSub), WebSockets (Real-time feedback).

## 🛠️ Tech Stack Adicional
* **Infra:** Docker Compose completo.
* **Messaging:** RabbitMQ (AMQP).
* **Cache:** Redis.
* **DSP Libs:** `librosa` (Python), `parselmouth` (Python Wrapper pro Praat).
* **Frontend:** `socket.io-client` ou WS nativo, `recharts` (para gráficos).

## 📋 Tarefas Detalhadas

### 1. Dockerização e Infra (DevOps)
- [ ] Criar `docker-compose.yml` unificando todos os serviços.
- [ ] Configurar volumes para persistência de dados (Postgres/RabbitMQ).

### 2. Refatoração Backend (Golang)
- [ ] Implementar `RabbitMQ Producer`: Em vez de chamar o Python, publica na fila `audio_processing_queue`.
- [ ] Implementar `WebSocket Hub`: Para manter conexão aberta com o Frontend e enviar updates de progresso ("Transcrevendo...", "Analisando...").
- [ ] Implementar `Redis Cache`: Se o mesmo hash de áudio for enviado, retornar resultado cacheado.

### 3. Evolução do Worker (Python)
- [ ] Transformar API (FastAPI) em Worker Consumidor (loop infinito ouvindo RabbitMQ).
- [ ] **Feature Nova:** Análise de Pitch (Entonação) usando `parselmouth`.
- [ ] **Feature Nova:** Análise de Pausas e WPM (Words Per Minute) usando `librosa`.

### 4. Frontend (Dashboard)
- [ ] Criar gráficos visuais: Linha de entonação, Barra de velocidade.
- [ ] Receber updates via WebSocket em vez de esperar a resposta HTTP única.

## ⚠️ Definição de "Pronto" (DoD)
O usuário grava um áudio longo (1 min). A UI mostra passos de progresso. Os gráficos de ritmo e entonação aparecem antes da correção gramatical (pois são mais rápidos de calcular). Tudo roda via `docker-compose up`.