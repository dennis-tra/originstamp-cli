ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=dennis-tra
ENTRY_POINT=cmd/stamp/main.go
VERSION=`cat VERSION`
TARGETS=darwin # windows linux

$(TARGETS): version
	GOOS=$@ GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_$@_amd64 $(ENTRY_POINT)

.PHONY: build
build: $(TARGETS)

.PHONY: test
test:
	go test -v $(ENTRY_POINT)

.PHONY: version
version:
	go run cmd/version/version.go

.PHONY: clean
clean:
	rm -r artifacts

.PHONY: release
release: build
	hub release create -p $(foreach target,$(TARGETS),-a $(ARTIFACTS_DIR)/stamp_$(target)_amd64) $(VERSION)

