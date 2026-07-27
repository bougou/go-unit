.PHONY: docs-serve docs-build docs-check

PNPM ?= pnpm

docs-serve:
	@echo "Open http://127.0.0.1:4321/go-unit/ (English) or http://127.0.0.1:4321/go-unit/zh/ (中文)"
	$(PNPM) install
	$(PNPM) dev

docs-build:
	$(PNPM) install
	$(PNPM) build

docs-check: docs-build
	$(PNPM) check-links
