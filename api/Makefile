.EXPORT_ALL_VARIABLES:



GOBASE := $(shell pwd)
PROJECT_BASE := $(GOBASE)/..

CGO_LDFLAGS := -g -O2 -Llib
CGO_CXXFLAGS :=-I$(PROJECT_BASE)/core/include -g -O2
GOBIN=$(GOBASE)/bin

build:
	@go install  ./...

run_dev:
	tmux  new-session -s "oneseismic dev" "tmux source-file oneseismic-tmux.conf" 
