
run: 
	@~/go/bin/templ generate
	@go run cmd/web/*.go

test:
	@go test -v ./..

migrateup:
	migrate -path db/migrations -database "postgresql://mariogarza:@localhost:5432/godeck?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://mariogarza:@localhost:5432/godeck?sslmode=disable" -verbose down
