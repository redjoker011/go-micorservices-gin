FROM iron/base

EXPOSE 8080
ADD microservices-gin-linux-amd64 /

ENTRYPOINT ["./microservices-gin-linux-amd64"]
