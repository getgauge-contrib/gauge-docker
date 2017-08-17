# Copyright (c) ThoughtWorks Inc. All rights reserved.
# Licensed under the MIT license. See LICENSE file in the project root for details.

FROM centos

WORKDIR /gauge

USER root

RUN \
  wget -O /tmp/gauge.zip `curl -s https://api.github.com/repos/getgauge/gauge/releases | grep browser_download_url | grep 'linux.x86_64.zip' | head -n 1 | cut -d '"' -f 4` && \
  unzip -d /tmp/gauge /tmp/gauge.zip && \
  cp /tmp/gauge/bin/gauge /usr/local/bin/ && \
  # http://www.mono-project.com/download/#download-lin-centos
  yum install yum-utils && \
  rpm --import "http://keyserver.ubuntu.com/pks/lookup?op=get&search=0x3FA7E0328081BFF6A14DA29AA6A19B38D3D831EF" && \
  yum-config-manager --add-repo http://download.mono-project.com/repo/centos7/ && \
  yum install mono-devel-4.8.0 &&\
  yum uninstall yum-utils && \
  yum clean all && \
  rm -rf gauge gauge.zip ~/.gauge/config ~/.gauge/logs &&\
  chmod 777 /gauge

USER default

RUN gauge install csharp