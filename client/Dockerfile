FROM node:23-alpine

WORKDIR /opt/app

RUN apk add --no-cache curl bash
RUN curl -fsSL https://bun.sh/install | bash
RUN apk del curl bash
RUN ln -s ~/.bun/bin/bun /usr/local/bin/bun


CMD [ "ash", "-c", "if [ ! -d node_modules ]; then bun install; fi; bun x nuxi dev --host=0.0.0.0" ]