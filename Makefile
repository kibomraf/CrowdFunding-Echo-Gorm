crowdfunding:
	docker run --name crowdfunding -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123qwe -d postgres
createdb:
	docker exec -it crowdfunding createdb --username=root --owner=root crowdfunding
dropdb:
	docker exec -it crowdfunding dropdb crowdfunding
migrateup:
	migrate -path database/migration -database "postgresql://root:123qwe@localhost:5432/crowdfunding?sslmode=disable" -verbose up
migratedown:
	migrate -path database/migration -database "postgresql://root:123qwe@localhost:5432/crowdfunding?sslmode=disable" -verbose down
.PHONY: crowdfunding createdb dropdb migrateup migratedown