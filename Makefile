server:
	go run ./...

dev:
	@echo "Starting dev environment..."
	@air -c ./.air.toml
	@echo "Dev environment started"

.PHONY: server dev