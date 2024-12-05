migrateup:
	migrate -path db/migration -database "postgres://postgres:21622292a@localhost:5432/mycloud_back?sslmode=disable" -verbose up

createDB:
	psql -U postgres -c "CREATE DATABASE mycloud_back"