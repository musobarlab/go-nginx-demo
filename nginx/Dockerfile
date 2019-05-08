# nginx/Dockerfile
FROM nginx:latest

ARG SITE_DIR=/data/www

# create site folder
RUN mkdir -p ${SITE_DIR}

# copy site file to container site folder
COPY index.html ${SITE_DIR}/index.html

# create SSL folder
RUN mkdir -p /etc/nginx/ssl

# copy ssl files to container ssl folder
COPY ./ssl /etc/nginx/ssl/

EXPOSE 80
EXPOSE 443

COPY nginx.conf /etc/nginx/nginx.conf