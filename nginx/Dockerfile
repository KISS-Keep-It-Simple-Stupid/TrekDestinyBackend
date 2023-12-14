FROM nginx:1.16.0-alpine

RUN rm /etc/nginx/conf.d/default.conf
RUN apk update
RUN apk add openssl
COPY nginx.conf /etc/nginx/conf.d
COPY ./back_cert.pem /etc/ssl
COPY ./back_key.key /etc/ssl
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
