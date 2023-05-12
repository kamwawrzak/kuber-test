FROM golang:1.19-buster AS build

WORKDIR /build

COPY go.mod go.sum /build/
COPY ./config/service-config.yaml /build/config/
RUN go mod download

COPY cmd /build/cmd
COPY app /build/app

RUN cd cmd && env go build -o kuber

#FROM busybox:1.35.0-uclibc as busybox

#FROM gcr.io/distroless/static:nonroot
FROM ubuntu:latest
#COPY --from=busybox /bin/sh /bin/sh
#RUN /bin/bash -c 'source /opt/ros/melodic/setup.bash'

WORKDIR /app
COPY --from=build ./build/cmd/kuber ./build/config/service-config.yaml /app/

ENTRYPOINT [ "./kuber" ]


