FROM resin/raspberrypi3-golang:1.9

ENV APP_HOME /var/app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY install-opencv.sh $APP_HOME
RUN chmod +x install-opencv.sh && ./install-opencv.sh


# Define default command
CMD ["tail", "-f", "/dev/null"]