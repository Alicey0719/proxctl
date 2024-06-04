FROM golang:1.22.4-bullseye

COPY ./app /opt/app/

WORKDIR /opt/app/
