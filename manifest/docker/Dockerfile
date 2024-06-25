FROM frolvlad/alpine-glibc:alpine-3

RUN apk update && apk add tzdata ca-certificates bash

ENV WORKDIR /app
ADD ./main $WORKDIR/main
RUN chmod +x $WORKDIR/main

WORKDIR $WORKDIR
CMD /app/main