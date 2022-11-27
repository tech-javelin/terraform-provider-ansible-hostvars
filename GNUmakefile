default: testacc

# Run acceptance tests
.PHONY: testacc install
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

install:
	go fmt
	go mod tidy
	go build -o terraform-provider-bindrole
	mkdir -p ~/.terraform.d/plugins/registry.terraform.io/techjavelin/bindrole/0.0.1/linux_amd64
	cp terraform-provider-bindrole ~/.terraform.d/plugins/registry.terraform.io/techjavelin/bindrole/0.0.1/linux_amd64
