postgres:
	docker run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root products

dropdb:
	docker exec -it postgres14 dropdb products

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/products?sslmode=disable" -verbose down

test:
	cd backend && go test -v -cover ./...

mock:
	cd backend && mockgen -package mock --destination service/mocks/mock.go github.com/tedoham/fullstack-coding-test/service ProductService 

.PHONY: postgres createdb dropdb migrateup migratedown test mock