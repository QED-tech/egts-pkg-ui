run-dev:
	air --build.cmd "go build -o tmp/app cmd/main.go" --build.bin "./tmp/app"
migrations-create:
	goose create $(name) sql
migrations-up:
	goose up