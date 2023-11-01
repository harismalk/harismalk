// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &VzOOBBrCPDataSource{}

func NewVzOOBBrCPDataSource() datasource.DataSource {
	return &VzOOBBrCPDataSource{}
}

// VzOOBBrCPDataSource defines the data source implementation.
type VzOOBBrCPDataSource struct {
	client *client.Client
}

func (d *VzOOBBrCPDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Trace(ctx, "start schema of datasource: aci_out_of_band_contract")
	resp.TypeName = req.ProviderTypeName + "_out_of_band_contract"
	tflog.Trace(ctx, "end schema of datasource: aci_out_of_band_contract")
}

func (d *VzOOBBrCPDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The out_of_band_contract datasource for the 'vzOOBBrCP' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Out Of Band Contract object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Out Of Band Contract object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the Out Of Band Contract object.`,
			},
			"intent": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Install Rules or Estimate Nummber of Rules.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the Out Of Band Contract object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the Out Of Band Contract object.`,
			},
			"owner_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"priority": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `null.`,
			},
			"scope": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile.`,
			},
			"target_dscp": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `contract level dscp value.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key or password used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
}

func (d *VzOOBBrCPDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Trace(ctx, "start configure of datasource: aci_out_of_band_contract")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
	tflog.Trace(ctx, "end configure of datasource: aci_out_of_band_contract")
}

func (d *VzOOBBrCPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "start read of datasource: aci_out_of_band_contract")
	var data *VzOOBBrCPResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setVzOOBBrCPId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("read of datasource aci_out_of_band_contract with id '%s'", data.Id.ValueString()))

	messageMap := setVzOOBBrCPAttributes(ctx, d.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end read of datasource: aci_out_of_band_contract")
}
