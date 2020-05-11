FROM golang:latest

# ENV GO111MODULE=on
ENV LANG C.UTF-8
ENV APP_HOME $GOPATH/src/go_api

RUN mkdir $APP_HOME
WORKDIR $APP_HOME

RUN go get -u github.com/gin-gonic/gin && \
  go get github.com/jinzhu/gorm && \
  go get github.com/jinzhu/gorm/dialects/postgres && \
  go get github.com/go-delve/delve/cmd/dlv

# コンテナ実行時のデフォルトを設定する
# ライブリロードを実行する
# CMD gin -i run
EXPOSE 3000 2345

ENTRYPOINT ["sh", "./script/web_entrypoint.sh"]
