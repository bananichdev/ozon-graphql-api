FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY ["app/go.mod", "app/go.sum", "./"]
RUN go mod download
RUN go install -tags "postgres" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY app ./
RUN go build -o ./bin/main ./cmd

# wait 2 seconds for init db
CMD sleep 2 && go test ./internal/db && migrate -database ${POSTGRESQL_URL} -path internal/db/migrations up && ./bin/main -${MODE}

EXPOSE 8000
