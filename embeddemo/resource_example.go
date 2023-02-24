package embeddemo

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type example struct{}

type commonModel struct {
	Foo types.String `tfsdk:"foo"`
}

type exampleResourceModel struct {
	commonModel
	Bar types.String `tfsdk:"bar"`
	ID  types.String `tfsdk:"id"`
}

func NewExampleResource() resource.Resource {
	return &example{}
}

func (r *example) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_example"
}

func (r *example) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bar": schema.StringAttribute{
				Required: true,
			},
			"foo": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *example) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan exampleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *example) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state exampleResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *example) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan exampleResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *example) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
