FROM ubuntu:23.04 as base

# install golang
RUN apt-get update && apt-get install -y golang

# copy the source code to the container
COPY . /app

# set the working directory
WORKDIR /app

# build the source code
RUN go build -o main .

# run the application
CMD ["/app/main"]

