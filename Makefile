ifndef GOOS
UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Darwin)
	GOOS := darwin
else ifeq ($(UNAME_S),Linux)
	GOOS := linux
else
$(error "$$GOOS is not defined. If you are using Windows, try to re-make using 'GOOS=windows make ...' ")
endif
endif

PACKAGES    := $(shell go list ./... | grep -v '/lib/')

BUILD_FLAGS := -ldflags "-X github.com/janotchain/janeta/version.GitCommit=`git rev-parse HEAD`"

janetaD_BINARY32 := janetad-$(GOOS)_386
janetaD_BINARY64 := janetad-$(GOOS)_amd64

janetaCLI_BINARY32 := janetacli-$(GOOS)_386
janetaCLI_BINARY64 := janetacli-$(GOOS)_amd64

VERSION := $(shell awk -F= '/Version =/ {print $$2}' version/version.go | tr -d "\" ")

janetaD_RELEASE32 := janetad-$(VERSION)-$(GOOS)_386
janetaD_RELEASE64 := janetad-$(VERSION)-$(GOOS)_amd64

janetaCLI_RELEASE32 := janetacli-$(VERSION)-$(GOOS)_386
janetaCLI_RELEASE64 := janetacli-$(VERSION)-$(GOOS)_amd64

janeta_RELEASE32 := janeta-$(VERSION)-$(GOOS)_386
janeta_RELEASE64 := janeta-$(VERSION)-$(GOOS)_amd64

all: test target release-all install

janetad:
	@echo "Building janetad to cmd/janetad/janetad"
	@go build $(BUILD_FLAGS) -o cmd/janetad/janetad cmd/janetad/main.go

janetacli:
	@echo "Building janetacli to cmd/janetacli/janetacli"
	@go build $(BUILD_FLAGS) -o cmd/janetacli/janetacli cmd/janetacli/main.go

install:
	@echo "Installing janetad and janetacli to $(GOPATH)/bin"
	@go install ./cmd/janetad
	@go install ./cmd/janetacli

target:
	mkdir -p $@

binary: target/$(janetaD_BINARY32) target/$(janetaD_BINARY64) target/$(janetaCLI_BINARY32) target/$(janetaCLI_BINARY64)

ifeq ($(GOOS),windows)
release: binary
	cd target && cp -f $(janetaD_BINARY32) $(janetaD_BINARY32).exe
	cd target && cp -f $(janetaCLI_BINARY32) $(janetaCLI_BINARY32).exe
	cd target && md5sum  $(janetaD_BINARY32).exe $(janetaCLI_BINARY32).exe >$(janeta_RELEASE32).md5
	cd target && zip $(janeta_RELEASE32).zip  $(janetaD_BINARY32).exe $(janetaCLI_BINARY32).exe $(janeta_RELEASE32).md5
	cd target && rm -f  $(janetaD_BINARY32) $(janetaCLI_BINARY32)  $(janetaD_BINARY32).exe $(janetaCLI_BINARY32).exe $(janeta_RELEASE32).md5
	cd target && cp -f $(janetaD_BINARY64) $(janetaD_BINARY64).exe
	cd target && cp -f $(janetaCLI_BINARY64) $(janetaCLI_BINARY64).exe
	cd target && md5sum  $(janetaD_BINARY64).exe $(janetaCLI_BINARY64).exe >$(janeta_RELEASE64).md5
	cd target && zip $(janeta_RELEASE64).zip  $(janetaD_BINARY64).exe $(janetaCLI_BINARY64).exe $(janeta_RELEASE64).md5
	cd target && rm -f  $(janetaD_BINARY64) $(janetaCLI_BINARY64)  $(janetaD_BINARY64).exe $(janetaCLI_BINARY64).exe $(janeta_RELEASE64).md5
else
release: binary
	cd target && md5sum  $(janetaD_BINARY32) $(janetaCLI_BINARY32) >$(janeta_RELEASE32).md5
	cd target && tar -czf $(janeta_RELEASE32).tgz  $(janetaD_BINARY32) $(janetaCLI_BINARY32) $(janeta_RELEASE32).md5
	cd target && rm -f  $(janetaD_BINARY32) $(janetaCLI_BINARY32) $(janeta_RELEASE32).md5
	cd target && md5sum  $(janetaD_BINARY64) $(janetaCLI_BINARY64) >$(janeta_RELEASE64).md5
	cd target && tar -czf $(janeta_RELEASE64).tgz  $(janetaD_BINARY64) $(janetaCLI_BINARY64) $(janeta_RELEASE64).md5
	cd target && rm -f  $(janetaD_BINARY64) $(janetaCLI_BINARY64) $(janeta_RELEASE64).md5
endif

release-all: clean
	GOOS=darwin  make release
	GOOS=linux   make release
	GOOS=windows make release

clean:
	@echo "Cleaning binaries built..."
	@rm -rf cmd/janetad/janetad
	@rm -rf cmd/janetacli/janetacli
	@rm -rf target
	@rm -rf $(GOPATH)/bin/janetad
	@rm -rf $(GOPATH)/bin/janetacli
	@echo "Cleaning temp test data..."
	@rm -rf test/pseudo_hsm*
	@rm -rf blockchain/pseudohsm/testdata/pseudo/
	@echo "Cleaning sm2 pem files..."
	@rm -rf crypto/sm2/*.pem
	@echo "Done."

target/$(janetaD_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/janetad/main.go

target/$(janetaD_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/janetad/main.go

target/$(janetaCLI_BINARY32):
	CGO_ENABLED=0 GOARCH=386 go build $(BUILD_FLAGS) -o $@ cmd/janetacli/main.go

target/$(janetaCLI_BINARY64):
	CGO_ENABLED=0 GOARCH=amd64 go build $(BUILD_FLAGS) -o $@ cmd/janetacli/main.go

test:
	@echo "====> Running go test"
	@go test $(PACKAGES)

benchmark:
	@go test -bench $(PACKAGES)

functional-tests:
	@go test -timeout=5m -tags="functional" ./test 

ci: test

.PHONY: all target release-all clean test benchmark
