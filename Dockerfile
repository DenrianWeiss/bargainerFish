FROM ubuntu:latest
RUN mkdir /app
WORKDIR /app
COPY barginerFish barginerFish
COPY template ./template/
COPY static ./static/
COPY db ./db/
RUN chmod +x barginerFish
ENTRYPOINT ["/app/barginerFish"]