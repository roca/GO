docker run \
        --name unifi-video \
        --cap-add SYS_ADMIN \
        --cap-add DAC_READ_SEARCH \
        -p 1935:1935 \
        -p 6666:6666 \
        -p 7080:7080 \
        -p 7442:7442 \
        -p 7443:7443 \
        -p 7444:7444 \
        -p 7445:7445 \
        -p 7446:7446 \
        -p 7447:7447 \
        -v /Users/romel.campbell/go/src/github.com/GOCODE/camera/rpi/opencv:/var/lib/unifi-video \
        -v /Users/romel.campbell/go/src/github.com/GOCODE/camera/rpi/opencv/videos:/var/lib/unifi-video/videos \
        -e TZ=America/Los_Angeles \
        -e PUID=99 \
        -e PGID=100 \
        -e DEBUG=1 \
        pducharme/unifi-video-controller