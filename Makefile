install-fly:
	curl -L https://fly.io/install.sh | sh

deploy:
	fly deploy

run:
	go run cmd/main.go