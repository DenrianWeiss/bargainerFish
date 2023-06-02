FROM ubuntu:latest
WORKDIR /usr/local/bin/exe
COPY barginerFish barginerFish
COPY template ./template/
COPY static ./static/
COPY db ./db/
RUN chmod +x barginerFish
ENTRYPOINT ["/usr/local/bin/exe/barginerFish"]