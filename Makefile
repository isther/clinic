run: down remove_image
	@docker-compose up -d

debug: down clean remove_image
	@docker-compose up -d

down:
	@docker-compose down

logs:
	@docker logs backend

clean:
	@sudo rm -rf data
	@sudo rm -rf logs

backend_image_name = "clinic_backend"
backend_image ="$(shell docker images | grep $(backend_image_name) | awk '{print $$1}')"
remove_image:

ifeq ($(backend_image),$(backend_image_name))
	@docker image rm $(backend_image_name)
endif

