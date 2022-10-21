docker run --name postgres-fga --network go-fga -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=final-project -d postgres
