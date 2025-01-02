FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go build -o backend ./cmd/main.go

FROM node:lts AS frontend-builder
WORKDIR /frontend
COPY public/package.json ./
COPY public/yarn.lock ./
RUN npm install
COPY public ./
RUN npm run build

FROM alpine:latest

COPY --from=builder /app/backend ./

COPY --from=frontend-builder /frontend/dist ./public

EXPOSE 8080
EXPOSE 3000

CMD ["./backend"]
