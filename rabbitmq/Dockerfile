FROM docker.io/rabbitmq:3.7.16-management-alpine

COPY plugins/ plugins/

RUN rabbitmq-plugins enable rabbitmq_message_deduplication 