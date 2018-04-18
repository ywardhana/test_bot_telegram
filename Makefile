.PHONY: all

pretty:
	gofmt -w *.go

bin: pretty sync
	go build -o kingsman main.go

run: bin
	./kingsman

REGISTRY  = registry.bukalapak.io/bukalapak
DDIR      = deploy
ODIR      = $(DDIR)/_output
NOCACHE   = --no-cache
VERSION   = $(shell git show -q --format=%h)
DEFENV    = production canary sandbox
SERVICES ?= cream
ENV      ?= $(DEFENV)
ACTION   ?= replace
FILE     ?= deployment
SQUAD     = default

all:
	consul compile build push deployment

test:
	govendor test -v -cover +local,^program

dep:
	go get -u github.com/kardianos/govendor
	govendor sync

sync:
	govendor sync

compile:
	@$(foreach var, $(SERVICES), GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(ODIR)/$(var)/bin/$(var) app/$(var)/main.go;)

$(ODIR):
	mkdir -p $(ODIR)

consul: $(ODIR)
	@wget https://releases.hashicorp.com/envconsul/0.6.2/envconsul_0.6.2_linux_amd64.tgz
	@tar -xf envconsul_0.6.2_linux_amd64.tgz -C $(ODIR)/
	@rm envconsul_0.6.2_linux_amd64.tgz

build:
	@$(foreach var, $(SERVICES), docker build $(NOCACHE) -t $(REGISTRY)/cream/$(var):$(VERSION) -f ./deploy/$(var)/Dockerfile .;)

push:
	@$(foreach var, $(SERVICES), docker push $(REGISTRY)/cream/$(var):$(VERSION);)

deployment: $(ODIR)
ifeq ($(ENV),$(DEFENV))
	kubelize deployment -v $(VERSION) $(SERVICES)
else
	kubelize deployment -e $(ENV) -v $(VERSION) $(SERVICES)
endif

$(ENV):
	$(foreach var, $(SERVICES), kubectl --namespace=$(SQUAD) $(ACTION) -f $(ODIR)/$(var)/$@/$(FILE).yml;)
