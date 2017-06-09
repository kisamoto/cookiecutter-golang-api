# Cookiecutter Go API

An opinionated template to create Go powered API's. Whilst this is primarily
for how _I_ create projects I am always looking for improvements and other 
options - Pull requests welcome. 

The aim here is to have the right balance between speed (both development
 and runtime), flexibility and modularity. 

## Usage

With [`cookiecutter`](https://github.com/audreyr/cookiecutter) installed, run:

```
$ cookiecutter gh:kisamoto/cookiecutter-golang-api
```

## More info

More information is found in the internal [`README`]({{cookiecutter.project_slug/README.md}})

## ToDo

- [] `Dockerfile` - 
- [] `docker-compose.yml`
- [] `Caddyfile` - Configured without Let's Encrypt for development usage but as close to production setup as possible
- [] `Makefile` - Should include commands to build API; deploy with Docker; setup cluster with docker-compose; unit test API; functionally test API
