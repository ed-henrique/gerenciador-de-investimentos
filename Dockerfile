FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /app

COPY . .

# As flags utilizadas são para reduzir o tamanho da imagem, descartando
# informações desnecessárias do executável
RUN go build --ldflags "-s -w" -o /go/bin/main cmd/cli/main.go

FROM alpine:3.19 AS exec

ENV TERM "xterm-256color"

WORKDIR /app

COPY --from=builder /go/bin/main .

ENTRYPOINT ["/app/main"]
CMD ["-b"]
