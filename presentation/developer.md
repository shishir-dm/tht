## Developer Setup

In order to enable developers to test their changes locally we propose using docker-compose.

Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. To learn more about all the features of Compose, see the list of features.

## How to?

In this example we assume that service-a and service-b are microservices and after understanding the tradeoff's we have chosen a [mono-repo](repo_microservices) approach, optionally we have also provided the skeleton for a multi-repo approach of [service-a](repo_service-a) and [service-b](repo_service-b) individually.

- Since we have chosen a mono-repo approach we have stored the docker-compose.yaml file in the root of the repo.
- This allows us to easily spin up the entire application i.e service-a, service-b and cockroach-db locally via a single command and optionally add profiles with which we can optionally spin-up only a set of required application and not all 3 (e.g only service-b and the db)

Example docker-compose file
```yaml
version: "3.8"
services:
  service-a:
    build: "service-a"
    ports:
      - "8080"
    restart: unless-stopped
    environment:
      DB_URL: cockroachdb:26257
      DB_USER: root
      DB_PASSWORD_FILE: /run/secrets/db-password
    networks:
      - backend

  service-b:
    build: "service-b"
    ports:
      - "8081"
    restart: unless-stopped
    environment:
      DB_URL: cockroachdb:26257
      DB_USER: root
      DB_PASSWORD_FILE: /run/secrets/db-password
    networks:
      - backend

  cockroachdb:
    image: cockroachdb/cockroach:v19.2.2
    ports:
      - "26257:26257"
    command: start-single-node --insecure
    secrets:
      - db-password
    environment:
      ROOT_PASSWORD_FILE: /run/secrets/db-password
    volumes:
      - "db-data:/cockroach/cockroach-data"
    networks:
      - backend

  nginx:
    build: "./hacks/nginx"
    ports:
      - "80:80"
    depends_on:
      - "cockroachdb"
      - "service-a"
      - "service-b"
    networks:
      - backend

networks:
  backend: {}
volumes:
  db-data: {}
secrets:
  db-password:
    file: hacks/db/password.txt
```

To start the dev setup ensure you have the following installed
- Docker Desktop > v4.6
- docker cli > v20.10.13
- docker-compose > v1.29.2

To start the dev environment with [docker-compose](repo_microservices/docker-compose.yaml), run the following:
1. Clone this repo with `git clone git@github.com:shishir-dm/tht.git`
2. Navigate to `repo_microservices`
   ```
   cd tht/repo_microservices
   ```
3. Run the following docker-compose commands in order
   ```
   docker-compose build
   docker-compose up -d
   ```
4. If this runs successfully you should be able to see the running docker containers with `docker ps`. And in your browser navigating to `http://localhost/servicea` or `http://localhost/serviceb` should return meaningful responses
5. To shut-down everything run `docker-compose down -v`

### To continue the presentation, head to the next section: [Compliance](compliance.md)