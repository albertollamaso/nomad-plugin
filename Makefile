SHELL=/bin/bash -o pipefail
GO ?= go

NAME := nomad
OUTPUT := lib$(NAME).so
DESTDIR := /usr/share/falco/plugins

ifeq ($(DEBUG), 1)
    GODEBUGFLAGS= GODEBUG=cgocheck=2
else
    GODEBUGFLAGS= GODEBUG=cgocheck=0
endif

all: build

clean:
	@rm -f lib$(NAME).so

build: clean
	@$(GODEBUGFLAGS) $(GO) build -buildmode=c-shared -buildvcs=false -o $(OUTPUT) ./plugin

install: build
	mv $(OUTPUT) $(DESTDIR)/
