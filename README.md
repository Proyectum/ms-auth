# ms-auth

## Generate API

```shell
oapi-codegen -package=api -generate "types,gin" resources/openapi.yml > internal/adapters/in/http/api/api.gen.go
```