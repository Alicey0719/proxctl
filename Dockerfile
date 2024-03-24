FROM golang:1.22.1-bullseye

COPY ./app /opt/app/

WORKDIR /opt/app/
