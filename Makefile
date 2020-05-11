ARTIFACTS_DIR=artifacts/${VERSION}
GITHUB_USERNAME=dennis-tra
ENTRY_POINT=cmd/stamp/main.go

.PHONY: install
install:
	go install $(ENTRY_POINT)

.PHONY: build
build:
	go build $(ENTRY_POINT)

.PHONY: test
test:
	go test -v $(ENTRY_POINT)

.PHONY: release
release:
	GOOS=windows GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_windows_amd64 $(ENTRY_POINT)
	GOOS=darwin GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_darwin_amd64 $(ENTRY_POINT)
	GOOS=linux GOARCH=amd64 go build -o $(ARTIFACTS_DIR)/stamp_linux_amd64 $(ENTRY_POINT)
# 	ghr -u $(GITHUB_USERNAME) -t $(shell cat github_token) --replace ${VERSION} $(ARTIFACTS_DIR)
