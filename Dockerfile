# Use a imagem oficial do Golang
FROM golang:1.16 as builder

# Defina o diretório de trabalho no contêiner
WORKDIR /

# Copie o código-fonte para o contêiner
COPY . .

# Compile o aplicativo Go
RUN go build -o main .

CMD ["./main"]

## Use uma imagem mínima do alpine para reduzir o tamanho da imagem final
#FROM alpine:latest
#
## Defina o diretório de trabalho no contêiner
#WORKDIR /
#
## Copie o executável do compilador de estágio anterior
##COPY --from=builder /app/main .
#
## Exponha a porta 8080, que a aplicação Go escutará
#EXPOSE 8080
#
## Execute o aplicativo Go quando o contêiner for iniciado
#CMD ["go run main.go"]