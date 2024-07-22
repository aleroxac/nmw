## ---------- UTILS
.PHONY: help
help: ## Show this menu
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## Clean all temp files
	@rm -rfv .build .run



## ---------- HELPERS
## setup the workspace
define setup
	for folder in .run .build; do \
		[ -d $1 ] && mkdir $1 || true; \
	done
endef

define teardown
	for folder in .run .build; do \
		[ -d $1 ] && rm -rf $1 || true; \
	done
endef



## ---------- BUILD
.PHONY: buildc
buildc: ## build code and generate the binary
	@$(call setup, .build)
	@go build -o .run/nmw
	@$(call teardown, .build)



## ---------- MAIN
.PHONY: runc
runc: ## run code
	@go run *.go

.PHONY: runb
runb: ## run binary
	@[ -f .run/nmw ] && .run/nmw || echo "Please, first that, build the code via 'make buildc' it and try again."
