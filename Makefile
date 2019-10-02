NAME := gruutbot
BUILD_DIR := build
PLUGINS_DIR := plugins
CONFIGS_DIR := configs
SCRIPTS_DIR := scripts

DATE=$(shell date '+%F %T')
BUILD_VERSIONED_DIR := $(BUILD_DIR)/$(NAME)-$(shell cat version.txt)

VERSION=$(NAME) $(shell cat version.txt) - $(DATE)
VFLAG=-X 'main.VERSION=$(VERSION)'

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

.PHONY: test clean lint linux darwin configs version

build: clean linux darwin

test:
	go test $(PKGS)

clean:
	find $(BUILD_DIR)/ -type f -not -name '.gitkeep' -print0 | xargs -0 rm -f --
	find $(BUILD_DIR)/ -type d -not -name 'build' -print0 | xargs -0 rm -rf --

bump:
	@$(SCRIPTS_DIR)/bump.sh -p v $(call args,patch)
	$(SCRIPTS_DIR)/bump.sh -p v -l > version.txt

PLATFORMS := linux darwin
os = $(word 1, $@)

$(PLATFORMS):
	GOOS=$(os) GOARCH=amd64 go build -ldflags "$(VFLAG)" -o $(BUILD_VERSIONED_DIR)-$(os)/$(NAME)-$(os)-amd64 ./cli
	mkdir -p $(BUILD_VERSIONED_DIR)-$(os)/$(CONFIGS_DIR) && cp $(CONFIGS_DIR)/* $(BUILD_VERSIONED_DIR)-$(os)/$(CONFIGS_DIR)

%::
	@: