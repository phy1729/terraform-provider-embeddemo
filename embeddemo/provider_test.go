package embeddemo_test

import (
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
	"github.com/phy1729/terraform-provider-embeddemo/embeddemo"
)

// map cannot be const and this is idiomatic for terraform.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){ //nolint:gochecknoglobals
	"embeddemo": providerserver.NewProtocol6WithError(embeddemo.New()),
}
