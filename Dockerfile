FROM busybox
MAINTAINER Gustavo Gimenez <gimenezanderson@gmail.com>
RUN mkdir /app
ADD build/geo /app
RUN mkdir /app/data
RUN chmod -R 777 /app
EXPOSE 4000

CMD ["./app/geo"]
