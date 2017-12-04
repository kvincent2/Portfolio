FROM alpine:3.5
RUN apk add --update ca-certificates

COPY portfolio /

COPY views/* /views/
COPY views/styles/* /views/styles/

WORKDIR "/"

CMD [ "/portfolio"]
EXPOSE 8080

LABEL 	maintainer="Kristina Vincent <vincentkb@live.com> <https://github.com/kvincent2>"
