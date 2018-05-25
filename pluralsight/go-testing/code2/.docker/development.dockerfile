# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM golang:latest
# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/github.com/GOCODE/pluralsight/go-testing/code2
WORKDIR /go/src/github.com/GOCODE/pluralsight/go-testing/code2
# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/github.com/GOCODE/pluralsight/go-testing/code2
# Download and install any required third party dependencies into the container.
RUN go get github.com/codegangsta/gin
RUN go get github.com/gorilla/mux
RUN go-wrapper download
RUN go-wrapper install
# Set the PORT environment variable inside the container
ENV APP_PORT 8080
ENV RECEIVER_PORT 5000
ENV VENDOR_PORT 4000
# Expose port 8080 to the host so we can access our application
EXPOSE 3000
EXPOSE 4000
EXPOSE 5000
EXPOSE 8080
# Now tell Docker what command to run when the container starts
#CMD ["go-wrapper", "run"]
CMD gin run