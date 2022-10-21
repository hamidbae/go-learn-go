docker run --name postgres-fga --network go-fga -p 5432:5432 -e POSTGRES_USER=username -e POSTGRES_PASSWORD=password -e POSTGRES_DB=final-project -d postgres
docker run --rm --name go-fga-final-project -p 8080:8080 hamidbae/go-fga:latest
https://stackoverflow.com/questions/66314534/env-variables-not-coming-through-godotenv-docker
