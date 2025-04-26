.PHONY: migrate migrate-fresh migrate-reset

migrate:
	@echo "Running migrations"
	@goose up

migrate-fresh:
	@echo "Droping all migrations"
	@goose reset
	@echo "Running migrations"
	@goose up

migrate-reset:
	@echo "Droping all migrations"
	@goose reset
