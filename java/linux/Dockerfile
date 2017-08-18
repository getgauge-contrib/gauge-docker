# Copyright (c) ThoughtWorks Inc. All rights reserved.
# Licensed under the MIT license. See LICENSE file in the project root for details.

FROM centos

ARG MAVEN_VERSION=3.5.0
ARG USER_HOME_DIR="/root"
ARG SHA=beb91419245395bd69a4a6edad5ca3ec1a8b64e41457672dc687c173a495f034
ARG BASE_URL=http://redrockdigimark.com/apachemirror/maven/maven-3/${MAVEN_VERSION}/binaries

WORKDIR /gauge

USER root

RUN \
  wget -O /tmp/gauge.zip `curl -s https://api.github.com/repos/getgauge/gauge/releases | grep browser_download_url | grep 'linux.x86_64.zip' | head -n 1 | cut -d '"' -f 4` && \
  unzip -d /tmp/gauge /tmp/gauge.zip && \
  cp /tmp/gauge/bin/gauge /usr/local/bin/ && \
  rm -rf /tmp/gauge /tmp/gauge.zip ~/.gauge/config ~/.gauge/logs && \
  yum install -y java-1.8.0-openjdk java-1.8.0-openjdk-devel which && \ 
  mkdir -p /usr/share/maven /usr/share/maven/ref \
  && curl -fsSL -o /tmp/apache-maven.tar.gz ${BASE_URL}/apache-maven-${MAVEN_VERSION}-bin.tar.gz \
  && echo "${SHA}  /tmp/apache-maven.tar.gz" | sha256sum -c - \
  && tar -xzf /tmp/apache-maven.tar.gz -C /usr/share/maven --strip-components=1 \
  && rm -f /tmp/apache-maven.tar.gz \
  && ln -s /usr/share/maven/bin/mvn /usr/bin/mvn && \
  rm -rf gauge gauge.zip ~/.gauge/config ~/.gauge/logs && \
  yum clean all && \
  chmod 777 /gauge

USER default

RUN gauge install java

ENV JAVA_HOME /usr/lib/jvm/java-1.8.0-openjdk/
ENV MAVEN_HOME /usr/share/maven
ENV MAVEN_CONFIG "$USER_HOME_DIR/.m2"

