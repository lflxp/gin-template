FROM centos:7
ADD gin-template /gin-template
EXPOSE 8080
ENTRYPOINT ["/gin-template"]
