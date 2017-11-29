FROM alpine:3.5
RUN apk add --update ca-certificates

COPY portfolio /

COPY views/* /views/
COPY views/styles/* /views/styles/

WORKDIR "/"

CMD [ "/portfolio"]
EXPOSE 8080
