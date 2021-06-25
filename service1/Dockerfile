FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
ADD .builds/binary  /binary 
ENTRYPOINT ["./binary"]