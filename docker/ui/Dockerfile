FROM alpine:3.15 as builder

RUN apk add npm

COPY src/vue vue
COPY docker/ui/.env.playground vue/.env

WORKDIR vue

RUN npm install

CMD ["npm", "run", "build"]

FROM alpine:3.15

RUN apk add nginx

COPY --from=builder /vue/dist /var/www/
COPY docker/ui/etc/nginx/default.conf /etc/nginx/http.d/default.conf

EXPOSE 80

STOPSIGNAL SIGQUIT

CMD ["nginx", "-g", "daemon off;"]
