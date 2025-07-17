# 🐳 Docker Complete Cheat Sheet

Your all-in-one guide to mastering Docker — from beginner basics to advanced usage.  
Perfect for developers, DevOps engineers, or anyone working with containers.

---

## 📘 Table of Contents

1. [Docker Overview](#docker-overview)  
2. [Docker Installation](#docker-installation)  
3. [Docker Basics](#docker-basics)  
4. [Managing Images](#managing-images)  
5. [Managing Containers](#managing-containers)  
6. [Volumes](#volumes)  
7. [Docker Networks](#docker-networks)  
8. [Dockerfile](#dockerfile)  
9. [Docker Compose](#docker-compose)  
10. [DockerHub & Image Sharing](#dockerhub--image-sharing)  
11. [Debugging & Logs](#debugging--logs)  
12. [Best Practices](#best-practices)  
13. [Common Issues & Fixes](#common-issues--fixes)

---

## 🚢 Docker Overview

Docker is a containerization platform that allows developers to package applications with all dependencies into a standardized unit.

---

## ⚙️ Docker Installation

- [Download Docker Desktop](https://www.docker.com/products/docker-desktop/)
- Install and verify:
  ```bash
  docker --version


```bash
## Docker basics

- docker build -t <image-name> .        # Build image from Dockerfile
- docker images                          # List local images
- docker run <image-name>               # Run container
- docker run -d -p 3000:3000 <image>    # Run in detached mode with port mapping
- docker ps                              # List running containers
- docker stop <container-id>            # Stop a container
- docker rm <container-id>              # Remove a container
- docker rmi <image-id>                 # Remove an image


## 🖼️ Managing Images

docker pull <image>                   # Pull from DockerHub
docker tag <local> <username>/<repo> # Tag before push
docker push <username>/<repo>        # Push to DockerHub
docker rmi <image-id>                # Remove image

## 💾 Volumes

docker volume create myvolume
docker run -v myvolume:/app/data ...

to  inspect

docker volume ls
docker volume inspect myvolume

## 🌐 Docker Networks

docker network create mynetwork
docker run --network=mynetwork ...
docker network ls

## 📄 Dockerfile

# Phase 1: Build
FROM node:18-alpine as builder
WORKDIR /app
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build

# Phase 2: Serve
FROM nginx:alpine
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80

## 📦 Docker Compose

version: '3.8'
services:
  frontend:
    build: ./frontend
    ports:
      - "3000:80"
  backend:
    build: ./backend
    ports:
      - "5000:5000"
    volumes:
      - ./data:/app/data



usage

docker-compose up --build
docker-compose down

## ☁️ DockerHub & Image Sharing

docker login
docker tag myapp yourname/myapp
docker push yourname/myapp

docker pull yourname/myapp
docker run -p 3000:80 yourname/myapp

## 🐞 Debugging & Logs

docker logs <container>
docker inspect <container>
docker exec -it <container> sh

## 🧠 Best Practices

- Use .dockerignore to skip unnecessary files
- Keep images small (use node:alpine)
- Use multistage builds
- Tag images with meaningful versions
- Clean unused containers/images with docker system prune




Made with ❤️ by [shaikh saim] — Happy Dockering!


