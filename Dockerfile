FROM golang:latest 
LABEL maintainer="rosie@rosie.dev"
RUN go get github.com/stretchr/testify/assert && go get gopkg.in/retry.v1
WORKDIR /app 
COPY /apiclient /app/
CMD ["go", "test", "-v", "-cover"]