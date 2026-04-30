# Clamav-rest Signature Monitor

This application is meant to be a sidecar, a stand-alone container or a cronjob to monitor the signature age of [clamav-rest](https://github.com/ajilach/clamav-rest) since the project decided not to implement this feature in the core API.
Feel free to fork or use this alongside `clamav-rest` according to the [LICENSE](./LICENSE)

## Nota Bene

* When clamav-rest restarts, it will start with the signatures from when the container image was created. After about two minutes, the signatures will be loaded. They are fetched on startup but need time to download and be consumed by clamav before they are fully loaded. During this time, clamav-rest-sigmon will return `420`.

## Endpoint

* `GET /health/signature-age`
  * With the option of sending the query parameter `maxAgeHours`, `GET /health/signature-age?maxAgeHours=10`. Setting this value overrides the env var `MAX_SIGNATURE_AGE_HOURS`.
  * Will return `200` if the signature age is less than `MAX_SIGNATURE_AGE_HOURS`, otherwise`420`

## Environment variables

|  Environment variable  |  Description  |  Default value  |
|------------------------|---------------|-----------------|
| `CRONJOB`  | If this has a value other than an empty string, the HTTP server will not be started. The check will be run and then the container exits. Exit code 0 means success, exit code 1 means parsing the `MAX_SIGNATURE_AGE_HOURS` variable failed, exit code 2 means the signatures are too old.  |  Not set  |  
|  `SIGMON_HTTP_PORT`  |  The port to run the HTTP server on  |  `9001`  |  
|  `SIGMON_HTTPS_PORT`  | The port to run the HTTPS server on  |  not set, TLS disabled  |  
|  `SSL_CERT`  |  Path to the TLS certificate  |  not set, TLS disabled  |  
|  `SSL_KEY`  |  Path to the TLS private key  |  not set, TLS disabled  |  
|  `CLAMAV_REST_URL`  |  URL to the clamav rest API  |  `http://localhost:9000`  |  
|  `MAX_SIGNATURE_AGE_HOURS`  | The maximum amount of hours for the signature age before this API starts returning `420`  |  `26`  |  

## Build

### Multi-platform build (amd64 and arm64)

`docker buildx build . -t docker.io/arizon/clamav-rest-sigmon:latest --platform linux/amd64,linux/arm64`

## Docker image

The docker image can be found on [Docker Hub](https://hub.docker.com/repository/docker/arizon/clamav-rest-sigmon).
