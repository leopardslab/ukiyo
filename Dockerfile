FROM alpine

COPY ukiyo .
COPY dbs .
RUN chmod 777 ./ukiyo

EXPOSE 8080

CMD ["./ukiyo"]
