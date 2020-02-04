FROM golang:1.11.0
RUN mkdir /app
ADD ./src/skill-microservices /app
WORKDIR /app
COPY /src/skill-microservices/go.mod .
COPY /src/skill-microservices/go.sum .
RUN sed -i 's/local/development/g' .env
RUN go mod download
COPY . .
EXPOSE ${APP_PORT}
RUN go build -o main .
CMD "/app/main"