FROM rcampbell/gocv

ENV APP_HOME /var/app

RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

RUN export PKG_CONFIG_PATH=/usr/lib/pkgconfig

COPY camera.go $APP_HOME
RUN go build camera.go
ENV PORT 3000
EXPOSE 3000
CMD ["./camera"]