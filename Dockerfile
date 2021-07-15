# B"H
FROM golang:1.16.6-alpine3.14 as  builder

WORKDIR /app/golnag/vscode-ext
COPY . /app/golnag/vscode-ext/

# ARG OS=linux
ARG VERSION=unknown
ARG ARTIFACT_NAME=vscode-ext
ARG ARTIFACT_FULL_NAME=vscode-ext
ARG GO_MODE_NAME=vscode-ext

ENV GOPATH=/app/golang
ENV GOBIN=/app/golang/bin

# ENV OS=$OS
ENV ARTIFACT_NAME=$ARTIFACT_NAME
ENV ARTIFACT_FULL_NAME=$ARTIFACT_FULL_NAME
ENV VERSION=$VERSION

# # Linux build
# RUN go build -o $ARTIFACT_FULL_NAME -ldflags="-X $GO_MODE_NAME/flags.BuildVersion=$VERSION" $GO_MODE_NAME
# # Windows build
# RUN env GOOS=windows GOARCH=amd64 go build -o $ARTIFACT_FULL_NAME.exe -ldflags="-X $GO_MODE_NAME/flags.BuildVersion=$VERSION" $GO_MODE_NAME
# Linux build
RUN go build -o $ARTIFACT_FULL_NAME $GO_MODE_NAME
# Windows build
RUN env GOOS=windows GOARCH=amd64 go build -o $ARTIFACT_FULL_NAME.exe $GO_MODE_NAME


FROM node:14-buster-slim as runtime

COPY --from=builder /app/golang/vscode-ext/vscode-ext /app