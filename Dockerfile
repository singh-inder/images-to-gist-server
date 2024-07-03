FROM node:21.7.3-bullseye-slim as builder

WORKDIR /app

COPY package*.json .

COPY . .

RUN npm ci

RUN npm run build

FROM node:21.7.3-bullseye-slim as final

WORKDIR /app

COPY --from=builder /app/dist dist

COPY package*.json .

RUN npm ci --omit=dev

CMD [ "node","dist/server.js" ]