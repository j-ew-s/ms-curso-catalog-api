# ms-curso-catalog-api


# Catalog Mod
go mod init github.com/j-ew-s/ms-curso-catolog-api

# Catalog API porta
:5200


# Comunica com outros serviços:

1 - AUTH GRPC   pela porta :5500 


# Pacotes que utilizamos 
- SQL : go get -u github.com/go-sql-driver/mysql


# DOCKER 


docker build --tag catalog-api:v0.0.1 -f Dockerfile   .

docker compose build

docker compose up


docker run -d  --name mongo-on-docker  -p 27017:27017 -e MONGO_INITDB_ROOT_USERNAME=mongoadmin -e MONGO_INITDB_ROOT_PASSWORD=secret mongo