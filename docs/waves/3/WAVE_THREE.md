# 🌊 NMW Roadmap: Wave 3 (Deep Dive & Quality)

**Status:** Bloqueado (Aguardando Wave 2)
**Foco:** Linguística Avançada, Testes Automatizados, Qualidade de Produção.

## 🎯 Objetivo Principal
Refinar a "inteligência" pedagógica. Sair do básico "está certo/errado" para análises profundas de "naturalidade" e "sotaque". Garantir que o sistema seja robusto e testável.

## 🏗️ Mudanças na Arquitetura
* A arquitetura infraestrutural mantém-se a da Wave 2.
* O foco muda para a **Lógica de Domínio** e **Pipeline de Qualidade**.

## 🛠️ Tech Stack Adicional
* **NLP:** `spacy` (Python), Modelos de Fonética (MFA ou Wav2Vec fine-tuned).
* **Testes:** `testcontainers-go`, `cypress` (ou Playwright).
* **Observabilidade:** OpenTelemetry (Tracing Distribuído).

## 📋 Tarefas Detalhadas

### 1. Análise Linguística (Advanced NLP)
- [ ] **Lexical Diversity:** Calcular se o usuário tem vocabulário pobre ou rico (Type-Token Ratio).
- [ ] **Disfluency Detection:** Detectar e marcar "uhh", "umm", gaguejadas.
- [ ] **Word Stress:** Comparar a sílaba tônica do usuário com a do dicionário (usando CMU Dict).

### 2. Testes e Qualidade (QA/SRE)
- [ ] **Testes de Integração:** Usar `testcontainers` no Go para subir um Postgres real e testar o fluxo de repositório.
- [ ] **Testes E2E:** Scriptar o fluxo completo do usuário no navegador.
- [ ] **Tracing:** Implementar OpenTelemetry para ver quanto tempo o áudio ficou na fila vs processando.

### 3. Features "Pro"
- [ ] **Shadowing Mode:** UI para tocar áudio nativo -> user grava -> UI sobrepõe as ondas sonoras.
- [ ] **Feedback Híbrido:** Permitir escolher entre modelos (OpenAI para precisão vs Local para privacidade) via configurações do usuário.

## ⚠️ Definição de "Pronto" (DoD)
O sistema fornece um relatório detalhado de pronúncia e estilo. Existem testes rodando no CI (GitHub Actions). Logs mostram o trace completo da requisição.