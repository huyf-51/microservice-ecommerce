FROM ubuntu:latest

RUN apt update

RUN apt install -y nginx-extras libnginx-mod-http-lua luarocks

RUN luarocks install lua-resty-jwt && luarocks install lua-cjson && luarocks install lua-resty-string

COPY nginx.conf pub-key.pem auth.lua /etc/nginx/

RUN ln -sf /dev/stdout /var/log/nginx/access.log && \
    ln -sf /dev/stderr /var/log/nginx/error.log

CMD ["nginx", "-g", "daemon off;"]