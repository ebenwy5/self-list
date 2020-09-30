.PHONY: build test run run-full

build:
	@echo " >>>>>>>>>> Building [app]..."
	@go build
	@echo " >>>>>>>>>> Finish building [app]"

test:
	@echo " >>>>>>>>>> Testing [app]..."
	@go test ./... -v
	@echo " >>>>>>>>>> Finish testing [app]"

run: build
	@echo " >>>>>>>>>> Running [app]..."
	@./self-list
	@echo " >>>>>>>>>> Finish running [app]"

run-full: build test run