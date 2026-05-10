FROM golang:1.25-trixie AS builder-server
WORKDIR /src
COPY . .
RUN go build -o wvtrserv .

FROM node:lts AS builder-web
WORKDIR /src
COPY ui/wvtr-front .
RUN npx pnpm build-only

FROM chainguard/wolfi-base:latest
WORKDIR /app
COPY --from=builder-server /src/wvtrserv .
COPY --from=builder-web /src/dist ui/wvtr-front/dist
COPY imgs imgs
RUN mkdir -p tmp/logs
ENTRYPOINT [ "/app/wvtrserv" ]
EXPOSE 4210
