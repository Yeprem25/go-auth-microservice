# Go Auth Microservice — Usman plan

Npatak: sovorel **Go, Gin, PostgreSQL, Docker, gRPC, JWT** — kaxucelov irakan auth microservice, vor@ GitHub-um kerevi vorpes portfolio.

Kanon: **kod@ grum es du**, es bacatrum em, stugum em, hushum em. Erb khrvas — grir indz ays fayli Phase/Task hamar@ ev sxal@.

---

## Verjnakan arkhitektura (inch enq karucum)

```
            HTTP/JSON                    gRPC
 Client ─────────────▶ api-gateway ─────────────▶ auth-service ───▶ PostgreSQL
 (curl/Postman)         (Gin, :8080)               (gRPC, :50051)
```

- **api-gateway** — Gin REST API: `/register`, `/login`, `/refresh`, `/me`. Inq@ password chi pahum, uxxaki gRPC-ov harcnum e auth-service-in.
- **auth-service** — gRPC server: grancum, login (bcrypt), JWT access + refresh token, token validation.
- **PostgreSQL** — users ev refresh tokens.
- Amen inch@ ashkhatum e **docker compose up**-ov.

## Verjnakan papkayi karucvack

```
GoAuthMicroservice/
├── proto/auth/v1/auth.proto
├── gen/                     # protoc-ov generacvac kod
├── auth-service/
│   ├── cmd/main.go
│   └── internal/{config,repository,service,grpcserver,token}/
├── api-gateway/
│   ├── cmd/main.go
│   └── internal/{handler,middleware,client}/
├── migrations/
├── docker-compose.yml
├── Makefile
└── README.md
```

> Mi ankhangstacir — sa hima chpetq e stexcel. Ays karucvack@ kerevan kga qayl-qayl.

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

- [ ] `User` struct (`ID`, `Email`, `PasswordHash`, `CreatedAt`)
- [ ] `UserRepository` **interface** + memory implementation (`map` + `sync.Mutex`)
- [ ] `POST /register` — email + password, validation (`binding:"required,email"`)
- [ ] Password-@ hash anel **bcrypt**-ov (`golang.org/x/crypto/bcrypt`)
- [ ] `POST /login` — stugel password-@
- [ ] Chisht status code-er: 201, 400, 401, 409

**Harc:** Inchu enq repository-n interface sarqum? (Patasxan@ kimanas Phase 3-um.)

---

## Phase 3 — JWT + middleware (2–3 or)
**Sovorum es:** JWT karucvack, signing, Gin middleware.

- [ ] Login-ic heto veradardznel **access token** (15 rope) + **refresh token** (7 or)
- [ ] `AuthMiddleware` — kardum e `Authorization: Bearer <token>`
- [ ] `GET /me` — pashtpanvac route, veradardznum e user-i tvyalner@
- [ ] `POST /refresh` — nor access token refresh-ov
- [ ] JWT secret@ — env variable-ic, **voch te kod-um**

**Harc:** Inchu access token@ karch e, isk refresh-@ erkar?

---

## Phase 4 — PostgreSQL + Docker (3–4 or)
**Sovorum es:** Docker, docker compose, SQL, migrations, Go-n DB-i het (`pgx`).

- [ ] Docker himunqner: image vs container, `docker run postgres`
- [ ] `docker-compose.yml` — miayn postgres
- [ ] Migration-ner (`golang-migrate`): `users`, `refresh_tokens` table-ner
- [ ] `PostgresUserRepository` — nuyn interface@, nor implementation 🎯 (ahan te inchu interface!)
- [ ] Refresh token-ner@ pahel DB-um, `/logout` — jnjel

**Patrast e, erb:** server@ restart es anum, user-@ chi korum.

---

## Phase 5 — gRPC (4–5 or) ⭐ amenakarevor@
**Sovorum es:** Protocol Buffers, gRPC, code generation, microservice-neri kap.

- [ ] Tex@ dnel `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc` (kam `buf`)
- [ ] `proto/auth/v1/auth.proto` — `Register`, `Login`, `Refresh`, `ValidateToken`
- [ ] Generacnel Go kod@
- [ ] Teghapokhel auth logic@ → `auth-service` (gRPC server, `:50051`)
- [ ] `api-gateway`-@ darnum e "barak": Gin handler → gRPC client → auth-service
- [ ] gRPC error code-er (`codes.NotFound`, `codes.Unauthenticated`) → HTTP status
- [ ] Stugel `grpcurl`-ov

**Harc:** Inchu microservice-ner@ irar het khosum en gRPC-ov, voch te REST-ov?

---

## Phase 6 — Amboxj system@ Docker-um (2 or)
**Sovorum es:** Dockerfile, multi-stage build, compose networking, healthcheck.

- [ ] Multi-stage `Dockerfile` erkusi hamar (image < 30MB)
- [ ] `docker-compose.yml`: postgres + auth-service + api-gateway
- [ ] `depends_on` + healthcheck
- [ ] `.env.example` fayl
- [ ] `Makefile`: `make up`, `make down`, `make proto`, `make test`

**Patrast e, erb:** `docker compose up --build` → amboxj system@ ashkhatum e.

---

## Phase 7 — Testing + CI + Portfolio (3 or)
**Sovorum es:** Go testing, table-driven tests, mocks, GitHub Actions.

- [ ] Unit test-er service layer-i hamar (mock repository-ov)
- [ ] Table-driven test-er validation-i hamar
- [ ] `.github/workflows/ci.yml` — `go vet`, `go test`, `golangci-lint`, docker build
- [ ] Lav **README**: diagram, inchpes gorcarkel, API orinakner, tech stack
- [ ] Graceful shutdown, structured logging (`slog`)

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
