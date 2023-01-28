FROM golang:alpine AS builder

WORKDIR /src
COPY ./ ./

RUN apk --no-cache add git

RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 go build -o /app . 
# final stage
FROM alpine

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata

WORKDIR /app

COPY --from=builder  /app .
COPY --from=builder /src/.env/ .

# COPY ./web/ web/
# COPY ./public/ public/
# COPY /src/client/dist/ .
ENV ENVIRONMENT=prod
EXPOSE 8080
ENTRYPOINT ./app
CMD ["./app"]