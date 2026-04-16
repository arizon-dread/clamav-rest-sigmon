# Clamav-rest Signature Monitor

This API is meant to be a sidecar or stand-alone container to monitor the signature age of [clamav-rest](https://github.com/ajilach/clamav-rest) since the project decided not to implement this feature in the core API.
Feel free to fork or use this alongside `clamav-rest` according to the [LICENSE](./LICENSE)

## Endpoint

* `GET /health/signature-age`
  * With the option of sending the query parameter `maxAgeHours`, `GET /health/signature-age?maxAgeHours=10`
  * Will return `200` if the signature age is less than `maxAgeHours``, otherwise`420`

## Environment variables

`SIGMON_HTTP_PORT` - The port to run the HTTP server on, default: `9001`
`SIGMON_HTTPS_PORT` - The port to run the HTTPS server on, default: TLS disabled
`SSL_CERT` - Path to the TLS certificate, default: TLS disabled
`SSL_KEY` - Path to the TLS private key, default: TLS disabled
`CLAMAV_REST_URL` - URL to the clamav rest API, default: `http://localhost:9000`
`MAX_SIGNATURE_AGE_HOURS` - The maximum amount of hours for the signature age before this API starts returning `420`, default: `26`
