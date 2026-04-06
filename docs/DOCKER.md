# Docker Guide

This guide shows the raw Docker commands for building, running, inspecting, and cleaning up the ASCII-Art-Web container without using `docker.sh`.

## Build the Image

Run the build from the project root:

```bash
docker image build -t ascii-art-web-docker .
```

The final `.` tells Docker to use the current directory as the build context.

## Run the Container

Start the web server in detached mode and publish port `8080`:

```bash
docker container run -d -p 8080:8080 --name dockerize ascii-art-web-docker
```

Open the application in your browser at `http://localhost:8080`.

## Check Images and Containers

```bash
docker images
docker ps -a
```

## Inspect the Container File System

Open a shell inside the running container:

```bash
docker exec -it dockerize /bin/bash
```

Then list the application files:

```bash
ls -l /app
```

Expected runtime content includes the `server` binary plus the `banners/`, `templates/`, and `static/` directories.

## Inspect Image Metadata

To view image labels and other metadata:

```bash
docker image inspect ascii-art-web-docker --format '{{json .Config.Labels}}'
```

## Stop and Remove the Container

```bash
docker stop dockerize
docker rm dockerize
```

## Optional Cleanup

Remove the image when you no longer need it:

```bash
docker image rm ascii-art-web-docker
```
