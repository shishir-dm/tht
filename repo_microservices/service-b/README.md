# Simple Service B

Example Go application packaged with Docker and Helm.

## Create a multi-stage docker image

To compile and package using Docker multi-stage builds

```bash
docker build . -t service-b
```

## To run the docker image

```bash
docker run -p 8080:8080 service-b
```

And then visit http://localhost:8081 in your browser.
