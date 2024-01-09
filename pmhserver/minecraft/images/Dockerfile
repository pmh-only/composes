FROM alpine as builder

RUN apk add --no-cache wget

COPY ./start.sh /bin/start.sh

ARG PROJECT_NAME
ARG VERSION_NAME
ARG BUILD_ID
ARG DOWNLOAD_NAME

RUN wget https://api.papermc.io/v2/projects/${PROJECT_NAME}/versions/${VERSION_NAME}/builds/${BUILD_ID}/downloads/${DOWNLOAD_NAME} -O /bin/server.jar

RUN chmod a+rwx /bin/start.sh /bin/server.jar

FROM alpine as runner

RUN apk add --no-cache openjdk17-jre

WORKDIR /app

VOLUME ["/app"]

EXPOSE 25565

COPY --from=builder /bin/server.jar /bin/server.jar
COPY --from=builder /bin/start.sh /bin/start.sh

ENTRYPOINT ["/bin/start.sh"]
