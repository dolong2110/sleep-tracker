.PHONY: migrate-create migrate-up migrate-down migrate-force init

PWD = $(shell pwd)
MPATH = $(PWD)/sleep-tracker/internal/migrations
PORT = 5432

# Default number of migrations to execute up or down
N = 2

# Commands for migrating tables
migrate-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(MPATH) -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(MPATH) -database postgres://postgres:password@localhost:$(PORT)/postgres?sslmode=disable force $(VERSION)