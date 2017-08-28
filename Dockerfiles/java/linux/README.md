# Gauge Java with OpenJDK 1.8 on Centos 7

## Contents

- [gauge-jdk8-centos](https://hub.docker.com/r/getgauge/gauge-jdk8-centos7/)
  - [Gauge] (https://github.com/getgauge/gauge/releases/latest)
  - JDK - OpenJDK 8
  - Centos 7
  - [Gauge-java plugin](https://github.com/getgauge/gauge-java/releases/latest)


## Usage

All images have the working directory set to `/gauge`. So ideally you should be able to clone your project in your host and run

```
docker run -it -v `pwd`:/gauge getgauge/gauge-jdk8-centos7 <gauge_command*>
```

<sup>*</sup> gauge command is specific to the project. Ex. java maven projects require `mvn test-compile gauge:execute`.

### Example

```
mkdir test && cd test
gauge init java # simulate a clone, by using gauge init on host.
docker pull getgauge/gauge-jdk8-centos7
docker run -it -v `pwd`:/gauge getgauge/gauge-jdk8-centos7 gauge run specs
```

## Reporting Issues

Comments on hub.docker.com are not actively monitored. Please report issues [here](https://github.com/getgauge-contrib/gauge-docker/issues/new).