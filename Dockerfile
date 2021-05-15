FROM golang:1.14.7-alpine as builder
ARG GITHUB_TOKEN 

RUN set -eux \
    && apk -U add git
RUN git config --global url."https://$GITHUB_TOKEN:@github.com/".insteadOf "https://github.com/"

WORKDIR /app

COPY go.mod . 
COPY go.sum .

RUN go mod tidy

RUN go mod download 

COPY . . 
RUN CGO_ENABLED=0 go build -o data_selection_pretexting ./cmd/main.go 

ENTRYPOINT ["/app/data_selection_pretexting"]