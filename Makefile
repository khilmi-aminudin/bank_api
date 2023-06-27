DB_SOURCE=postgresql://root:secret@localhost:5432/bank_db?sslmode=disable

postgresql :
	docker run --name postgresql -p 5432:5432 -e TZ=Asia/Jakarta -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

execdb :
	docker exec -it postgresql psql -d bank_db

uuid-db :
	docker exec -it postgresql psql -d bank_db -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'

createdb :
	docker exec -it postgresql createdb --username=root --owner=root bank_db

dropdb :
	docker exec -it postgresql dropdb bank_db

rundb :
	docker start postgresql

initmigrate :
	migrate create -ext sql -dir db/migrations -seq init_schema

migrateup :
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up

migratedown :
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose down

migrateup1 :
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose up 1

migratedown1 :
	migrate -path db/migrations -database "$(DB_SOURCE)" -verbose down 1

sqlc :
	sqlc generate

db_docs:
	dbdocs build docs/db.dbml

db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

test :
	go test -v -cover ./...

runserver :
	go run cmd/main.go

build:
	docker build -t ewallet .

mock :
	mockery --name=Repository --dir=./repositories --output=db/mocks --outpkg=mocks

.PHONY : postgresql execdb createdb initmigrate migrateup migratedown migrateup1 migratedown1 sqlc db_docs db_schema test runserver mock