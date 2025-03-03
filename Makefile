.PHONY: migrate-up migrate-down migrate-status

kill:
	@pid=$$(lsof -t -i :3000); if [ -n "$$pid" ]; then kill -9 $$pid; echo "Processo $$pid finalizado."; else echo "Nenhum processo na porta 3000."; fi

generate:
	cd cmd/app && mv wire.txt wire.go && wire && mv wire.go wire.txt && cd ../..

migrate-up:
	go run cmd/migrate/main.go up

migrate-down:
	go run cmd/migrate/main.go down

migrate-status:
	go run cmd/migrate/main.go status

migrate-create:
	go run cmd/migrate/main.go -command=create:migration -name=$(name) -type=$(type)

seed-create:
	go run cmd/migrate/main.go -command=create:seed -name=$(name) -type=$(type)

seed-up:
	go run cmd/migrate/main.go -command=seed
