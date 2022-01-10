# Go Mantil Sample Project

Retrieves data from meteoprovider and send it to a specified email address.

## Configuration
`configuration/environment.yml` contains required env var to get started with the project.

## Deployment via Mantil

### Infrastructure creation
* `mantil aws install --aws-profile=$PROFILE`
* `mantil stage new development`
* `mantil stage new production`

### Deployment
`mantil deploy --stage development`
`mantil deploy --stage production`


## Invocation
`mantil invoke wheater/get --stage production --data '{"stationId":"10637"}'`

`mantil invoke wheater/get --stage development --data '{"stationId":"10637"}'`