# base go image
# FROM golang:1.18-alpine as builder

# This is where the docker image we're building
# RUN mkdir /app

# copy everything from the current folder to the folder we build
# COPY . /app

# setting working directory
# WORKDIR /app

# build the go code
# RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

# add the executable flag to app/brokerApp just to make sure it's executable
# RUN chmod +x /app/brokerApp

# build a tiny docker image

# we do that by saying from and the smallest one can find is Alpine:latest
FROM alpine

# this is on a new docker image
RUN mkdir /app

# now copy from builder which named on line 2 and copy from /app/brokerApp to /app
# COPY --from=builder /app/brokerApp /app
COPY brokerApp /app

# command to run
# when run this , this should first of all build all of the code on one docker image and create a much smaller docker image-
# and just copy over the executable.
CMD [ "/app/brokerApp" ]