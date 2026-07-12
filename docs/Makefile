.PHONY: docs-serve docs-build

HUGO ?= hugo

docs-serve:
	@echo "Open http://127.0.0.1:1313/go-unit/ (English) or http://127.0.0.1:1313/go-unit/zh/ (中文)"
	$(HUGO) mod get
	$(HUGO) server -D

docs-build:
	$(HUGO) mod get
	$(HUGO) --minify

docs-check:
	@$(HUGO) version | grep -q extended || (echo "Lotus Docs requires Hugo Extended (e.g. asdf install hugo extended-0.164.0)" >&2; exit 1)
