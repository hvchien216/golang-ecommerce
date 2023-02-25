FROM golang:1.19

RUN apt-get update

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.11.0 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.11.0 \