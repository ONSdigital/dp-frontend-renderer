FROM onsdigital/dp-concourse-tools-ubuntu

WORKDIR /app/

COPY ./build/dp-frontend-renderer .

ENTRYPOINT ./dp-frontend-renderer
