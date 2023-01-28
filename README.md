# Go covid playground

## running on docker

- this will run on port 8080 and mongodb on port 27017 you can see mongodb URL in `cmd/infrastructure/datastore/db.go`

```bash
docker compose up -d
```

### stop docker

```bash
docker compose down --volume
```

| endpoint       | description                                      |
| -------------- | ------------------------------------------------ |
| /v1/covid      | fetch data from API Covid Thailand               |
| /v1/covid/get  | see all data from API in mongodb if it's working |
| /v1/covid/case | fetch covid case by year                         |
| /v1/covid/job  | fetch covid case by job                          |
