# Imagem base
FROM golang:1.24.4-alpine3.22

# Diretório de trabalho
WORKDIR /app

# Copia os arquivos
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Compila a aplicação
RUN go build -o main .

# Expõe a porta
EXPOSE 5000

# Comando para rodar a API
CMD ["./main"]
