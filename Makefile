GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BIN_DIR=./bin
all: clean build
build:
	mkdir $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/quitter ./cmd/quitter
	$(GOBUILD) -o $(BIN_DIR)/sc ./cmd/sc
install:
	$(GOCMD) install ./cmd/quitter
	$(GOCMD) install ./cmd/sc
clean:
	rm -rf $(BIN_DIR)
	$(GOCLEAN)