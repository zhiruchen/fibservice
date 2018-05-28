FROM ubuntu:18.04

COPY bin/fibservice /data/apps/fibservice/bin/

EXPOSE 8898 8899

CMD ["-p", ":8898", "-pp", ":8899"]
ENTRYPOINT ["/data/apps/fibservice/bin/fibservice"]

