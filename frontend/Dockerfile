FROM node:12.14.1-alpine

WORKDIR /usr/src/frontend

RUN apk add git

COPY package*.json ./

RUN npm install

COPY . ./

CMD [ "npm", "run", "start" ]
