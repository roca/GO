FROM rcampbell/rpi-golang

RUN apt-get update
RUN apt-get autoremove libopencv-dev python-opencv

ENV APP_HOME /var/app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY install-opencv.sh $APP_HOME
RUN chmod +x install-opencv.sh && ./install-opencv.sh

CMD ["bash"]