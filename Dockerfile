FROM alpine:latest
RUN mkdir /app

COPY ./subs-service.bin /app

# Run the server executable
CMD [ "/app/subs-service.bin" ]