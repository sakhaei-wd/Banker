postgres:
    docker run --name postgres-0 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -d -p 5432:5432 postgres:alpine	

createdb:    
    docker exec -it postgres-0 createdb --username=postgres --owner=postgres postgres

dropdb:
    docker exec -it postgres-0 dropdb postgres

sqlc:
    sqlc generate
    
.PHONY: postgres createdb dropdb