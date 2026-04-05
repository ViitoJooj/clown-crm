# 🤡 Clown CRM - Docs

[English](./README.en.md) | **Português**

## Setup Rápido

```bash
git clone https://github.com/ViitoJooj/clown-crm.git
cd clown-crm
docker-compose up -d
go run cmd/api/main.go
```

Outro terminal:
```bash
cd view
dx serve
```

Pronto. Acessa `localhost:8080`.

## O que tem

- **Contatos**: CRUD, busca, tags, campos custom
- **Deals**: Pipeline drag-and-drop, conversão, métricas
- **Tasks**: Criar, atribuir, deadline, notificações
- **Chat**: WebSocket real-time, histórico
- **Calls**: Log de chamadas, Twilio opcional

## Stack

- Go 1.25 + Gin (backend)
- Dioxus 0.7 (frontend Rust)
- PostgreSQL
- JWT auth
- WebSocket

## Configuração

`.env` na raiz:

```env
DATABASE_URL=postgres://user:pass@localhost/clowncrm
JWT_SECRET=qualquer-coisa-longa-aqui
JWT_REFRESH_SECRET=outra-coisa-longa

# Twilio (opcional)
TWILIO_ACCOUNT_SID=ACxxx
TWILIO_AUTH_TOKEN=xxx
```

Migrations:

```bash
# Com migrate
migrate -path sql -database $DATABASE_URL up

# Manual
for f in sql/*.sql; do psql $DATABASE_URL < $f; done
```

## API Endpoints

### Auth

`POST /api/auth/register`
```json
{"email": "user@mail.com", "password": "pass", "name": "Nome"}
```

`POST /api/auth/login`
```json
{"email": "user@mail.com", "password": "pass"}
```

Retorna tokens JWT.

### Contatos

- `GET /api/contacts` - Lista (query: page, limit, search)
- `GET /api/contacts/:id` - Um contato
- `POST /api/contacts` - Criar
- `PUT /api/contacts/:id` - Editar
- `DELETE /api/contacts/:id` - Deletar

### Deals

- `GET /api/deals` - Lista
- `GET /api/deals/:id` - Detalhes
- `POST /api/deals` - Criar
- `PUT /api/deals/:id` - Editar
- `PUT /api/deals/:id/stage` - Mover de etapa
- `PUT /api/deals/:id/won` - Marcar como ganha
- `PUT /api/deals/:id/lost` - Marcar como perdida

### Tasks

- `GET /api/tasks` - Lista (filtros: status, assigned_to, overdue)
- `POST /api/tasks` - Criar
- `PUT /api/tasks/:id/complete` - Concluir
- `DELETE /api/tasks/:id` - Deletar

### WebSocket Chat

`WS /ws/chat?token=JWT_TOKEN`

Mensagem:
```json
{
  "from": "uuid",
  "to": "uuid",
  "message": "texto",
  "timestamp": "ISO8601"
}
```

## Estrutura

```
clown-crm/
├── cmd/api/          # main.go
├── internal/
│   ├── domain/       # Entities (Contact, Deal, Task...)
│   ├── repository/   # DB layer
│   ├── services/     # Business logic
│   └── http/         # Controllers
├── sql/              # Migrations (001_*.sql)
├── view/             # Frontend Dioxus
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── styles/
└── pkg/              # Utils (twilio, websocket...)
```

## Problemas

### Postgres não conecta

```bash
docker ps  # Vê se tá rodando
docker-compose logs db  # Vê o erro
docker-compose restart db
```

### Frontend não compila (webkit2gtk)

Ubuntu/Debian:
```bash
sudo apt install libwebkit2gtk-4.1-dev
```

Fedora:
```bash
sudo dnf install webkit2gtk4.1-devel
```

Arch:
```bash
sudo pacman -S webkit2gtk-4.1
```

### "JWT token invalid"

Confere se o JWT_SECRET no `.env` tá igual ao que o código tá usando.

### Migrations deram erro

Reseta tudo:
```bash
dropdb clowncrm && createdb clowncrm
ls sql/*.sql | sort | xargs -I{} psql clowncrm < {}
```

## Contribuindo

1. Fork
2. Branch (`git checkout -b fix/alguma-coisa`)
3. Commit (`git commit -m 'fix: corrige alguma coisa'`)
4. Push (`git push origin fix/alguma-coisa`)
5. PR

Usa conventional commits:
- `feat:` feature nova
- `fix:` bug fix
- `docs:` docs
- `refactor:` refactor

Testes são bem-vindos mas não obrigatórios pra PR pequeno.

## Roadmap (talvez)

- [ ] Export CSV/Excel
- [ ] Email integration (Gmail/Outlook)
- [ ] Mobile app (React Native ou Flutter)
- [ ] Analytics dashboard
- [ ] WhatsApp integration
- [ ] Custom fields UI
- [ ] Temas/dark mode
- [ ] API rate limiting

## Licença

MIT. Faz o que quiser.

---

Feito com ☕ por [@ViitoJooj](https://github.com/ViitoJooj)
