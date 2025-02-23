FROM ubuntu:latest

RUN apt update

RUN apt install -y nginx-extras libnginx-mod-http-lua luarocks

RUN luarocks install lua-resty-jwt && luarocks install lua-cjson && luarocks install lua-resty-string

COPY nginx.conf auth.lua /etc/nginx/

CMD ["nginx", "-g", "daemon off;"]
