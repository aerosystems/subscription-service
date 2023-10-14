FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/logs

COPY ./lookup-service/lookup-service.bin /app

# Run the server executable
CMD [ "/app/lookup-service.bin" ]