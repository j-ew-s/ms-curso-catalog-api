FROM golang:1.17.3  AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o  /ms-curso-catalog-api



FROM gcr.io/distroless/base-debian10 

WORKDIR /
COPY --from=build /ms-curso-catalog-api  /catalog-api

EXPOSE 5200

USER nonroot:nonroot

CMD ["/catalog-api"]Ÿ