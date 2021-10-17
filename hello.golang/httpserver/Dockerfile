FROM ubuntu
ENV MY_SERVICE_PORT=8080
LABEL my.label="label1"
ADD ./bin/amd64/httpserver /httpserver
EXPOSE 8080
ENTRYPOINT /httpserver
