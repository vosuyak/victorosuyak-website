FROM golang:1.11.0
RUN mkdir /app
ADD ./src/education-microservices /app
WORKDIR /app
COPY /src/education-microservices/go.mod .
COPY /src/education-microservices/go.sum .
RUN go mod download
COPY . .
EXPOSE 8081
RUN go build -o main .
CMD "/app/main"