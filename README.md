# atlassian.carcgl.com/bitbucket/ls/lms

Licensing service - LMS(learning management system) user info micro-service
This is a micro service response for connecting the LMS system and fetch operator info
This micro service provide a

# always run buf mod update after adding a dependency to your buf.yaml.

buf mod update

# run postgresql docker container

docker run -d --name postgres-1 -e POSTGRES_DB=lego_db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -p 5432:5432 postgres:15.1-alpine

docker run --name postgres-1 --rm -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e PGDATA=/var/lib/postgresql/data/pgdata -v /tmp:/var/lib/postgresql/data -p 5432:5432 -it postgres:15.1-alpine

# exec to postgresql

docker exec -it postgres-1 bash

# login to database

psql --username=postgres --dbname=postgres

# show connect info

\c

# show table list

\dt

# show content in "operators" table

SELECT \* FROM operators;
SELECT id FROM operators;

# delete complete data from an existing table

TRUNCATE TABLE operators;

# drop table

DROP TABLE operators

# exit

\q

# Build Docker image / push image to docker hub

build
docker build -f docker/lms.Dockerfile -t lms-ms:0.0.1 .

run
docker run -it --name lms-ms lms-ms:0.0.1

run (with environment variable)
docker run -it -e DB_ADDR=172.17.0.2 --name lms-ms -p9094:9094 lms-ms:0.0.1

stop
docker stop lms-ms

check container
docker inspect lms-ms | grep IPAddress
docker inspect postgres-1 | grep IPAddress

clean builder layer image
docker image prune --filter label=stage=builder
