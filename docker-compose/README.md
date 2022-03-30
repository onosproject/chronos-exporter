<!--
SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>

SPDX-License-Identifier: Apache-2.0
-->
# RASA Web Chat demonstrator

This `docker-compose` distribution allows running the RASA
model locally and make a page available at:

* http://localhost:8181

Docker Compose [reference](https://docs.docker.com/compose/).

## Running it

To run it do:

```bash
cd docker-compose
docker-compose up
```

Use `Ctrl-C` to stop the distribution.

If you prefer to leave running in the background (as a daemon) use:
```bash
docker-compose up -d
```
and it can be stopped with
```bash
docker-compose down
```

The compose file consists of 3 containers

1. the `rasa-model-server` - our own docker image that trains 
   the local model and saves it to a TGZ in a docker image.
2. the `rasa` server that runs the model and the websocket.
   The REST API is also enabled.
3. The `nginx` server - serves the static `index.html` file 
   and also does reverse proxy pass for RASA

## Rebuilding
The Docker images are built dynamically. If you need to rebuild, run:
```bash
docker-compose build
docker-compose up
```
Note: if you get an error about `go: inconsistent vendoring` then you should `make images` again from top level directory.

## Monitoring
From another terminal window (when in the same `docker-compose` directory)
you can monitor the status with:
```bash
docker-compose ps
```
