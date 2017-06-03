# Base this docker container off the official golang docker image.
# Docker containers inherit everything from their base.
FROM gobuffalo/buffalo
# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /go/src/github.com/GOCODE/buffalo
WORKDIR /go/src/github.com/GOCODE/buffalo
# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY . /go/src/github.com/GOCODE/buffalo
# Download and install any required third party dependencies into the container.
# Set the PORT environment variable inside the container
ENV PORT 8080
# Expose port 8080 to the host so we can access our application
EXPOSE 3000
# Now tell Docker what command to run when the container starts
#CMD ["go-wrapper", "run"]
ENTRYPOINT ["tail", "-f", "/dev/null"]