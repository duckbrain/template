# My Template Stack

This repo is to set up a template I can clone for new projects. Parts may be refactored out of here as they solidify, so they can be updated in the cloned projects.

- Code gen/compiling/testing (CI) should not require binaries besides `go` and `yarn`
- Dev environment
  - `.editorconfig`
  - [ ] `docker-compose up` and running (avoid Dockerfile for dev; include public image if needed)
    - [ ] Handle running as current user?
  - [ ] `go run` CLI for codegen/tasks
  - [ ] Commands work from root of project
  - [ ] vscode configuration
- CI/CD (GitLab CI)
  - [ ] multistage Dockerfile deployment
  - [ ] Beluga Deployment
  - [ ] Pages deployment (vite index.html)
  - [ ] Review apps
- Backend (Go)
  - [ ] golangci-lint
  - [x] go mod
  - [ ] gqlgen
  - [ ] Models
    - [ ] Pop/Fizz w/o config file?
    - [ ] Repository layer (sqlboiler?)
  - [ ] Tests
    - [ ] Optional DB tests?
    - [ ] Fixtures/schaffolding
    - [ ] Migration test/database
  - [ ] model/migration/dblayer code-generator
  - [ ] cobra for built-in CLI
  - [ ] database migrations
  - [ ] HTML templating w/ VScode syntax support
- Front-end (Typescript/Vue)
  - [ ] eslint? (publish a preset?)
  - [ ] vite
  - [ ] prettier
  - [ ] bootstrap-vue
  - [ ] vue-router
  - [ ] components
    - [ ] layout template
    - [ ] dev demo output

- Functionality
  - [ ] Users/auth
  - [ ] Password set/reset email
  - [ ] Permissions
  - [ ] i18n?

- Build types
  - Debug: source code/tooling present
  - Release: self-containtained single binary

Files/Directories:

- `assets/` front-end source/assets
  - 
- `public/` (gitignored), output of assets 