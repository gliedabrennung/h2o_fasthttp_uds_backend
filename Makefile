CADDYFILE=./Caddyfile
MAIN_PATH=cmd/app/main.go

.PHONY: dev

dev:
	@echo "Запуск инфраструктуры (Ctrl+C для завершения)..."

	sudo caddy run --config $(CADDYFILE) & \
	CADDY_PID=$$!; \
	\
	go run $(MAIN_PATH) & \
	APP_PID=$$!; \
	\
	wait