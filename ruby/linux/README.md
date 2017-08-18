# Gauge Ruby with Ruby 2.3 on Centos 7

## Contents

- [gauge-ruby23-centos](https://hub.docker.com/r/getgauge/gauge-ruby23-centos7/)
  - [Gauge] (https://github.com/getgauge/gauge/releases/latest)
  - Ruby - MRI Ruby 2.3
  - Centos 7
  - [Gauge-ruby plugin](https://github.com/getgauge/gauge-ruby/releases/latest)


## Usage

All images have the working directory set to `/gauge`. So ideally you should be able to clone your project in your host and run

```
docker run -it -v `pwd`:/gauge getgauge/gauge-ruby23-centos7 <gauge_command*>
```

<sup>*</sup> gauge command is specific to the project. Ex. java maven projects require `mvn test-compile gauge:execute`.

### Example

```
mkdir test && cd test
gauge init ruby # simulate a clone, by using gauge init on host.
docker pull getgauge/gauge-ruby23-centos7
docker run -it -v `pwd`:/gauge getgauge/gauge-ruby23-centos7 gauge run specs
```

## Reporting Issues

Comments on hub.docker.com are not actively monitored. Please report issues [here](https://github.com/getgauge-contrib/gauge-docker/issues/new).