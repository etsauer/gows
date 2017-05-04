# gows - A Tiny Web Server

The gows project is a web server with a very small footprint that's made for simply serving static content.

## Build

First we're going to build the go binary.

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gows .
```

Then we build the image

```
docker build -t etsauer/gows:latest .
```

## Run

```
docker run -p 8080:8080 -v /path/to/static/content:/opt/_site -e GOWS_DIR=/opt/_site etsauer/gows:latest
```
