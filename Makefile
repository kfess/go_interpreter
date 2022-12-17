.PHONY: build run test cover fmt vet clean

# Color for Test Success / Fail
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m
COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

GO_CMD = go
GO_OS = linux
GO_BUILD = $(GO_CMD) build
GO_RUN = $(GO_CMD) run
GO_CLEAN = $(GO_CMD) clean
GO_TEST = $(GO_CMD) test -v
GO_FMT = $(GO_CMD) fmt
GO_VET = $(GO_CMD) vet
GO_LDFLAGS = -ldflags="-s -w"

EXECUTABLE = bin/main
TARGET_DIR = bin
GO_PKGROOT = ./...

build:
	$(GO_BUILD) -o $(EXECUTABLE)

run: build
	./$(EXECUTABLE)

test:
	env GOOS=$(GO_OS) $(GO_TEST) $(GO_PKGROOT) | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

cover:
	$(GO_TEST) -cover ./... -coverprofile=cover.out.tmp
	go tool cover -html=cover.out.tmp -o cover.html
	rm cover.out.tmp

fmt:
	$(GO_FMT) $(GO_PKGROOT)

vet:
	$(GO_VET) $(GO_PKGROOT)

clean:
	rm -rf $(TARGET_DIR)
	if [ -e "cover.html" ]; then rm cover.html; fi