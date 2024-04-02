# Use the official Golang image to create a build artifact.
# This is the first stage of a multi-stage build.
FROM golang:1.22 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything
COPY . .

RUN go build -C jobrunner -o jobrunner .

FROM ubuntu:22.04

RUN apt-get update && apt-get install -y ca-certificates
RUN update-ca-certificates

WORKDIR /root/

COPY --from=builder /app/jobrunner/jobrunner /bin/jobrunner

RUN chmod +x /bin/jobrunner

ENTRYPOINT ["/bin/jobrunner"]
