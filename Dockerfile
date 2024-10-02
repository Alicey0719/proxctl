FROM golang:1.23.2-bullseye

COPY ./app /opt/app/

WORKDIR /opt/app/
