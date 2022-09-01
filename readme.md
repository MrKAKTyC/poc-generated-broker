# Pre requirements

installed kubernetes cluster, [service catalog](https://github.com/kubernetes-retired/service-catalog)

install [oapi generator](https://github.com/deepmap/oapi-codegen)

download [specification file](https://github.com/openservicebrokerapi/servicebroker/blob/master/openapi.yaml)

# Code generation

To generate a server stub use the following command

`oapi-codegen --package [package-name] -generate server --old-config-style openapi.yaml > serv.go`

and for generating structures according to specs

`oapi-codegen --package [package-name] -generate types --old-config-style openapi.yaml > model.go`

Move 

# Adding Broker to the service catalog

To add a broker to the service catalog use the following command

`svcat register [broker-name] --url [broker-url]`

To check the status and available classes and plan use the following command

`svcat get brokers`

`svcat get classes`

`svcat get plans`