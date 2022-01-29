# Figment sensor log processor

Application written in Go which processes sensor logs and evaluates them accordingly.

---

### Local setup

#### Prerequisites

- Go installed
- Docker installed (if running application via Docker)

You can start your local environment in different ways:

```
# Run application locally
make run

---

# Run application via Docker
make run-docker
```

### Test coverage

You can run all tests with following command:

```
# Running this command from the root of the project
make test
```

### Notices

Altering `.log` file changes the application output.
