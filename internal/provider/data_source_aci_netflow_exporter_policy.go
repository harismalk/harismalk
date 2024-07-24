// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"fmt"

	customtypes "github.com/CiscoDevNet/terraform-provider-aci/v2/internal/custom_types"
	"github.com/ciscoecosystem/aci-go-client/v2/client"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &NetflowExporterPolDataSource{}

func NewNetflowExporterPolDataSource() datasource.DataSource {
	return &NetflowExporterPolDataSource{}
}

// NetflowExporterPolDataSource defines the data source implementation.
type NetflowExporterPolDataSource struct {
	client *client.Client
}

func (d *NetflowExporterPolDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_netflow_exporter_policy")
	resp.TypeName = req.ProviderTypeName + "_netflow_exporter_policy"
	tflog.Debug(ctx, "End metadata of datasource: aci_netflow_exporter_policy")
}

func (d *NetflowExporterPolDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_netflow_exporter_policy")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The netflow_exporter_policy datasource for the 'netflowExporterPol' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Netflow Exporter Policy object.",
			},
			"parent_dn": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Netflow Exporter Policy object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the Netflow Exporter Policy object.`,
			},
			"dscp": schema.StringAttribute{
				CustomType:          customtypes.NetflowExporterPoldscpStringType{},
				Computed:            true,
				MarkdownDescription: `The DSCP value of the Netflow Exporter Policy object.`,
			},
			"destination_address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The destination IP address of the remote node.`,
			},
			"destination_port": schema.StringAttribute{
				CustomType:          customtypes.NetflowExporterPoldstPortStringType{},
				Computed:            true,
				MarkdownDescription: `The destination port of the remote node.`,
			},
			"name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `The name of the Netflow Exporter Policy object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the Netflow Exporter Policy object.`,
			},
			"owner_key": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"source_ip_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The type of the source IP address: It can be one of the available management IP address for a given leaf or a custom IP Address.`,
			},
			"source_address": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The source IP address.`,
			},
			"version": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The collector version.`,
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Computed:            true,
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
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
	tflog.Debug(ctx, "End schema of datasource: aci_netflow_exporter_policy")
}

func (d *NetflowExporterPolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_netflow_exporter_policy")
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
	tflog.Debug(ctx, "End configure of datasource: aci_netflow_exporter_policy")
}

func (d *NetflowExporterPolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_netflow_exporter_policy")
	var data *NetflowExporterPolResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.ParentDn.IsNull() || data.ParentDn.IsUnknown() {
		data.ParentDn = basetypes.NewStringValue("uni/infra")
	}

	setNetflowExporterPolId(ctx, data)

	// Create a copy of the Id for when not found during getAndSetNetflowExporterPolAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))

	getAndSetNetflowExporterPolAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_netflow_exporter_policy data source",
			fmt.Sprintf("The aci_netflow_exporter_policy data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End read of datasource aci_netflow_exporter_policy with id '%s'", data.Id.ValueString()))
}
