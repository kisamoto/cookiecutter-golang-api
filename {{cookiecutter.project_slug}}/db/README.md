# Database Migrations

## Flyway

Migrations are managed manually by [Flyway](https://flywaydb.org/). During a release any schema migrations must be coordinated and run before a code deploy.

Migrations are merely ordered SQL files where Flyway keeps a record of already run files, detecting and applying any new SQL files in the `migrations` directory.

Configuration is managed in the `flyway.conf` file.

## Creating migration files

`touch migrations/V__$(date "+%Y%m%dT%H%M%S")-description.sql`

where `description` is a brief file name. 

e.g. `touch migrations/V__$(date "+%Y%m%dT%H%M%S")-initial_setup.sql`

## Running migrations

Make sure the `flyway` CLI tool is [installed on the system](https://flywaydb.org/documentation/commandline/). Double check the `flyway.conf` variables but these will be overwritten by CLI parameters.

```
$ cd {{cookiecutter.project_slug}}/db/
$ flyway 
```
