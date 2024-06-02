# Etapa de construção
FROM golang:1.21.0-bullseye AS build

WORKDIR /app

# Copia os arquivos go.mod, go.sum e main.go
COPY go.mod ./
COPY go.sum ./
COPY main.go ./

# Baixa as dependências
RUN go mod download

# Copia todas as pastas
COPY config/ ./config
COPY controllers/agent/ ./controllers/agent
COPY controllers/middleware/ ./controllers/middleware
COPY controllers/receptor/ ./controllers/receptor
COPY Logs/ ./Logs
COPY middleware/ ./middleware
COPY models/ ./models
COPY repositories/agent/ ./repositories/agent
COPY repositories/receptor/ ./repositories/receptor
COPY routers/ ./routers

# Constroi a imagem
RUN go build -o /server

# Etapa final
FROM gcr.io/distroless/base-debian10

WORKDIR /

# Cria o usuário não root
RUN addgroup --system nonroot && adduser --system --ingroup nonroot nonroot

# Copia o binário construído e o arquivo .env
COPY --from=build /server /server
COPY .env ./.env

# Copia o diretório de logs criado na etapa anterior
COPY --from=build /Logs /Logs

# Ajusta permissões no estágio final
RUN chown -R nonroot:nonroot /Logs && chmod -R 777 /Logs

# Define a porta que será exposta
EXPOSE 8080

# Define um usuário não root
USER nonroot:nonroot

# Copia o script de inicialização
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

# Define o ponto de entrada para o contêiner
ENTRYPOINT ["/entrypoint.sh"]
