FROM frolvlad/alpine-glibc:alpine-3

ENV WORKDIR /app
ADD ./main $WORKDIR/main
RUN chmod +x $WORKDIR/main

WORKDIR $WORKDIR
CMD ./main