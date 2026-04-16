# Clamav-rest Signature Monitor

This API is meant to be a sidecar or stand-alone container to monitor the signature age of [clamav-rest](https://github.com/ajilach/clamav-rest) since the project decided not to implement this feature in the core API.
Feel free to fork or use this alongside `clamav-rest` according to the [LICENSE](./LICENSE)

## Endpoint

* `GET /health/signature-age`
  * With the option of sending the query parameter `maxAgeHours`, `GET /health/signature-age?maxAgeHours=10`
  * Will return `200` if the signature age is less than `maxAgeHours``, otherwise`420`
