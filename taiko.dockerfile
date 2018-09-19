FROM node

LABEL taiko.version="0.2.0"\
    gauge.version="1.0.1"\
    description="debian image with node, dependencies for headless chrome, gauge and taiko npm packages"\
    repository="getgauge/taiko"


ENV NPM_CONFIG_PREFIX=/home/node/.npm-global
ENV PATH=$PATH:/home/node/.npm-global/bin

# requirements for chrome headless
# https://github.com/Googlechrome/puppeteer/issues/290#issuecomment-322838700
RUN apt-get update &&\
  apt-get install -y git-all gconf-service libasound2 libatk1.0-0 libc6 libcairo2 libcups2 libdbus-1-3 \
    libexpat1 libfontconfig1 libgcc1 libgconf-2-4 libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 \
    libnspr4 libpango-1.0-0 libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 \
    libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 libxss1 libxtst6 ca-certificates \
    fonts-liberation libappindicator1 libnss3 lsb-release xdg-utils wget &&\
    rm -rf /var/lib/apt/lists/* &&\
    mkdir -p /gauge &&\
    chown -R node:node /gauge &&\
    chmod 755 /gauge

USER node

RUN npm install -g taiko @getgauge/cli

WORKDIR /gauge