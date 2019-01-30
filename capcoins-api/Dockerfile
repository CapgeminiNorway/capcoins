FROM alpine:3.8
LABEL maintainer="Yilmaz Guleryuz 397125+zeusbaba@users.noreply.github.com"

# copy executable prepared via go build
COPY capcoins-api ./
RUN chmod +x ./capcoins-api

EXPOSE 8088
CMD ./capcoins-api
