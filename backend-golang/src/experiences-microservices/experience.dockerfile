FROM golang:1.11.0
RUN mkdir /app
ADD ./src/experiences-microservices /app
WORKDIR /app
COPY /src/experiences-microservices/go.mod .
COPY /src/experiences-microservices/go.sum .
RUN sed -i 's/local/development/g' .env
RUN go mod download
COPY . .
EXPOSE ${APP_PORT}
RUN go build -o main .
CMD "/app/main"