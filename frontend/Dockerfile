FROM node:lts-alpine as build
WORKDIR /app
COPY package*.json yarn.lock ./
RUN yarn
COPY . .
RUN yarn build

FROM node:lts-alpine AS prod
COPY --from=build /app/build .
COPY --from=build /app/package.json .
COPY --from=build /app/yarn.lock .
RUN yarn install --production
EXPOSE 3000
CMD ["node", "."]