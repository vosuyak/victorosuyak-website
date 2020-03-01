FROM golang:1.11.0
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN sed -i 's/local/development/g' /app/.env
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
EXPOSE 8081
RUN go build -o main .
CMD "/app/main"