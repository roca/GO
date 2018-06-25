FROM resin/rpi-raspbian:latest

ENV APP_HOME /var/app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

RUN cd $APP_HOME

# Update path
ENV	PATH	/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

RUN apt-get -y update && apt-get -y install gcc g++ subversion libjpeg62-turbo-dev imagemagick libav-tools libv4l-dev cmake git make 

RUN git clone https://github.com/jacksonliam/mjpg-streamer.git && \
    cd mjpg-streamer/mjpg-streamer-experimental && \
    export LD_LIBRARY_PATH=. && \
    make

EXPOSE 8080

RUN	echo 'mjpg-streamer/mjpg-streamer-experimental/mjpg_streamer -i "mjpg-streamer/mjpg-streamer-experimental/input_raspicam.so -fps 20" -o "mjpg-streamer/mjpg-streamer-experimental/output_http.so"' > /var/app/start.sh && chmod a+rx /var/app/start.sh

ENTRYPOINT ["/var/app/start.sh"]