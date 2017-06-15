docker-build:
	@echo Building {{ cookiecutter.repo }}:build
	@echo
	docker build -t {{ cookiecutter.repo }}:build . -f Dockerfile.build
	@echo
	@echo Extracting api binary
	@echo
	docker create --name extract {{ cookiecutter.repo }}:build
	docker cp extract:/go/src/{{ cookiecutter.repo }}/api ./api
	docker rm -f extract
	@echo
	@echo Building {{ cookiecutter.repo }}:latest
	@echo
	docker build --no-cache -t {{ cookiecutter.repo }}:latest .
	@echo
	@echo Cleaning up
	@echo
	rm ./api
