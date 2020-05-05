FROM node:14-alpine
WORKDIR /usr/src/app
RUN chmod 777 /usr/src/app
RUN \
    # Add edge repo
    echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/main' >> /etc/apk/repositories && \
    echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories && \
    echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories 

    # Update nodejs
RUN apk add --update --no-cache nodejs nodejs-npm
    
    # Update packages
RUN apk --no-cache upgrade

    # Install Bash
RUN apk --no-cache add bash

    # Install libx11
RUN apk --no-cache add libx11 
    # Install SSL
    # Alpine 3.5 switched from OpenSSL to LibreSSL
RUN apk --no-cache add libressl

RUN apk --no-cache add libc6-compat 

RUN apk add --no-cache --update --virtual .build-deps \	
      build-base \	
      libffi-dev \	
      openssl-dev

RUN apk add --no-cache \
      chromium \
      nss \
      freetype \
      freetype-dev \
      harfbuzz \
      ca-certificates \
      ttf-freefont \
      nodejs \
      yarn

#Tell Puppeteer to skip installing Chrome. We'll be using the installed package.
ENV PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium-browser
COPY package*.json ./
COPY . .
RUN npm install

