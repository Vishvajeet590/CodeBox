FROM node:lts as dependencies
WORKDIR /codebox-web
COPY package.json ./
RUN npm install

FROM node:lts as builder
WORKDIR /codebox-web
COPY . .
COPY --from=dependencies /codebox-web/node_modules ./node_modules
RUN npm run build

FROM node:lts as runner
WORKDIR /codebox-web
ENV NODE_ENV production

COPY --from=builder /codebox-web/public ./public
COPY --from=builder /codebox-web/.next ./.next
COPY --from=builder /codebox-web/node_modules ./node_modules
COPY --from=builder /codebox-web/package.json ./package.json
COPY --from=builder /codebox-web/next.config.js ./
COPY --from=builder /codebox-web/tailwind.config.js ./

EXPOSE 3000
CMD ["npm", "run","start"]
