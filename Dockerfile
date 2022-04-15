FROM golang:1.17.2

RUN apt update && apt install -y wkhtmltopdf

WORKDIR /src/

RUN mkdir api.molino-de-barranco

WORKDIR /src/api.molino-de-barranco

COPY . .

ENV PORT=${PORT}
ENV MONGO_URI=${MONGO_URI}
ENV DB_NAME=${DB_NAME}
ENV EMAIL=${EMAIL}
ENV PASSWORD=${PASSWORD}

RUN go build -o cmd cmd/main.go

EXPOSE 8000

CMD ["./cmd/main"]
