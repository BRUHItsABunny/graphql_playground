# GRAPHQL Playground
I was bored and wanted to try some Graph+Postgres
Based on [this article](https://hasura.io/blog/building-a-graphql-api-with-golang-postgres-and-hasura/)

Additions on top of the aforementioned article:
* DB Migrations [like this](https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md)

### Notes

Run Docker Postgres:

`docker run --name postgres-db -e POSTGRES_PASSWORD=docker -e POSTGRES_DB=playground_db -p 5432:5432 -d postgres`

To Shell into Postgres:

`docker exec -it postgres-db psql -U postgres -W`

To create a database in Postgres after connecting:
`create database playground_db;`

For the CMD's above, the DSN would be:

`postgres://postgres:docker@localhost:5432/playground_db?sslmode=disable`
