build:
	GOOS=linux GOARCH=amd64 go build -o terraform.d/plugins/registry.terraform.io/Skedulo/kustomization/1.0.0/linux_amd64/terraform-provider-kustomization
	GOOS=darwin GOARCH=amd64 go build -o terraform.d/plugins/registry.terraform.io/Skedulo/kustomization/1.0.0/darwin_amd64/terraform-provider-kustomization

test:
	TF_ACC=1 go test -v ./kustomize
