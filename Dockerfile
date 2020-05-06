FROM node:14-alpine
WORKDIR /usr/src/app
RUN chmod 777 /usr/src/app
RUN \
    # Add edge repo
    echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/main' >> /etc/apk/repositories \
    echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/community' >> /etc/apk/repositories \
    echo '@edge http://dl-cdn.alpinelinux.org/alpine/edge/testing' >> /etc/apk/repositories \

    # Update nodejs
    apk add --update --no-cache nodejs nodejs-npm \
    
    # Update packages
    apk --no-cache upgrade \

    # Install Bash
    apk --no-cache add bash \

    # Install libx11
    apk --no-cache add libx11 \ 
    # Install SSL
    # Alpine 3.5 switched from OpenSSL to LibreSSL
    apk --no-cache add libressl \

    apk --no-cache add libc6-compat \ 

    apk add --no-cache --update --virtual .build-deps \	
      build-base \	
      libffi-dev \	
      openssl-dev

COPY package*.json ./
COPY . .
RUN npm install

