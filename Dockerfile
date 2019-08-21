FROM onsdigital/dp-concourse-tools-ubuntu

WORKDIR /app/

COPY ./build/dp-frontend-renderer .
ADD taxonomy-redirects.yml .

ENTRYPOINT ./dp-frontend-renderer
