FROM scratch
MAINTAINER Gustavo Gimenez <gimenezanderson@gmail.com>
ADD geo /
EXPOSE 4000

CMD ["./geo"]
