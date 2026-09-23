INSTALL_DIR = $(HOME)/source
BINARIES = bookgen check-source colorize compose export extract-images extract-text imagerender inventory pipeline planbook proof proofscan read-annotations scaffold
MSG ?= update

.PHONY: build clean clobber add commit push cmds lint

lint:
	@golangci-lint run

cmds: build

build:
	@mkdir -p $(INSTALL_DIR)
	@echo "Building bookmill tools: $(BINARIES)"
	@go build -o $(INSTALL_DIR)/ $(addprefix ./cmd/,$(BINARIES))

clean:
	@for bin in $(BINARIES); do \
		rm -f $(INSTALL_DIR)/$$bin; \
	done

clobber: clean
	@find . \( -path './.git' -o -path './.git/*' \) -prune -o -type d -name node_modules -print -exec rm -rf {} +

add:
	@git add -A

commit: lint
	@git add -A
	@git commit -m "$(MSG)" || true

push: lint
	@git add -A
	@git commit -m "$(MSG)" || true
	@git push
