.PHONY: docker
docker:
	@rm webook || true
	@go mod tidy
	@GOOS=linux GOARCH=arm go build -tags=k8s -o webook .
	@docker rmi -f frankie/webook:v0.0.1
	@docker build -t frankie/webook:v0.0.1 .