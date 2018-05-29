FROM rcampbell/rpi-golang

RUN go get -d -u gobot.io/x/gobot/...

# Define default command
CMD ["bash"]
