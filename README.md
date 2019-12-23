# go-http-rest-api

--------
##Docker-compose for two sql databases:

```version: "3.1"

volumes:
    db-data:

services:
    postgresql:
      image: postgres:alpine
      # restart: always
      container_name: Postgresql
      ports:
        - 5432:5432
      volumes:
        # - db-data:/var/lib/postgresql/data
        - ./pg-init-scripts:/docker-entrypoint-initdb.d
      environment:
        - POSTGRES_USER=db_user
        - POSTGRES_PASSWORD=db_user_pass
        # - POSTGRES_DB=restapi_dev
        - POSTGRES_MULTIPLE_DATABASES=restapi_dev,restapi_test

    # DB connection & admin
    adminer:
      image: adminer
      container_name: PG-adminer
      ports:
        - "8081:8080"
```

## Lib for migrations
```
https://github.com/golang-migrate/migrate
```