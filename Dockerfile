FROM ubuntu:latest
COPY barginerFish barginerFish
COPY template ./template/
COPY static ./static/
COPY db ./db/
RUN chmod +x barginerFish
ENTRYPOINT ["barginerFish"]