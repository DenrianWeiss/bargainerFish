FROM ubuntu:latest
RUN apt-get update -y
COPY barginerFish /usr/local/bin/exe/barginerFish
COPY template /usr/local/bin/exe/template
COPY static /usr/local/bin/exe/static
COPY db /usr/local/bin/exe/db
RUN chmod +x /usr/local/bin/exe/barginerFish
WORKDIR /usr/local/bin/exe
ENTRYPOINT ["/usr/local/bin/exe/barginerFish"]