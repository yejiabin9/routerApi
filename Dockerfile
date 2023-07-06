FROM alpine
ADD routerApi /routerApi
ADD filebeat.yml /filebeat.yml
ENTRYPOINT [ "/routerApi" ]
