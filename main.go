package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/phy1729/terraform-provider-embeddemo/embeddemo"
)

func main() {
	providerserver.Serve(context.Background(), embeddemo.New, providerserver.ServeOpts{ //nolint:errcheck
		Address: "registry.terraform.io/example/embeddemo",
	})
}
