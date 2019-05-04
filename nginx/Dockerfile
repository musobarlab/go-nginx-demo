#nginx/Dockerfile
FROM nginx:latest

ARG SITE_DIR=/data/www

RUN mkdir -p ${SITE_DIR}

COPY index.html ${SITE_DIR}/index.html

EXPOSE 80

COPY nginx.conf /etc/nginx/nginx.conf