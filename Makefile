
migrate-up:
	go run main.go migrate:up

migrate-up-to:
	go run main.go migrate:up-to $(version)

migrate-down:
	go run main.go migrate:down

migrate-down-to:
	go run main.go migrate:down-to $(version)

migrate-status:
	go run main.go migrate:status

migrate-create:
	go run main.go migrate:create $(name)

migrate-refresh:
	go run main.go migrate:refresh

seeder:
	go run main.go seeder

seeder-only:
	go run main.go seeder --only $(name)

module-create:
	go run main.go make:module $(name)

worker-create:
	go run main.go make:worker $(name)

worker-run:
	go run main.go worker