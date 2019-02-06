GO ?= go
GOFILES := $(shell find . -name "*.go" -type f)
EXECUTABLE := line-login

.PHONY: build
build: $(EXECUTABLE)

.PHONY: line-login
$(EXECUTABLE): $(SOURCES)
	$(GO) build -v -o bin/$@

.PHONY: lint
lint:
	@hash revive > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		$(GO) get -u github.com/mgechev/revive; \
	fi
	revive -config .revive.toml ./... || exit 1

clean:
	$(GO) clean -modcache -cache -x -i ./...
