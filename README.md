# fibservice

## deploy with docker

* Dockerfile

* docker build
```shell
docker build -t fibservice:v1 .

Sending build context to Docker daemon  78.31MB
Step 1/4 : FROM ubuntu:18.04
 ---> 452a96d81c30
Step 2/4 : COPY bin/fibservice /data/apps/fibservice/bin/
 ---> 56317a268655
Step 3/4 : EXPOSE 8898
 ---> Running in 4013156a76f1
Removing intermediate container 4013156a76f1
 ---> 425341d0df19
Step 4/4 : ENTRYPOINT ["/data/apps/fibservice/bin/fibservice", "-p", ":8898"]
 ---> Running in f35940d298ec
Removing intermediate container f35940d298ec
 ---> 3e29b091d3ee
Successfully built 3e29b091d3ee
Successfully tagged fibservice:v1
```

* docker run
```shell
docker run -d -p 8898:8898 fibservice:v1
```
```shell
docker ps -a |grep fibservice
5c22922f882d        fibservice:v1       "/data/apps/fibservi…"   Less than a second ago   Up 46 seconds               0.0.0.0:8898->8898/tcp             wonderful_jepsen
```
