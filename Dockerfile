FROM paas-docker-artifactory.gannettdigital.com/paas-scratch-base:latest
MAINTAINER Fill-Me <fill-me@gannett.com>

COPY {NAME} /

COPY views/* /views/
COPY views/styles/* /views/styles/

WORKDIR "/"

CMD [ "/{NAME}"]
EXPOSE 8080
