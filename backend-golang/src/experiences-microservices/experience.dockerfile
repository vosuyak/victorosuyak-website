FROM golang:1.11.0
RUN mkdir /app
ADD ./src/experiences-microservices /app
WORKDIR /app
COPY /src/experiences-microservices/go.mod .
COPY /src/experiences-microservices/go.sum .
RUN go mod download
COPY . .
EXPOSE 8082
RUN go build -o main .
CMD "/app/main"