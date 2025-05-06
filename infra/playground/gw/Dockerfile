FROM nginx:alpine

WORKDIR /app

COPY ./infra/playground/gw/conf.d /etc/nginx/conf.d

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]