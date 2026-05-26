# AGENTS.md

## Build and test commands
- **Install dependencies:** `go mod tidy`
- **Build the application:** `make build`
- **Run locally:** `make web`
- **Run all tests:** `make test_local`
- **Run tests with coverage:** `make test_local_coverage`

## Project overview
This microservice resolves search results for food vendors and products for the webview search detail and prepares the tracking payload.

Key entrypoints:
- `cmd/web/main.go`: Application entry point and configuration loading.
- `routes/routes.go`: API route definitions and handler registration.
- `handlers/`: HTTP request handlers and input validation.
- `services/`: Core business logic and orchestration.
- `repositories/`: External service integrations and data access.

### Repository layout
```
.
├── api/             # OpenAPI/Swagger specifications
├── builder/         # Construction of complex data models
├── cmd/web/         # Main application entry point
├── common/          # Shared metrics, headers, and observability utilities
├── constants/       # Application-wide constants
├── handlers/        # HTTP handlers for API endpoints
├── mappers/         # Data transformation logic between models
├── models/          # Go structs for requests, responses, and internal data
├── repositories/    # Data retrieval from external APIs and Redis
├── routes/          # API routing and middleware configuration
├── services/        # Business logic implementation
├── utils/           # Helper functions and configuration management
├── Makefile         # Build, test, and run task definitions
├── go.mod           # Go module dependencies
└── README.md        # Project documentation and setup guide
```

## Key technologies
- **Language:** Go (see `go.mod`)
- **Package manager:** Go Modules
- **Framework:** Gorilla Mux (routing, used in `routes/routes.go`)
- **Testing:** `go test`, `testify` (assertions, see `services/search_test.go`)
- **Other notable libraries:**
  - `peya-go`: PedidosYa internal shared library.
  - `zerolog`: Structured JSON logging (see `common/observability.go`).
  - `dd-trace-go`: Datadog APM tracing integration.
  - `resty`: HTTP client for external service calls.
  - `redis`: Distributed caching (see `repositories/redis_repository.go`).

## Code style guidelines
- Follow standard Go conventions and maintain consistency with the existing codebase.
- **Linting:** Hardcoded secrets detection via `gitleaks` (see `.pre-commit-config.yaml`).

## Testing instructions
- To run the full test suite, use `make test_local`.
- To run a specific package test: `go test ./path/to/package`.
- To run a specific test: `go test -run TestName ./path/to/package`.

## Security considerations
- **Secret Management:** Never commit hardcoded secrets. Use `gitleaks` locally to verify changes.
- **Vault:** The application retrieves secrets from Vault (configured in `utils/config.go`).

## Extra instructions
- **Pull Request Template:** When preparing a Pull Request, follow the template at `.github/PULL_REQUEST_TEMPLATE.md`.
- **Changelog Update:** When preparing a Pull Request, you MUST create a new entry in `CHANGELOG.md` respecting its existing structure (e.g., using `#### [Version](Link)` followed by `> Day Month Year` and categories like `#### 🚀 New Features` or `#### 🐛 Bug Fix`).

### Agent development cycle (default, unless overridden)
- This is the default workflow for this repository.
- **Override:** Only override this workflow if the task explicitly says to use a different workflow (e.g., "override the default workflow" / "follow this workflow instead") or provides its own step-by-step "Workflow / Way of working". If overridden, do not merge workflows.
- **Plan once:** Before coding, propose a TODO list oriented to iterative implementation and STOP. Ask for approval once. After approval, proceed without asking for the plan again unless new information invalidates it.
- **Test-first (when supported):** If the repo has an existing test framework AND documented test command(s) (as listed in this AGENTS.md), write/update tests that specify the new behavior before implementing the production change.
- **Quality gate (must answer “yes” before finishing):**
  - **Task alignment:** Does the change meet every requirement from the original request (and nothing unrelated)?
  - **Tests for new logic:** Did I add/adjust unit tests covering the success path and relevant error or edge cases (when supported in this repo)?
  - **Idiomatic + consistent:** Does the implementation follow repo conventions and language idioms?
  - **Clarity + simplicity:** Is the code easy to read and minimizes complexity?
  - **Error handling:** Are failure modes handled explicitly using the repo’s idioms (exceptions, Result types, validations, retries), with no silent failures?
- **Final verification (only using verified commands listed above):** Run the applicable validation commands that exist in this repo and are listed in "Build and test commands":
  - build/compile validation (if listed)
  - tests covering what you changed (single/scope if documented, otherwise full test command)
  - lint/format/typecheck (if listed)
