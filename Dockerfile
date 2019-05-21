FROM centos:7
ENV UI_PORT=8080
ADD spa-template /spa-template
ADD templates /templates
ADD static /static
ENTRYPOINT ["/spa-template"]