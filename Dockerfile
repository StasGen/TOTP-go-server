FROM golang:1.16-alpine
WORKDIR /app

ARG SECRET

COPY go.mod go.sum ./

RUN go mod download


COPY *.go ./

# Set necessary environmet variables needed for our image and build the API server.
# ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
# RUN go build -ldflags="-s -w" -o apiserver .
# RUN go build -a -installsuffix cgo -o app .


RUN go build -o /docker-gs
CMD [ "/docker-gs" ]





# FROM scratch

# Copy binary and config files from /build to root folder of scratch container.
# COPY --from=builder ["/build/app", "/"]

# Command to run when starting the container.
# ENTRYPOINT ["/app"]
