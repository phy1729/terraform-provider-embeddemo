package embeddemo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type embeddemoProvider struct{}

func New() provider.Provider {
	return embeddemoProvider{}
}

func (p embeddemoProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "embeddemo"
}

func (p embeddemoProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p embeddemoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p embeddemoProvider) DataSources(context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p embeddemoProvider) Resources(context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewExampleResource,
	}
}
