FROM golang:latest as builder

WORKDIR /helmAPI

COPY go.* .

RUN go mod download

COPY . .

RUN go build -o helmAPI



FROM ubuntu:18.04

WORKDIR /helmAPI

COPY --from=builder /helmAPI/helmAPI .

COPY kubeconf.yaml .

RUN apt-get update &&\
    apt-get -y install curl

WORKDIR /helm

RUN curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 | bash

WORKDIR /helmAPI

ENV KUBECONFIG=/helmAPI/kubeconf.yaml

EXPOSE 8080

ENTRYPOINT ["./helmAPI"]