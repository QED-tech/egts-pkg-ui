FROM nginx:alpine

WORKDIR /app

COPY ./infra/production/gw/nginx/conf.d /etc/nginx/conf.d
COPY ./infra/production/gw/nginx/snippets /etc/nginx/snippets

COPY ./frontend/dist/ /app/public

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]