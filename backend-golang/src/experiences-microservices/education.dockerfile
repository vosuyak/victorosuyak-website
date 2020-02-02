FROM golang:1.11.0
RUN mkdir /app
ADD ./src/experience-microservices /app
WORKDIR /app
COPY /src/experience-microservices/go.mod .
COPY /src/experience-microservices/go.sum .
RUN go mod download
COPY . .
EXPOSE 8081
RUN go build -o main .
CMD "/app/main"