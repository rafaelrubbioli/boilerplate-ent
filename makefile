run:
	go run cmd/api.go

migrate:
	go run cmd/migrate.go

gen-ent:
	go run -mod=mod entgo.io/ent/cmd/ent generate ./pkg/ent/schema

gen-graphql:
	go generate pkg/graphql/handler.go
