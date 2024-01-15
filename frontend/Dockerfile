# syntax=docker/dockerfile:1
# node project multistage build

FROM node:lts-alpine as build
WORKDIR /app
COPY package*.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn build

FROM node:lts-alpine AS prod

# copy only the necessary files
COPY --from=build /app/build ./build
COPY --from=build /app/node_modules ./node_modules
COPY --from=build /app/package*.json /app/yarn.lock ./

EXPOSE 3000

CMD ["yarn", "dev"]