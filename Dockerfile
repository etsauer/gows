FROM scratch

ADD gows /bin/gows

WORKDIR /opt

EXPOSE 8080

CMD ["/bin/gows"]
