FROM nginx:alpine

COPY ./conf.d /etc/nginx/conf.d

EXPOSE 80 443

CMD ["nginx", "-g", "daemon off;"]