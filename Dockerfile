FROM golang:1.20 AS builder

RUN apt-get update -y 
RUN wget -O /usr/local/bin/dumb-init https://github.com/Yelp/dumb-init/releases/download/v1.2.5/dumb-init_1.2.5_x86_64

RUN chmod +x /usr/local/bin/dumb-init

WORKDIR /build/bin/
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

RUN make build

FROM scratch

WORKDIR /tmp

WORKDIR /config
WORKDIR /

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/local/bin/dumb-init /usr/bin/dumb-init
COPY --from=builder ["/build/bin/api", "/"]
COPY --from=builder ["/build/config/config.yaml", "/config"]

# Export necessary port.
EXPOSE 8082
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

# Command to run when starting the container.
CMD ["./api"]