FROM postgres:alpine
LABEL authors="Dmitriy Panin"

ENV POSTGRES_USER=admin
ENV POSTGRES_PASSWORD=admin_password

COPY ./setupDB.sh /docker-entrypoint-initdb.d/

EXPOSE 5432
