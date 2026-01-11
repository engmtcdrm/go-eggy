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
	@cd examples; \
	echo "Size before build:"; \
	ls -la |grep ___examples; \
	ls -lh |grep ___examples; \
	echo "\n\nSize after build:"; \
	CGO_ENABLED=0 go build --ldflags "-s -w" -o ___examples; \
	strip ___examples; \
	ls -la |grep ___examples; \
	ls -lh |grep ___examples; \
	cd ..

# Run for package examples
run:
	@cd examples; \
	go run main.go; \
	cd ..
