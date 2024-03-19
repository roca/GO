FROM linuxkit/grub-dev:e94da02aac3a39fec34047051a40d367e61cd0a0-amd64

RUN mkdir -p /etc/default

RUN echo 'GRUB_CMDLINE_LINUX_DEFAULT="loglevel=4 slub_debug=P page_poison=1 psi=1"' > /etc/default/grub