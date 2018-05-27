FROM golang
RUN apt-get update 
RUN apt-get install -y unzip
RUN apt-get install -y cmake cmake-curses-gui

RUN cd /opt \
     && wget https://github.com/Itseez/opencv/archive/3.4.1.zip \
     && unzip 3.4.1.zip \
     && cd opencv-3.4.1 \
     && mkdir release \
     && cd release \
     && cmake -D CMAKE_BUILD_TYPE=RELEASE -D CMAKE_INSTALL_PREFIX=/usr/local -D BUILD_NEW_PYTHON_SUPPORT=ON -D INSTALL_C_EXAMPLES=ON -D INSTALL_PYTHON_EXAMPLES=ON  -D BUILD_EXAMPLES=ON .. \
     && make -j4 \
     && make install \
     && ldconfig \
     && rm /opt/3.4.1.zip \
     && rm -R /opt/opencv-3.4.1


RUN go get -d -u gobot.io/x/gobot/...
RUN go get -u -d gocv.io/x/gocv
RUN apt-get install -y sudo
RUN sudo adduser root sudo
WORKDIR $GOPATH/src/gocv.io/x/gocv
RUN yes | make deps
RUN make download
RUN make build
RUN make clean
