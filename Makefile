# Variables
IMAGE_NAME = afya-chain-db
CONTAINER_NAME = afya-chain-db

# Build the Docker image
build:
	docker build -t $(IMAGE_NAME) .

# Run the Docker container
run:
	docker run --name $(CONTAINER_NAME) -d -p 3306:3306 $(IMAGE_NAME)
	go run .

# Stop the Docker container
stop:
	docker stop $(CONTAINER_NAME)

# Remove the Docker container
rm:
	docker rm $(CONTAINER_NAME)

# Remove the Docker image
rmi:
	docker rmi $(IMAGE_NAME)

# Clean up: stop and remove the container
cleanup: stop rm

# Show the status of the Docker container
status:
	docker ps -a | grep $(CONTAINER_NAME)

# Help: show available commands
help:
	@echo "Available commands:"
	@echo "  make build       - Build the Docker image"
	@echo "  make run         - Run the Docker container"
	@echo "  make stop        - Stop the Docker container"
	@echo "  make rm          - Remove the Docker container"
	@echo "  make rmi         - Remove the Docker image"
	@echo "  make cleanup      - Stop and remove the container"
	@echo "  make status      - Show the status of the Docker container"
	@echo "  make help        - Show this help message"
