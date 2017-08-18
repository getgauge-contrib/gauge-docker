# Gauge C# with Mono 4.8 on Centos 7

## Contents

- [gauge-mono48-centos]((https://hub.docker.com/r/getgauge/gauge-mono48-centos7/))
  - [Gauge] (https://github.com/getgauge/gauge/releases/latest)
  - Mono 4.8 (gauge-csharp does not support mono 5.x yet)
  - Centos 7
  - [Gauge-csharp plugin](https://github.com/getgauge/gauge-csharp/releases/latest)

## Usage

All images have the working directory set to `/gauge`. So ideally you should be able to clone your project in your host and run

```
docker run -it -v `pwd`:/gauge getgauge/gauge-mono48-centos7 <gauge_command*>
```

<sup>*</sup> gauge command is specific to the project. Ex. java maven projects require `mvn test-compile gauge:execute`.

### Example

```
mkdir test && cd test
gauge init csharp # simulate a clone, by using gauge init on host.
docker pull getgauge/gauge-mono48-centos7
docker run -it -v `pwd`:/gauge getgauge/gauge-mono48-centos7 gauge run specs
```

## Reporting Issues

Comments on hub.docker.com are not actively monitored. Please report issues [here](https://github.com/getgauge-contrib/gauge-docker/issues/new).