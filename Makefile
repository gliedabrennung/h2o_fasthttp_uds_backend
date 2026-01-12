H2O_CONF = ./h2o.conf
MAIN_PATH = cmd/app/main.go
H2O_BIN = h2o

.PHONY: dev

dev:
	@echo "Запуск H2O и Go приложения..."

	sudo $(H2O_BIN) -t -c $(H2O_CONF)

	sudo $(H2O_BIN) -c $(H2O_CONF) & \
	H2O_PID=$$!; \
	\
	go run $(MAIN_PATH) & \
	APP_PID=$$!; \
	\
	trap "echo 'Остановка...'; sudo kill $$H2O_PID; kill $$APP_PID" SIGINT SIGTERM; \
	wait