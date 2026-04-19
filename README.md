# Clamav-rest Signature Monitor

This application is meant to be a sidecar, a stand-alone container or a cronjob to monitor the signature age of [clamav-rest](https://github.com/ajilach/clamav-rest) since the project decided not to implement this feature in the core API.
Feel free to fork or use this alongside `clamav-rest` according to the [LICENSE](./LICENSE)

## Endpoint

* `GET /health/signature-age`
  * With the option of sending the query parameter `maxAgeHours`, `GET /health/signature-age?maxAgeHours=10`
  * Will return `200` if the signature age is less than `maxAgeHours``, otherwise`420`

## Environment variables

|  Environment variable  |  Description  |  Default value  |
| `CRONJOB`  | If this has a value other than an empty string, the HTTP server will not be started. The check will be run and then the container exits. Exit code 0 means success, exit code 1 means parsing the `MAX_SIGNATURE_AGE_HOURS` variable failed, exit code 2 means the signatures are too old.  |  Not set  |
|  `SIGMON_HTTP_PORT`  |  The port to run the HTTP server on  |  `9001`  |
|  `SIGMON_HTTPS_PORT`  | The port to run the HTTPS server on  |  not set, TLS disabled  |
|  `SSL_CERT` - Path to the TLS certificate  |  not set, TLS disabled  |
|  `SSL_KEY`  |  Path to the TLS private key  |  not set, TLS disabled  |
|  `CLAMAV_REST_URL`  |  URL to the clamav rest API  |  `http://localhost:9000`  |
|  `MAX_SIGNATURE_AGE_HOURS`  | The maximum amount of hours for the signature age before this API starts returning `420`  |  `26`  |
