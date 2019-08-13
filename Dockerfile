FROM onsdigital/dp-concourse-tools-ubuntu

WORKDIR /app/

COPY ./build/dp-frontend-renderer .
ADD ./assets/locales /app/assets/locales

ENTRYPOINT ./dp-frontend-renderer
