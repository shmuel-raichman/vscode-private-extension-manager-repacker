# B"H
FROM golang as builder

WORKDIR /app/golang/ext
COPY . /app/golang/ext/

# ARG OS=linux
# ARG VERSION=unknown
ARG ARTIFACT_NAME=repacker
ARG ARTIFACT_FULL_NAME=repacker
ARG GO_MODE_NAME=vscode-ext

ENV GOPATH=/app/golang
ENV GOBIN=/app/golang/bin

# ENV OS=$OS
ENV ARTIFACT_NAME=$ARTIFACT_NAME
ENV ARTIFACT_FULL_NAME=$ARTIFACT_FULL_NAME
# ENV VERSION=$VERSION

RUN mkdir bin
# Linux build
RUN go build -o bin/$ARTIFACT_FULL_NAME $GO_MODE_NAME
# Windows build
RUN env GOOS=windows GOARCH=amd64 go build -o bin/$ARTIFACT_FULL_NAME.exe $GO_MODE_NAME

# Linux build
# RUN go build -o $ARTIFACT_FULL_NAME -ldflags="-X $GO_MODE_NAME/flags.BuildVersion=$VERSION" $GO_MODE_NAME
# Windows build
# RUN env GOOS=windows GOARCH=amd64 go build -o $ARTIFACT_FULL_NAME.exe -ldflags="-X $GO_MODE_NAME/flags.BuildVersion=$VERSION" $GO_MODE_NAME

FROM node:14-buster-slim

WORKDIR /app

COPY --from=builder /app/golang/ext/bin/repacker /app/repacker

ENTRYPOINT /bin/sh
CMD ["cat"]