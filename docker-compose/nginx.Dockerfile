FROM nginx
COPY html /usr/share/nginx/html

COPY server.conf /etc/nginx/conf.d/default.conf
