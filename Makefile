include .project

BUILDS=\
  darwin-arm64  \
#   darwin-amd64  \

#   linux-amd64   \
# 	linux-386     \
# 	linux-arm     \
# 	linux-arm64   \

# 	windows-386   \
#   windows-amd64 \

MAKEFLAGS += --silent

DISTS_PATTERN=dist/$(NAME)-%-$(VERSION).tgz
DISTS=$(BUILDS:%=$(DISTS_PATTERN))

BINS_PATTERN=bin/%/$(NAME)
BINS=$(BUILDS:%=$(BINS_PATTERN))

BASE := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
# BASE := $(dir $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY:
all: $(DISTS)


$(BINS): OS = $(word 1,$(subst -, ,$*))
$(BINS): ARCH = $(word 2,$(subst -, ,$*))
$(BINS): $(BINS_PATTERN):
	echo "building bin: [$(OS) $(ARCH)]"
	mkdir -p `dirname $@`
	cd src && GOPATH=$(BASE) GOOS=$(OS) GOARCH=$(ARCH) go build -o $(BASE)/$@
	touch $@


$(DISTS): OS = $(word 1,$(subst -, ,$*))
$(DISTS): ARCH = $(word 2,$(subst -, ,$*))
$(DISTS): $(DISTS_PATTERN): $(BINS)
	echo "building dist: [$(OS) $(ARCH)]"
	mkdir -p `dirname $@`
	cd bin/$(OS)-$(ARCH) && tar czf $(BASE)/$@ .
	touch $@

.PHONY: clean
clean:
	rm -rf dist bin

.PHONY: release
release: all
	git tag -a $(VERSION) -m "releasing $(VERSION)"
	git push --tags origin master
