FROM onsdigital/dp-concourse-tools-ubuntu

WORKDIR /app/

COPY ./build/dp-frontend-renderer .
COPY ./assets/locales /app/assets/locales
ADD taxonomy-redirects.yml .

ENTRYPOINT ./dp-frontend-renderer
