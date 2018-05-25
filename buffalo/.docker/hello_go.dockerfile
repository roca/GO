FROM ubuntu:latest
# Create a directory inside the container to store all our application and then make it the working directory.
RUN mkdir -p /var/app
WORKDIR /var/app
# Copy the example-app directory (where the Dockerfile lives) into the container.
COPY ./hello_go/bin/hello_go /var/app
# Download and install any required third party dependencies into the container.
# Set the PORT environment variable inside the container
ENV PORT 3000
# Expose port 8080 to the host so we can access our application
EXPOSE 3000
# Now tell Docker what command to run when the container starts
#CMD ["go-wrapper", "run"]
ENTRYPOINT ["./hello_go"]