INSTALL_DIR = $(HOME)/source
BINARIES = bookgen imagerender imageswap pipeline planbook scaffold
MSG ?= update

.PHONY: build clean clobber add commit push cmds

cmds: build

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

clobber: clean
	@find . \( -path './.git' -o -path './.git/*' \) -prune -o -type d -name node_modules -print -exec rm -rf {} +

add:
	@git add -A

commit: build
	@git add -A
	@git commit -m "$(MSG)" || true

push: build
	@git add -A
	@git commit -m "$(MSG)" || true
	@git push
