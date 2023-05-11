FROM golang:1.19-buster AS build

WORKDIR /build

COPY go.mod go.sum /build/
COPY ./config/service-config.yaml /build/config/
RUN go mod download

COPY cmd /build/cmd
COPY app /build/app

RUN cd cmd && env go build -o kuber


FROM gcr.io/distroless/static:nonroot

WORKDIR /app
COPY --from=build ./build/cmd/kuber ./build/config/service-config.yaml /app/

ENTRYPOINT [ "./kuber" ]


