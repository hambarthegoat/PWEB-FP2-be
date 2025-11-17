run:
	go run cmd/api/main.go

migrate: 
	go run migrations/auto_migrate.go

seed:
	go run cmd/seed/main.go