postgres(){
    psql -U mariogarza -d pata-app
}

migrateup(){
    migrate -path db/migrations -database "postgresql://mariogarza:@localhost:5432/pata-app?sslmode=disable" -verbose up
}

migratedown(){
    migrate -path db/migrations -database "postgresql://mariogarza:@localhost:5432/pata-app?sslmode=disable" -verbose down
}

"$@"
