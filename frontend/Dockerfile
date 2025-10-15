FROM node:20-alpine AS deps

WORKDIR /app

COPY package.json packag-lock.json* ./

RUN npm ci

FROM node:20-alpine AS builder

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules

COPY . .

RUN npm run build

FROM nginx:alpine AS production

COPY --from=builder /app/dist /usr/share/nginx/html

COPY docker/nginx/nginx.conf /etc/nginx/nginx.conf
COPY docker/nginx/conf.d/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

FROM node:20-alpine AS development

WORKDIR /app

COPY package.json package-lock.json* ./

RUN npm ci

COPY . .

EXPOSE 3000

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]