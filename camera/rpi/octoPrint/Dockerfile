FROM resin/rpi-raspbian:latest

ENV APP_HOME /var/app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

RUN sudo apt update
RUN sudo apt install python-pip python-dev python-setuptools python-virtualenv git libyaml-dev build-essential
RUN mkdir OctoPrint && cd OctoPrint
RUN virtualenv venv
RUN source venv/bin/activate
RUN pip install pip --upgrade
RUN pip install https://get.octoprint.org/latest

EXPOSE 5000

ENTRYPOINT [ "pi@raspberrypi ~ $ ~/OctoPrint/venv/bin/octoprint serve" ]