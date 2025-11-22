# TODOs

## worker
- [ ] use pathlib instead of os
- [ ] review the need of a temporary file during the transcription
- [ ] implent a class for input and output
- [ ] look for the faster_whisper documentation to see which options are available
- [ ] integrate faster_whisper with ollama
- [ ] consume the model params via env vars or dotenv

## backend
- [ ] criar endpoint /api/analyze, que será chamada pelo frontend passando arquivo .wav, e então o backend deverá chamar o serviço de AI, obter a transcrição, guardar os dados no banco de dados, e retornar a transcrição e os dados obtidos pela AI de volta para o frontend
    - receber chamada do front
    - chamar o worker
    - salvar o arquivo de audio localmente
    - guardar dados no banco
        - attempts
            - audio_file_path: TEXT
            - transcription: TEXT
            - ai_analysis: JSONB
    - retornar dados para o frontend
