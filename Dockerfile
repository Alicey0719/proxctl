FROM golang:1.22.2-bullseye

COPY ./app /opt/app/

WORKDIR /opt/app/
