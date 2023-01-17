FROM onsdigital/dp-concourse-tools-ubuntu-20:ubuntu20.4-rc.1

WORKDIR /app/

COPY ./build/dp-frontend-renderer .
ADD taxonomy-redirects.yml .

ENTRYPOINT ./dp-frontend-renderer
