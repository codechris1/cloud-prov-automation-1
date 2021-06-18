FROM golang:latest
ARG APP_NAME
ENV APP_NAME=${APP_NAME}

WORKDIR /go
COPY cmd/${APP_NAME}/* src/${APP_NAME}/

RUN go env -w GO111MODULE=off
RUN go install ${APP_NAME}
RUN rm -Rf src/*

CMD ["sh","-c","${APP_NAME}"]