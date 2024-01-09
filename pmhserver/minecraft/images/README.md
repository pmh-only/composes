# :package: `papermc-image`
Automatically build container images for minecraft PaperMC softwares.

*powered by Github Actions.

## Quick preview
Simply,
```sh
docker run -itp 25565:25565 -v .:/app ghcr.io/pmh-only/paper
```

or.. velocity proxy
```sh
docker run -itp 25565:25565 -v .:/app ghcr.io/pmh-only/velocity
```

jvm memory limit:
```sh
docker run -itp 25565:25565 -v .:/app ghcr.io/pmh-only/paper -Xms1G -Xmx1G
```

[Aikar's gc flag tuning](https://aikar.co/2018/07/02/tuning-the-jvm-g1gc-garbage-collector-flags-for-minecraft/) is enabled by default. You can disable this feature with `DISABLE_TUNING` environment variable.
```sh
docker run -itp 25565:25565 -e DISABLE_TUNING=true -v .:/app ghcr.io/pmh-only/paper -Xms1G -Xmx1G
```

## docker-compose
```yml
version: '3'

services:
  minecraft:
    image: ghcr.io/pmh-only/paper
    command: -Xms1G -Xmx1G
    restart: always
    user: 1000:1000
    tty: true
    stdin_open: true
    volumes:
      - .:/app:rw
    ports:
      - '25565:25565'
```

wanna update your paper bukkits? just type:
```sh
docker compose pull
docker compose up -d
```
