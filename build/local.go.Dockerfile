FROM golang:1.19

RUN apt-get update

RUN GO111MODULE=on go install github.com/vektra/mockery/v2@v2.20.0

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.11.0 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.11.0 \