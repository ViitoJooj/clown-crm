# 🤡 Clown CRM

CRM simples e direto. Sem firulas, sem mensalidade por usuário, sem enrolação.

[English](#english) | [Português](#português)

---

## Português

### O que é isso?

Um CRM que você pode hospedar onde quiser. Feito com Go (backend rápido) e Dioxus/Rust (frontend que não trava).

Gerencia contatos, negociações, tarefas e chamadas. Tem chat em tempo real e se você quiser pode integrar com Twilio pra logar ligações automaticamente.

**Não tem:**
- Cobrança por usuário
- Bloqueio de features na "versão grátis"
- Dashboard confuso com 50 métricas que ninguém usa

**Tem:**
- Código aberto (mexe como quiser)
- Pipeline de vendas drag-and-drop
- Chat interno sem delay
- Auth com JWT (seguro e rápido)

### Rodando local

```bash
git clone https://github.com/ViitoJooj/clown-crm.git
cd clown-crm

# Suba o postgres
docker-compose up -d

# Backend
go run cmd/api/main.go

# Frontend (outro terminal)
cd view
dx serve
```

Acessa em `http://localhost:8080` e já era.

### Stack

- **Backend:** Go 1.25 + Gin (rápido e simples)
- **Frontend:** Dioxus 0.7 (Rust, zero JavaScript se quiser)
- **DB:** PostgreSQL (confiável)
- **Auth:** JWT com refresh token
- **Real-time:** WebSocket (chat + notificações)

### Estrutura

```
clown-crm/
├── cmd/api/          # Entry point
├── internal/
│   ├── domain/       # Modelos e validação
│   ├── http/         # Controllers
│   ├── repository/   # Acesso ao banco
│   └── services/     # Lógica de negócio
├── sql/              # Migrations
├── view/             # Frontend Dioxus
└── pkg/              # Utils
```

---

## English

### What's this?

Self-hosted CRM. No monthly fees, no user limits, no BS.

Built with Go backend and Dioxus/Rust frontend. Handles contacts, deals, tasks, and calls. Has real-time chat and optional Twilio integration.

**Doesn't have:**
- Per-user pricing
- Feature locks
- Overcomplicated dashboards

**Has:**
- Open source (modify as needed)
- Drag-and-drop pipeline
- Internal chat with WebSocket
- Secure JWT auth

### Running locally

Same commands as above. Works on Linux/Mac/Windows.

### Contributing

PRs welcome. Keep it simple, keep it fast.

### License

MIT - do whatever you want with it.

---

Made by [@ViitoJooj](https://github.com/ViitoJooj)
