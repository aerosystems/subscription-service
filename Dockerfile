FROM alpine:latest
RUN mkdir /app

COPY ./subscription-service.bin /app

# Run the server executable
CMD [ "/app/subscription-service.bin" ]