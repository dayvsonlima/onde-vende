FROM node:13-alpine

RUN mkdir /app

WORKDIR /app

COPY package.json /app/package.json

RUN npm install --silent

RUN npm install react-scripts -g --silent

EXPOSE 3000

CMD ["npm", "start"]
