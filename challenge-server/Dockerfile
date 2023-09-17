# Use the Official Golang image
FROM golang:1.16 as builder

# Choosing to workdir / will make the next commands run from the root of the container
WORKDIR /

# Copy the source code to the container
COPY .. .

# Compile the Go application
RUN go build -o main .

CMD ["./main"]
