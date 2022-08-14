## Example microservices

This repository is meant to be an example of a mono-repo approach for developing a system with two microservices `service-a` and `service-b`.

For more details see [Developer Setup](../presentation/developer.md)

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
