FROM alpine:3.15

RUN apk add npm

COPY src/vue vue
COPY docker/ui/.env.playground vue/.env

WORKDIR vue

RUN npm install

CMD ["npm", "run", "serve"]
