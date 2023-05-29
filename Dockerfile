# Use a imagem oficial do Go 1.20 como a imagem base
FROM golang:1.20 AS build

# diretorio padrao
WORKDIR /app

# copia o mod e o main principal
COPY go.mod ./
COPY main.go ./


# importa APIs externas
RUN go mod download github.com/joho/godotenv
RUN go mod download github.com/golang-jwt/jwt/v5
RUN go mod download github.com/go-sql-driver/mysql

# copia todas as pastas

COPY routers/ ./routers
COPY config/ ./config
COPY controllers/agent/ ./controllers/agent
COPY controllers/middleware/ ./controllers/middleware
COPY controllers/receptor/ ./controllers/receptor
COPY middleware/ ./middleware
COPY models/ ./models
COPY repositories/agent/ ./repositories/agent
COPY repositories/receptor/ ./repositories/receptor

# constroi a imagem
RUN go build -o /server


#importa uma versao ultra resumida do linux
FROM gcr.io/distroless/base-debian10

WORKDIR /

# copia o server para o novo diretorio
COPY --from=build /server /server
COPY .env ./.env
EXPOSE 8080


USER nonroot:nonroot

ENTRYPOINT [ "/server" ]



