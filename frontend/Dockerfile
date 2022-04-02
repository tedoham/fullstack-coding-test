# Build stage
FROM node:17-alpine

# ENV NODE_ENV=production

WORKDIR /app
COPY package.json ./
COPY yarn.lock ./
COPY ./ ./

RUN yarn install 
RUN yarn build 

CMD ["yarn", "run", "dev"]