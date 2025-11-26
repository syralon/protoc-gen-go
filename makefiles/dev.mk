.PHONY: fmt
# format code
fmt:
	golangci-lint fmt
	buf format -w

.PHONY: lint
# run linter
lint:
	GOWORK=off golangci-lint run


.PHONY: lint-fix
# run linter with auto fix
lint-fix:
	GOWORK=off golangci-lint run --fix
