# Usa a imagem oficial do Go na versão correta
FROM golang:1.24.2

# Define diretório de trabalho
WORKDIR /file-service

# Copia os arquivos da aplicação
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# Compila a aplicação e cria o binário "main" dentro do container
RUN go build -o /app/main .

# Expõe a porta da aplicação
EXPOSE 8080

# Comando para rodar a aplicação, agora com caminho absoluto
CMD ["/app/main"]