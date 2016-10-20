FROM ubuntu:16.04

WORKDIR /app/

COPY ./build/dp-frontend-renderer .

ENTRYPOINT ./dp-frontend-renderer
