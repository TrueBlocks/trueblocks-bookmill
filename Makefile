INSTALL_DIR = $(HOME)/source
BINARIES = bookblurb bookcover imagerender imageswap pipeline planbook scaffold
MSG ?= update

.PHONY: build clean add commit push

build:
	@for bin in $(BINARIES); do \
		echo "Building $$bin..."; \
		go build -o $(INSTALL_DIR)/$$bin ./cmd/$$bin/.; \
		rm -f $$bin; \
	done

clean:
	@for bin in $(BINARIES); do \
		rm -f $(INSTALL_DIR)/$$bin; \
	done

add:
	@git add -A

commit: build
	@git add -A
	@git commit -m "$(MSG)" || true

push: build
	@git add -A
	@git commit -m "$(MSG)" || true
	@git push
