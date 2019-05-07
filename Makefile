VERSION=1.0.0
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BIN_DIR=./bin
all: clean build
build:
	mkdir $(BIN_DIR)
	$(GOBUILD) -o $(BIN_DIR)/quitter ./cmd/quitter
	$(GOBUILD) -o $(BIN_DIR)/sc ./cmd/sc
	$(GOBUILD) -o $(BIN_DIR)/rm ./cmd/rm
install:
	$(GOCMD) install ./cmd/quitter
	$(GOCMD) install ./cmd/sc
	$(GOCMD) install ./cmd/rm
clean:
	rm -rf $(BIN_DIR)
	$(GOCLEAN)