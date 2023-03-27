FROM golang:1.18-alpine


RUN  apk update \
  && apk upgrade \
  && apk add --update openjdk11 tzdata curl unzip bash \
  && rm -rf /var/cache/apk/*

RUN apk add g++ \
    && rm -rf /var/cache/apk/*


RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools




RUN gcc --version

RUN java -version

RUN python --version

RUN go version


WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o worker ./judge/
RUN ls
ENTRYPOINT [ "./worker" ]

