# syntax=docker/dockerfile:1
# node project multistage build

FROM node:lts-alpine AS build

WORKDIR /app

COPY package.json .
COPY yarn.lock .

RUN yarn install --frozen-lockfile
COPY . .
RUN yarn build

FROM node:lts-alpine

WORKDIR /app

COPY --from=build /app .
COPY . .

EXPOSE 5173

CMD ["yarn", "start"]