FROM golang:1.25-alpine AS builder-server
WORKDIR /src
COPY . .
RUN go build -o wvtrserv .

FROM node:lts-alpine AS builder-web
WORKDIR /src
COPY ui/wvtr-front .
RUN npx pnpm build-only

FROM scratch
WORKDIR /app
COPY --from=builder-server /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder-server /src/wvtrserv .
COPY --from=builder-web /src/dist ui/wvtr-front/dist
COPY imgs imgs
ENTRYPOINT [ "/app/wvtrserv" ]
VOLUME [ "/app/db", "/app/tmp/logs" ]
EXPOSE 4210
