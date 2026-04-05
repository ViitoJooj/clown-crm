# 🤡 Clown CRM - Docs

**English** | [Português](./README.pt-BR.md)

## Quick Setup

```bash
git clone https://github.com/ViitoJooj/clown-crm.git
cd clown-crm
docker-compose up -d
go run cmd/api/main.go
```

Another terminal:
```bash
cd view
dx serve
```

Done. Go to `localhost:8080`.

## Features

- **Contacts**: CRUD, search, tags, custom fields
- **Deals**: Drag-and-drop pipeline, conversion metrics
- **Tasks**: Create, assign, deadlines, notifications
- **Chat**: Real-time WebSocket, history
- **Calls**: Call logging, optional Twilio integration

## Stack

- Go 1.25 + Gin (backend)
- Dioxus 0.7 (Rust frontend)
- PostgreSQL
- JWT auth
- WebSocket

## Configuration

`.env` in root:

```env
DATABASE_URL=postgres://user:pass@localhost/clowncrm
JWT_SECRET=something-long-and-random
JWT_REFRESH_SECRET=another-long-secret

# Twilio (optional)
TWILIO_ACCOUNT_SID=ACxxx
TWILIO_AUTH_TOKEN=xxx
```

Run migrations:

```bash
# With migrate CLI
migrate -path sql -database $DATABASE_URL up

# Manual
for f in sql/*.sql; do psql $DATABASE_URL < $f; done
```

## API Endpoints

### Auth

`POST /api/auth/register`
```json
{"email": "user@mail.com", "password": "pass", "name": "Name"}
```

`POST /api/auth/login`
```json
{"email": "user@mail.com", "password": "pass"}
```

Returns JWT tokens.

### Contacts

- `GET /api/contacts` - List (query: page, limit, search)
- `GET /api/contacts/:id` - Single contact
- `POST /api/contacts` - Create
- `PUT /api/contacts/:id` - Update
- `DELETE /api/contacts/:id` - Delete

### Deals

- `GET /api/deals` - List
- `GET /api/deals/:id` - Details
- `POST /api/deals` - Create
- `PUT /api/deals/:id` - Update
- `PUT /api/deals/:id/stage` - Move stage
- `PUT /api/deals/:id/won` - Mark as won
- `PUT /api/deals/:id/lost` - Mark as lost

### Tasks

- `GET /api/tasks` - List (filters: status, assigned_to, overdue)
- `POST /api/tasks` - Create
- `PUT /api/tasks/:id/complete` - Complete
- `DELETE /api/tasks/:id` - Delete

### WebSocket Chat

`WS /ws/chat?token=JWT_TOKEN`

Message format:
```json
{
  "from": "uuid",
  "to": "uuid",
  "message": "text",
  "timestamp": "ISO8601"
}
```

## Project Structure

```
clown-crm/
├── cmd/api/          # main.go
├── internal/
│   ├── domain/       # Entities (Contact, Deal, Task...)
│   ├── repository/   # DB layer
│   ├── services/     # Business logic
│   └── http/         # Controllers
├── sql/              # Migrations (001_*.sql)
├── view/             # Dioxus frontend
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── styles/
└── pkg/              # Utils (twilio, websocket...)
```

## Troubleshooting

### Postgres won't connect

```bash
docker ps  # Check if running
docker-compose logs db  # See errors
docker-compose restart db
```

### Frontend won't compile (webkit2gtk)

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

Make sure JWT_SECRET in `.env` matches what the code uses.

### Migration errors

Reset everything:
```bash
dropdb clowncrm && createdb clowncrm
ls sql/*.sql | sort | xargs -I{} psql clowncrm < {}
```

## Contributing

1. Fork
2. Branch (`git checkout -b fix/something`)
3. Commit (`git commit -m 'fix: fixes something'`)
4. Push (`git push origin fix/something`)
5. PR

Use conventional commits:
- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation
- `refactor:` refactoring

Tests are welcome but not required for small PRs.

## Roadmap (maybe)

- [ ] CSV/Excel export
- [ ] Email integration (Gmail/Outlook)
- [ ] Mobile app
- [ ] Analytics dashboard
- [ ] WhatsApp integration
- [ ] Custom fields UI
- [ ] Themes/dark mode
- [ ] API rate limiting

## License

MIT. Do whatever you want.

---

Made with ☕ by [@ViitoJooj](https://github.com/ViitoJooj)
