# Usar a imagem base do Go
FROM golang:1.22

RUN apt-get update && apt-get install -y netcat-openbsd

# Configurar o diretório de trabalho
WORKDIR /app

# Copiar os arquivos da aplicação
COPY . .

# Comando padrão
CMD ["go", "run", "main.go"]