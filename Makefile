postgres:
	@echo "Starting db_user_service..."
	docker run --name db_user_service -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	@echo "Creating database..."
	docker exec -it db_user_service createdb --username=root --owner=root streamfair_user_service_db

dropdb:
	@echo "Dropping database..."
	docker exec -it db_user_service dropdb streamfair_user_service_db

createmigration:
	@echo "Creating migration..."
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	@echo "Migrating up..."
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/streamfair_user_service_db?sslmode=disable" -verbose up

migratedown:
	@echo "Migrating down..."
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/streamfair_user_service_db?sslmode=disable" -verbose down

dbclean: migratedown migrateup
	clear

sqlc:
	sqlc generate

# testout, dbtestout, apitestout are used to redirect test output to a file
OUT ?= 0

testout: OUT=1
testout: test

dbtestout: OUT=1
dbtestout: dbtest

apitestout: OUT=1
apitestout: apitest

test:
	@if [ $(OUT) -eq 1 ]; then \
		go test -v -cover ./... > tests.log; \
	else \
		go test -v -cover ./... ; \
	fi

dbtest:
	@if [ $(OUT) -eq 1 ]; then \
		go test -v -cover ./db/sqlc > db_tests.log; \
	else \
		go test -v -cover ./db/sqlc ; \
	fi

apitest:
	@if [ $(OUT) -eq 1 ]; then \
		go test -v -cover ./api > api_tests.log; \
	else \
		go test -tags=-coverage -v -cover ./api ; \
	fi

coverage_html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out

server:
	@go run main.go

mock:
	mockgen -source=db/sqlc/store.go -destination=db/mock/store_mock.go

clean:
	rm -f coverage.out tests.log db_tests.log api_tests.log

.PHONY: createdb dropdb postgres migrateup migratedown sqlc test dbtest apitest testout dbtestout apitestout dbclean server mock clean debug