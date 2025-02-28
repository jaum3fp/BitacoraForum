FROM node:23-alpine3.20

WORKDIR /opt/app

RUN npm install -g vite

CMD ["sh", "-c", "if [ ! -d node_modules ]; then npm install; fi; vite --host=0.0.0.0"]