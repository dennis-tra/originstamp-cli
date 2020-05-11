ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=dennis-tra
ENTRY_POINT=cmd/stamp/main.go
VERSION=`cat VERSION`

.PHONY: install
install: version
	go install $(ENTRY_POINT)

.PHONY: build
build: version
	go build $(ENTRY_POINT)

.PHONY: test
test:
	go test -v $(ENTRY_POINT)

.PHONY: version
version:
	go run cmd/version/version.go

.PHONY: release
release: version
	GOOS=windows GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_windows_amd64 $(ENTRY_POINT)
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_darwin_amd64 $(ENTRY_POINT)
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_linux_amd64 $(ENTRY_POINT)
	hub release create -p -a $(ARTIFACTS_DIR)/stamp_windows_amd64 -a $(ARTIFACTS_DIR)/stamp_darwin_amd64 -a $(ARTIFACTS_DIR)/stamp_linux_amd64 $(VERSION)
