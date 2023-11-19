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
var _ datasource.DataSource = &PimRouteMapEntryDataSource{}

func NewPimRouteMapEntryDataSource() datasource.DataSource {
	return &PimRouteMapEntryDataSource{}
}

// PimRouteMapEntryDataSource defines the data source implementation.
type PimRouteMapEntryDataSource struct {
	client *client.Client
}

func (d *PimRouteMapEntryDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of datasource: aci_pim_route_map_entry")
	resp.TypeName = req.ProviderTypeName + "_pim_route_map_entry"
	tflog.Debug(ctx, "End metadata of datasource: aci_pim_route_map_entry")
}

func (d *PimRouteMapEntryDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of datasource: aci_pim_route_map_entry")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The pim_route_map_entry datasource for the 'pimRouteMapEntry' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the Pim Route Map Entry object.",
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
			},
			"action": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `route action.`,
			},
			"annotation": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The annotation of the Pim Route Map Entry object.`,
			},
			"description": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The description of the Pim Route Map Entry object.`,
			},
			"grp": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Multicast group ip/prefix.`,
			},
			"name": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name of the Pim Route Map Entry object.`,
			},
			"name_alias": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `The name alias of the Pim Route Map Entry object.`,
			},
			"order": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: `PIM route map entry order.`,
			},
			"rp": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Multicast RP Ip.`,
			},
			"src": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: `Multicast Source Ip.`,
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
	tflog.Debug(ctx, "End schema of datasource: aci_pim_route_map_entry")
}

func (d *PimRouteMapEntryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of datasource: aci_pim_route_map_entry")
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
	tflog.Debug(ctx, "End configure of datasource: aci_pim_route_map_entry")
}

func (d *PimRouteMapEntryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "Start read of datasource: aci_pim_route_map_entry")
	var data *PimRouteMapEntryResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setPimRouteMapEntryId(ctx, data)

	// Create a copy of the Id for when not found during setPimRouteMapEntryAttributes
	cachedId := data.Id.ValueString()

	tflog.Debug(ctx, fmt.Sprintf("Read of datasource aci_pim_route_map_entry with id '%s'", data.Id.ValueString()))

	setPimRouteMapEntryAttributes(ctx, &resp.Diagnostics, d.client, data)

	if data.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Failed to read aci_pim_route_map_entry data source",
			fmt.Sprintf("The aci_pim_route_map_entry data source with id '%s' has not been found", cachedId),
		)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, "End read of datasource: aci_pim_route_map_entry")
}
