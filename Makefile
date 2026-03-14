INSTALL_DIR = $(HOME)/source

BINARIES = imagerender imageswap pipeline planbook scaffold

.PHONY: build clean

build:
	@for bin in $(BINARIES); do \
		echo "Building $$bin..."; \
		go build -o $(INSTALL_DIR)/$$bin ./cmd/$$bin/.; \
	done

clean:
	@for bin in $(BINARIES); do \
		rm -f $(INSTALL_DIR)/$$bin; \
	done
