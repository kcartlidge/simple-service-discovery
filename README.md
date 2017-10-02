# Simple-Service-Discovery

An easy to use service discovery tool and health check, based upon endpoints defined in an ```ssd.ini``` file:

``` ini
[SETTINGS]
# Number of seconds between health checks.
poll-seconds = 30
port = 1337

[ENDPOINTS]
# Billing site endpoints, including failover servers.
billing-test = http://billing-test.example.com:8080
billing-prod = https://billing1.example.com
billing-prod = https://billing2.example.com
billing-prod = https://billing3.example.com

# Authentication API endpoints, including failover servers.
auth-uat = http://my-auth.example.com:8080
auth-prod =  https://my-auth-1.example.com
auth-prod =  https://my-auth-2.example.com
```

## Features

* No service registration, just update the ```ssd.ini``` file to make changes.
* Polling system checks the status of the endpoints on a regular basis.
* Simple JSON responses for valid, invalid, or both in one.

## Status

Fine to use. The services API (```/services```) works as does the health check.

It is missing the monitoring of the ```ssd.ini``` file for changes, so you'll need to manually restart if the file is changed. This will be fixed shortly.

## How to use it

Populate the ```ssd.ini``` file and drop it on a server somewhere alongside the built executable of ```simple-service-discovery```. This is a single file and can be downloaded ready-built for *Linux*, *Windows* and *MacOS* from the ```builds``` subfolder.

Start it going.

*Optional: Point a DNS entry at the running instance so no knowledge of machines, IP addresses etc are required by clients.*

### Important note

If your endpoints are not fast-responding, please be aware that the polling frequency is based purely on time elapsed and does not allow for the time taken for the checks to be performed. In other words, ensure your frequency is longer than your anticipated checks duration, otherwise you risk overlapping checks.

It should be remembered that this is a service discovery tool, and the health check is an added bonus. Therefore if your use case is primarily as service discovery you probably don't need a high granularity frequency anyway (something like 300 seconds is probably fine).

### GET /services

This returns the endpoint(s) as an array of JSON object:

``` json
[
  {
    "name": "auth-prod",
    "endpoint": "https://my-auth-1.example.com",
    "status": 200
  },
  {
    "name": "auth-prod",
    "endpoint": "https://my-auth-2.example.com",
    "status": 200
  }
]
```

All endpoints with the given key name are returned, along with their last known HTTP status code.

---

## Building and running

### One-off build and run

``` sh
cd <project-root>
go build -o builds/linux/ssd && ./builds/linux/ssd # Linux
go build -o builds/macos/ssd && ./builds/macos/ssd # Mac OS
go build -o builds\windows\ssd.exe && .\builds\windows\ssd.exe # Windows
```

### Build all platform executables

These are tiny, and so are committed to the repository in the ```builds``` folder.

``` sh
cd <project-root>
./make/make-linux.sh && ./builds/linux/ssd # Linux
./make/make-macos.sh && ./builds/macos/ssd # Mac OS
./make/make-windows.bat && .\builds\windows\ssd.exe # Windows
```
