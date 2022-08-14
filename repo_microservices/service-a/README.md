# Simple Service A

Example Go application packaged with Docker and Helm.

## Create a multi-stage docker image

To compile and package using Docker multi-stage builds

```bash
docker build . -t service-a
```

## To run the docker image

```bash
docker run -p 8080:8080 service-a
```

And then visit http://localhost:8080 in your browser.
