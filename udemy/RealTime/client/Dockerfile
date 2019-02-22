FROM node:8
MAINTAINER Romel Campbell

ARG NODE_ENV=production
ARG PORT=3000
ENV NODE_ENV=${NODE_ENV}
ENV PORT=${PORT}


ENV APP_HOME /home/node/app
ADD . $APP_HOME/
WORKDIR $APP_HOME

RUN yarn install

EXPOSE 3000

CMD ["npm", "start"]
