FROM node:alpine

WORKDIR /app

COPY package.json package-lock.json .

RUN npm install --omit=dev

COPY . .

CMD ["npm", "start"]
