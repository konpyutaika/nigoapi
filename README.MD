# NiGoAPI

## Purpose

The aim of this project is to provide a convenient Go client to interact with NiFi components ([NiFi](https://nifi.apache.org/) Api and [NiFi Registry](https://nifi.apache.org/registry.html) API)

## Documentations

- [NiFi client](pkg/registry/README.md)
- [NiFi Registry client](pkg/registry/README.md)

## Generate client

To generate the clients please run following command : 

```console
./regenerate.sh
```

## Credits

This client is : 

- inspired by the python client [nipyapi](https://github.com/Chaffelson/nipyapi)
- using [swagger-codegen](https://swagger.io/tools/swagger-codegen/)