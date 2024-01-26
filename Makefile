run:
	npx tailwindcss -i ./assets/styles/tailwind.css -o ./assets/styles/style.css
	@templ generate
	@go build -o ./tmp/main ./cmd/web/*.go

test:
	@go test -v ./..

migrateup:
	migrate -path db/migrations -database "postgresql://mariogarza:@localhost:5432/godeck?sslmode=disable" -verbose up
	psql -d godeck -U mariogarza -c "INSERT INTO users (username, email, password) VALUES ('mariomoo', 'm@moo.com', 'password');"


migratedown:
	migrate -path db/migrations -database "postgresql://mariogarza:@localhost:5432/godeck?sslmode=disable" -verbose down

reset: migratedown migrateup
