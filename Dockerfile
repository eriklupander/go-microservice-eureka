FROM ofayau/ejre:8-jre
MAINTAINER Micro Service <micro.service@gmail.com>

EXPOSE 8080

ADD bin/goeureka goeureka
ADD templates/*.json templates/

ENTRYPOINT ["./goeureka"]