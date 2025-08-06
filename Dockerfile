FROM golang:1.24.6-bullseye

COPY ./app /opt/app/

WORKDIR /opt/app/
