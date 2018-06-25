FROM rcampbell/rpi-golang

RUN apt-get update
RUN apt-get autoremove libopencv-dev python-opencv

ENV APP_HOME /var/app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

# need this some where ??????
# ln -s /usr/lib/arm-linux-gnueabihf/libQtOpenGL.so /usr/lib/arm-linux-gnueabihf/libQt5OpenGL.so.5.3.2

COPY install-opencv.sh $APP_HOME
RUN chmod +x install-opencv.sh && ./install-opencv.sh

CMD ["bash"]