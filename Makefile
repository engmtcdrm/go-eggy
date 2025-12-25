.PHONY: menv test testv build run buildexample runexample

PARENT_DIR := $(notdir $(CURDIR))

menv:
	@echo "Current directory: $(CURDIR)"
	@echo "Parent directory name: $(PARENT_DIR)"

test:
	@go test ./...

testv:
	@go test -v ./...

# Build for package examples
build:
	@cd example; \
	echo "Size before build:"; \
	ls -la |grep ___example; \
	ls -lh |grep ___example; \
	echo "\n\nSize after build:"; \
	CGO_ENABLED=0 go build --ldflags "-s -w" -o ___example; \
	strip ___example; \
	ls -la |grep ___example; \
	ls -lh |grep ___example; \
	cd ..

# Run for package examples
run:
	@cd example; \
	go run main.go; \
	cd ..
