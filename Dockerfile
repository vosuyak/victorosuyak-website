FROM node:12.14.1 AS build
RUN mkdir /app
ADD /frontend /app
WORKDIR /app
RUN ls
# COPY package.json package-lock.json ./
RUN npm install
RUN npm audit fix
COPY . .
RUN npm run build --prod
