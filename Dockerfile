FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env .env
ENV PORT=${PORT}
ENV LOG_LEVEL=${LOG_LEVEL}
ENV DB_URL=${DB_URL}
ENV FAST_FOREX_API=${FAST_FOREX_API}

EXPOSE 8080

CMD ["go", "run", "./main.go", "rest"]
