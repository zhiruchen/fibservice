FROM ubuntu:18.04

COPY bin/fibservice /data/apps/fibservice/bin/

EXPOSE 8898

ENTRYPOINT ["/data/apps/fibservice/bin/fibservice", "-p", ":8898"]

