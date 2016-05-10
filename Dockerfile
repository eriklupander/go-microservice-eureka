# Replacing standard java-8 Docker image with a 32 bit, embedded jre 8 version, reduces the memory consumption with some 50%
# FROM java:8
FROM ofayau/ejre:8-jre
MAINTAINER Micro Service <micro.service@gmail.com>

EXPOSE 8080

ADD bin/goeureka goeureka
ADD templates/*.json templates/

ENTRYPOINT ["./goeureka"]