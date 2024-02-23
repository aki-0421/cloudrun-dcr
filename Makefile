.PHONY: deploy_edge
deploy_edge:
	cd edge && wrangler deploy

.PHONY: run
run:
	PORT=8080 go run main.go
