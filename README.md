# Setup for VPS Websites

setup.sh will install packages and setup the environment.

## Running

Go needs to be able to open a port, which requires elevated permissions. To run, use

```bash
sudo /usr/local/go/bin/go run .
```

## Running from docker

```docker
docker build -t vps.docker .
docker run -it -p 443:443 -p 80:80 vps.docker
```
