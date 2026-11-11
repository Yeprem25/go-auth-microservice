# Go Auth Microservice — Usman plan

Npatak: sovorel **Go, Gin, PostgreSQL, Docker, gRPC, JWT** — kaxucelov irakan auth microservice, vor@ GitHub-um kerevi vorpes portfolio.

Kanon: **kod@ grum es du**, es bacatrum em, stugum em, hushum em. Erb khrvas — grir indz ays fayli Phase/Task hamar@ ev sxal@.

---

## Verjnakan arkhitektura (inch enq karucum)

```
            HTTP/JSON                         gRPC
 Client ─────────────▶ MainBackend ─────────────────▶ GoAuthMicroservice
 (curl/Postman)        (Gin, :8080)                   (gRPC, :50051)
                            │                               │
                            ▼                               ▼
                       Postgres (notes)              Postgres (users, refresh_tokens)
```

Amen microservice — **arandzin GitHub repo**, bolor@ `Go Protfolio/` papkayum:

- **GoAuthMicroservice** (`go-auth-microservice`) — **miayn auth**, miayn gRPC (HTTP chuni): Register, Login (bcrypt), Refresh, Logout, ValidateToken. Ir sepakan DB-n.
- **MainBackend** (`go-main-backend`) — Gin REST API client-i hamar. Auth request-ner@ uxarkum e gRPC-ov auth-service-in, ir middleware-@ token@ stugum e `ValidateToken`-ov. Uni ir biznes feature-@ (Notes) ev ir DB-n.
- **Database per service** — amen service iren DB-n uni, irar DB chen kardum.
- Amen inch@ ashkhatum e **docker compose up**-ov.

## Verjnakan papkayi karucvack

```
Go Protfolio/
├── GoAuthMicroservice/          # repo: go-auth-microservice
│   ├── proto/auth/v1/auth.proto # "paymanagir@" — inchpes khosel ays service-i het
│   ├── gen/auth/v1/             # generacvac kod (commit-vum e! MainBackend-@ import e anum)
│   ├── cmd/auth-service/main.go
│   ├── internal/{config,database,model,repository,service,grpcserver}/
│   ├── migrations/
│   └── Dockerfile, docker-compose.yml, Makefile, README.md
│
└── MainBackend/                 # repo: go-main-backend
    ├── cmd/main-backend/main.go
    ├── internal/{handler,middleware,authclient,notes}/
    ├── migrations/
    └── Dockerfile, docker-compose.yml (bolor service-ner@), Makefile, README.md
```

**Inchpes e MainBackend-@ stanum gRPC kod@:** `import "github.com/Yeprem25/go-auth-microservice/gen/auth/v1"` — sovorakan Go dependency. Local zargacman hamar `go.mod`-um: `replace github.com/Yeprem25/go-auth-microservice => ../GoAuthMicroservice`.

---

## Phase 0 — Setup (1 or)
**Sovorum es:** git, Go modules, project setup.

- [x] `git init`, stexcel `.gitignore` (Go)
- [x] GitHub-um stexcel datark repo `go-auth-microservice`, kapel remote-@
- [x] `go mod init github.com/<tvoy-username>/go-auth-microservice`
- [x] Stexcel `README.md` 2–3 toxov (inch e nakhagic@)
- [x] Arajin commit + push

**Patrast e, erb:** GitHub-um tesnum es repo-n README-ov.

---

## Phase 1 — Go himunqner + arajin Gin server (2–3 or)
**Sovorum es:** package, struct, method, error handling, Gin router.

- [ ] Ancnel [Tour of Go](https://go.dev/tour) — Basics + Methods (ete arden giteS, ancir)
- [x] `api-gateway/cmd/main.go` — Gin server `:8080`-um
- [x] `GET /health` → `{"status":"ok"}`
- [x] Port@ kardal environment variable-ic (`PORT`), default 8080
- [x] Stugel `curl localhost:8080/health`-ov

**Harcer, voronc patasxan@ petq e imanas:** Inch e `:=`-@? Inchu Go-um `error`-@ veradardznum en, voch te `throw`? Inch e `gin.Context`-@?

---

## Phase 2 — REST auth, memory-um (3–4 or)
**Sovorum es:** JSON binding, validation, HTTP status codes, layer-ner (handler → service → repository), interface-ner.

- [x] `User` struct (`ID`, `Email`, `PasswordHash`, `CreatedAt`)
- [x] `UserRepository` **interface** + memory implementation (`map` + `sync.Mutex`) — *skipped: angam Postgres-ov*
- [x] `POST /register` — email + password, validation (`binding:"required,email"`)
- [x] Password-@ hash anel **bcrypt**-ov (`golang.org/x/crypto/bcrypt`)
- [x] `POST /login` — stugel password-@
- [x] Chisht status code-er: 201, 400, 401, 409

**Harc:** Inchu enq repository-n interface sarqum? (Patasxan@ kimanas Phase 3-um.)

---

## Phase 3 — JWT + middleware (2–3 or)
**Sovorum es:** JWT karucvack, signing, Gin middleware.

- [x] Login-ic heto veradardznel **access token** (15 rope) + **refresh token** (7 or)
- [x] `AuthMiddleware` — kardum e `Authorization: Bearer <token>`
- [x] `GET /me` — pashtpanvac route, veradardznum e user-i tvyalner@
- [x] `POST /refresh` — nor access token refresh-ov
- [x] JWT secret@ — env variable-ic, **voch te kod-um**

**Harc:** Inchu access token@ karch e, isk refresh-@ erkar?

---

## Phase 4 — PostgreSQL + Docker (3–4 or)
**Sovorum es:** Docker, docker compose, SQL, migrations, Go-n DB-i het (`pgx`).

- [x] Docker himunqner: image vs container, `docker run postgres`
- [x] `docker-compose.yml` — miayn postgres
- [x] Migration-ner (`golang-migrate`): `users`, `refresh_tokens` table-ner — *Atlas-ov*
- [x] `PostgresUserRepository` — nuyn interface@, nor implementation 🎯 (ahan te inchu interface!)
- [x] Refresh token-ner@ pahel DB-um, `/logout` — jnjel

**Patrast e, erb:** server@ restart es anum, user-@ chi korum.

---

## Phase 5 — GoAuthMicroservice-@ darnum e gRPC service (4–5 or) ⭐ amenakarevor@
**Sovorum es:** Protocol Buffers, gRPC server, code generation, refactoring.

- [ ] **5.1** Tex@ dnel `protoc-gen-go`; grel `proto/auth/v1/auth.proto` (`Register`, `Login`, `Refresh`, `Logout`, `ValidateToken`); `make proto`; commit `gen/`
- [ ] **5.2** Refactor papkaner@: `api-gateway/` → root `cmd/auth-service/` + `internal/` (`git mv`-ov, vor history-n chkorchi)
- [ ] **5.3** `internal/grpcserver` — `AuthServiceServer`-i implementation, vor@ kanchum e arden exac `service` layer-@ (service-@ chi pokhvum! 🎯)
- [ ] **5.4** Service error-ner → gRPC code-er (`AlreadyExists`, `Unauthenticated`, `InvalidArgument`)
- [ ] **5.5** `main.go`: Gin-i poxaren `grpc.NewServer()` `:50051`-um + gRPC health check + reflection
- [ ] **5.6** Jnjel Gin handler-ner@ ev middleware-@ (dranq kgnan MainBackend)
- [ ] **5.7** Stugel `grpcurl`-ov: register → login → validate → refresh → logout

**Harc:** Inchu microservice-ner@ irar het khosum en gRPC-ov, voch te REST-ov?

---

## Phase 6 — MainBackend (nor repo) (4–5 or)
**Sovorum es:** gRPC client, service-neri kap, error mapping, biznes feature, database per service.

- [ ] **6.1** `Go Protfolio/MainBackend` — `git init`, GitHub repo `go-main-backend`, `go mod init`, `replace` → `../GoAuthMicroservice`
- [ ] **6.2** Gin server + `/health`, config `.env`-ic (`AUTH_GRPC_ADDR`)
- [ ] **6.3** `internal/authclient` — gRPC connection auth-service-in (timeout-ov, `context`)
- [ ] **6.4** HTTP → gRPC: `/register`, `/login`, `/refresh`, `/logout`; gRPC code → HTTP status
- [ ] **6.5** `AuthMiddleware` — `ValidateToken` gRPC-ov, `user_id`-@ dnel `gin.Context`-um
- [ ] **6.6** Notes feature: ir Postgres, migration, `POST/GET/DELETE /notes` — user-@ tesnum e miayn ir note-er@
- [ ] **6.7** Stugel: auth-service-@ kangnecnes → MainBackend-@ veradardznum e `503`, voch te panic

**Harc:** Inchu MainBackend-@ chi kardum auth-i DB-n uxxaki? Inch kliner, ete kardar?

---

## Phase 7 — Amboxj system@ Docker-um (2 or)
**Sovorum es:** multi-stage build, compose networking, healthcheck, service discovery.

- [ ] Multi-stage `Dockerfile` amen service-i hamar (image < 30MB)
- [ ] `MainBackend/docker-compose.yml`: auth-db + auth-service (`build: ../GoAuthMicroservice`) + main-db + main-backend
- [ ] `depends_on` + healthcheck-ner (gRPC health-@ auth-i hamar)
- [ ] `.env.example` erku repo-um
- [ ] `Makefile` erku repo-um: `make up`, `make down`, `make proto`, `make test`

**Patrast e, erb:** `MainBackend`-um `docker compose up --build` → amboxj system@ ashkhatum e.

---

## Phase 8 — Testing + CI + Portfolio (3 or)
**Sovorum es:** Go testing, table-driven tests, mocks, GitHub Actions.

- [ ] Unit test-er auth `service` layer-i hamar (mock repository-ov)
- [ ] MainBackend handler test-er (mock authclient-ov — ahan interface-i ogut@ noric)
- [ ] `.github/workflows/ci.yml` erku repo-um — `go vet`, `go test`, `golangci-lint`, docker build
- [ ] Lav **README** erku repo-um: diagram, inchpes gorcarkel, API orinakner, tech stack, link irar
- [ ] Graceful shutdown, structured logging (`slog`), request ID-n ancnel HTTP → gRPC metadata

---

## Bonus (yete uzes avelin)
- [ ] Rate limiting login-i vra
- [ ] Email verification
- [ ] Swagger docs gateway-i hamar
- [ ] Integration test-er `testcontainers-go`-ov
- [ ] Prometheus metrics

---

## Inchpes ashkhatel indz het
- **"Sksum em Phase X"** — kbacatrem tema-n ev kdnem arajin task@
- **"Prcrel em Task X"** — kstugem kod@, khushem inch karelia lavacnel
- **"Khrvel em..."** + sxal@ — kognem, bayc **nakh hushum, voch te patrast kod**
- **"Bacatri X-@"** — kbacatrem orinakov
