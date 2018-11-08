## Golang+postgres and Docker quick start

### The application consists of:
* golang
** glide
** httprouter
** godotenv
* postgres 10.4


### Getting started:
* install [docker](https://docs.docker.com/install/) if not already installed
* install [docker-compose](https://docs.docker.com/compose/install/) if not already installed

## Install GIT
```bash
apt-get update && apt-get install git
```

### Run docker container

```bash
docker-compose up --build
```

To get container name run follow command
```bash
docker-compose ps | grep <service_name> | awk '{ print $1 }'
```

To exec commands inside containers use `docker exec -i -t <container_name> bash`

For example to install new API dependency:
```bash
docker exec -i -t dev_api bash
glide get <package_name>
```

- API host [http://api.mercadolibre.com](http://api.mercadolibre.com)

To get access via domain add to /etc/hosts
```
127.0.0.1    api.mercadolibre.com
```
