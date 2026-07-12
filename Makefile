.PHONY: test docs-serve docs-build

test:
	go test ./...

docs-serve docs-build:
	$(MAKE) -C docs $@
