.PHONY: menv test testv build run buildexample runexample

PARENT_DIR := $(notdir $(CURDIR))

menv:
	@echo "Current directory: $(CURDIR)"
	@echo "Parent directory name: $(PARENT_DIR)"

test:
	@go test ./...

testv:
	@go test -v ./...

# Build for application
build:
	@echo "Size before build:"; \
	ls -la |grep $(PARENT_DIR); \
	ls -lh |grep $(PARENT_DIR); \
	echo "\n\nSize after build:"; \
	CGO_ENABLED=0 go build --ldflags "-s -w"; \
	strip $(PARENT_DIR); \
	ls -la |grep $(PARENT_DIR); \
	ls -lh |grep $(PARENT_DIR)

# Run for application
run:
	@go run .

# Build for package examples
buildexample:
	@cd example; \
	echo "Size before build:"; \
	ls -la |grep 'example'; \
	ls -lh |grep example; \
	echo "\n\nSize after build:"; \
	CGO_ENABLED=0 go build --ldflags "-s -w"; \
	strip example; \
	ls -la |grep example; \
	ls -lh |grep example; \
	cd ..

# Run for package examples
runexample:
	@cd example; \
	go run main.go; \
	cd ..
