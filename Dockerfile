FROM golang:alpine AS builder

WORKDIR /app

# Add go modules and env files to the WORKDIR and install dependencies.
ADD go.mod ./

# Add code to the WORKDIR and trigger the build process which will assess code quality
# and check if unit tests are passing. Golang binary will be found under /app/mc
ADD . ./

RUN go build -o /app/ -ldflags="-w -s"


# Create final application image.
FROM alpine:3.19 AS final

COPY --from=builder /app/input.txt /input.txt
COPY --from=builder /app/mc /mc

ENTRYPOINT ["/mc"]
