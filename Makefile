.PHONY: build

lint:
	./scripts/makescript.sh --lint

test:
	./scripts/makescript.sh --test

bench:
	./scripts/makescript.sh --bench

build:
	./scripts/makescript.sh --build

run:
	./scripts/makescript.sh --run
