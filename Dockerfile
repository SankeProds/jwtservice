FROM golang

LABEL maintainer="Carlos Borrero <carlosborrero@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/SankeProds/jwtservice

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v

# Install the package
RUN go install -v

# This container exposes port 8080 to the outside world
EXPOSE 4321

# Run the executable
CMD ["jwtservice"]