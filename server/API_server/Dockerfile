# API Imageの作成
FROM golang:1.12.4-alpine
# Gopkg.toml 等を事前にコピーして dep ensure が実行できるようにする
COPY src/ /go/src/api/

# dep ensure を行うプロジェクトルートに移動する
WORKDIR /go/src/api/

# 必要なパッケージをイメージにインストールする
RUN apk update \
  && apk add --no-cache git \
  && export GO111MODULE=on \
  && go build -o binary/orenotorero

EXPOSE 3000