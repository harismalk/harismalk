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
var _ datasource.DataSource = &L3extConsLblDataSource{}

func NewL3extConsLblDataSource() datasource.DataSource {
	return &L3extConsLblDataSource{}
}

// L3extConsLblDataSource defines the data source implementation.
type L3extConsLblDataSource struct {
	client *client.Client
}

func (d *L3extConsLblDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Trace(ctx, "start schema of datasource: aci_l3out_consumer_label")
	resp.TypeName = req.ProviderTypeName + "_l3out_consumer_label"
	tflog.Trace(ctx, "end schema of datasource: aci_l3out_consumer_label")
}

func (d *L3extConsLblDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_consumer_label datasource for the 'l3extConsLbl' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the L3out Consumer Label object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the L3out Consumer Label object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the L3out Consumer Label object.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the L3out Consumer Label object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the L3out Consumer Label object.`,
			},
			"owner": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The owner of the target relay. The DHCP relay is any host that forwards DHCP packets between clients and servers. The relays are used to forward requests and replies between clients and servers when they are not on the same physical subnet.`,
			},
			"owner_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Specifies the color of a policy label.`,
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

func (d *L3extConsLblDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Trace(ctx, "start configure of datasource: aci_l3out_consumer_label")
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
	tflog.Trace(ctx, "end configure of datasource: aci_l3out_consumer_label")
}

func (d *L3extConsLblDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Trace(ctx, "start read of datasource: aci_l3out_consumer_label")
	var data *L3extConsLblResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setL3extConsLblId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("read of datasource aci_l3out_consumer_label with id '%s'", data.Id.ValueString()))

	messageMap := setL3extConsLblAttributes(ctx, d.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end read of datasource: aci_l3out_consumer_label")
}
