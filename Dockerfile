FROM golang:1.10

# For later timing purposes
RUN apt-get update && apt-get install -y wget

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    # Clean packages
    && rm dockerize-linux-amd64-$DOCKERIZE_VERSION.tar.gz && apt-get clean

# We use Disconnnect24 as the name is hardcoded into patch's source code.
WORKDIR /go/src/github.com/Disconnect24/Mail-GO
COPY docker/get.sh /go/src/github.com/Disconnect24/Mail-GO
RUN sh get.sh

# Copy needed parts of the Mail-GO source into builder's source,
COPY *.go ./
COPY patch patch
COPY utilities utilities

# Build to name "app".
RUN GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# Wait until there's an actual MySQL connection we can use to start.
CMD ["dockerize", "-wait", "tcp://database:3306", "-timeout", "60s", "/go/src/github.com/Disconnect24/Mail-GO/app"]
