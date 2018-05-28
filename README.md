# fibservice

## deploy with docker

* Dockerfile

* docker build
```shell
docker build -t fibservice:v1 .

Sending build context to Docker daemon  94.69MB
Step 1/5 : FROM ubuntu:18.04
 ---> 452a96d81c30
Step 2/5 : COPY bin/fibservice /data/apps/fibservice/bin/
 ---> 14a949f378fb
Step 3/5 : EXPOSE 8898 8899
 ---> Running in 6e787873a01a
Removing intermediate container 6e787873a01a
 ---> 20c846e98270
Step 4/5 : CMD ["-p", ":8898", "-pp", ":8899"]
 ---> Running in aa86c72aa30d
Removing intermediate container aa86c72aa30d
 ---> e1b23ded9a13
Step 5/5 : ENTRYPOINT ["/data/apps/fibservice/bin/fibservice"]
 ---> Running in 7da55be436fe
Removing intermediate container 7da55be436fe
 ---> 5e8fb4bd1330
Successfully built 5e8fb4bd1330
Successfully tagged fibservice:v1
```

* docker run
```shell
docker run -d -p 8898:8898 -p 8899:8899 fibservice:v1
```

```shell
docker ps -a |grep fibservice
0ff82ec07398        fibservice:v1       "/data/apps/fibservi…"   About a minute ago   Up About a minute   0.0.0.0:8898-8899->8898-8899/tcp   nifty_bartik
```
