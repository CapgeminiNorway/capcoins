FROM alpine:3.8
LABEL maintainer="Yilmaz Guleryuz 397125+zeusbaba@users.noreply.github.com"

# copy executable prepared via go build
COPY capcoins-api ./
RUN chmod +x ./capcoins-api

RUN mkdir -p config
RUN chmod -R +x config
COPY config/bot-dialect.json ./config/

EXPOSE 8088
CMD ./capcoins-api
