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
