# FoldingAtHome Prometheus Exporter

Simple exporter to get infromation from Folding at Home API for specified users

## Config file
Config file is called `config.yaml` and should be in the same folder with executable file.

There is available three variables:
```
users:
  - tzlblsvo3dov
  - t1bur1an
sleep: 5 #minutes
port: 2112
```
List of users, delay between api calls and port number. Pretty simple. 

## Building

Run `go build -o fah-exporter`

## Docker

Project contains Dockerfile so you can simply build it by command: `docker build -t fah-exporter:1.0 .` and run it `docker run -d -p 2112:2112 fah-exporter:1.0`
