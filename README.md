# Gauge Docker

Docker images with Gauge installed.

## Platforms

### Linux

Linux images are centos based. 

- [gauge-openjdk8-centos](link_tbd) 
  - Gauge (latest)
  - JDK - OpenJDK 8
  - Centos 7
  - Gauge-java plugin

- [gauge-ruby23-centos](link_tbd) 
  - Gauge (latest)
  - Ruby 2.3
  - Centos 7
  - Gauge-ruby plugin

- [gauge-mono-centos](link_tbd) 
  - Gauge (latest)
  - Mono 4.8 (gauge-csharp does not support mono 5.x yet)
  - Centos 7
  - Gauge-csharp plugin


### Windows

TODO

## Usage

All images have the working directory set to `/gauge`. So ideally you should be able to clone your project in your host and run

```
docker run -it -v `pwd`:/gauge getgauge/<image_name> <gauge_command*>
```

<sup>*</sup> gauge command is specific to the project. Ex. maven projects require `mvn test-compile gauge:execute`.

## License

MIT License. See [LICENSE](LICENSE).