GO_CHECKS_APP=goChecksApp

up: compile
	@echo "Stopping pods (if running)..."
	podman-compose down
	@echo "Building images (when required) and starting pods..."
	podman-compose up --build -d
	@echo "Images are up and running..."

compile:
	@echo "Compiling go-checks-api app..."
	cd ./go-checks-api && env GOOS=linux CGO_ENABLED=0 go build -o ${GO_CHECKS_APP} .
	@echo "Done!"

