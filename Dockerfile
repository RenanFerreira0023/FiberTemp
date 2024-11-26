# Etapa de construção
FROM golang:1.21.0-bullseye AS build

WORKDIR /app

# Copia os arquivos go.mod, go.sum e main.go
COPY go.mod go.sum main.go ./

# Baixa as dependências
RUN go mod download

# Copia o código fonte
COPY . .

# Constroi a imagem
RUN go build -o /server

# Etapa final
FROM gcr.io/distroless/base-debian10

WORKDIR /

# Copia o binário construído e o arquivo .env
COPY --from=build /server /server
COPY .env ./.env

# Copia o diretório de logs
COPY --from=build /Logs /Logs

# Ajusta permissões no estágio final
RUN chmod -R 777 /Logs

# Define a porta que será exposta
EXPOSE 8080

# Define um usuário não root
USER nonroot:nonroot

# Copia o script de inicialização
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Define o ponto de entrada para o contêiner
ENTRYPOINT ["/entrypoint.sh"]
