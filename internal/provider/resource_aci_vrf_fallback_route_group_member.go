// Code generated by "gen/generator.go"; DO NOT EDIT.
// In order to regenerate this file execute `go generate` from the repository root.
// More details can be found in the [README](https://github.com/CiscoDevNet/terraform-provider-aci/blob/master/README.md).

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FvFBRMemberResource{}
var _ resource.ResourceWithImportState = &FvFBRMemberResource{}

func NewFvFBRMemberResource() resource.Resource {
	return &FvFBRMemberResource{}
}

// FvFBRMemberResource defines the resource implementation.
type FvFBRMemberResource struct {
	client *client.Client
}

// FvFBRMemberResourceModel describes the resource data model.
type FvFBRMemberResourceModel struct {
	Id         types.String `tfsdk:"id"`
	ParentDn   types.String `tfsdk:"parent_dn"`
	Annotation types.String `tfsdk:"annotation"`
	Descr      types.String `tfsdk:"description"`
	Name       types.String `tfsdk:"name"`
	NameAlias  types.String `tfsdk:"name_alias"`
	RnhAddr    types.String `tfsdk:"fallback_member"`
}

type FvFBRMemberIdentifier struct {
	RnhAddr types.String
}

func (r *FvFBRMemberResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_vrf_fallback_route_group_member")
	resp.TypeName = req.ProviderTypeName + "_vrf_fallback_route_group_member"
	tflog.Debug(ctx, "End metadata of resource: aci_vrf_fallback_route_group_member")
}

func (r *FvFBRMemberResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_vrf_fallback_route_group_member")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The vrf_fallback_route_group_member resource for the 'fvFBRMember' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the VRF Fallback Route Group Member object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"parent_dn": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "The distinguished name (DN) of the parent object.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
			},
			"annotation": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the VRF Fallback Route Group Member object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the VRF Fallback Route Group Member object.`,
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name of the VRF Fallback Route Group Member object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the VRF Fallback Route Group Member object.`,
			},
			"fallback_member": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The address of the VRF Fallback Route Group Member object.`,
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_vrf_fallback_route_group_member")
}

func (r *FvFBRMemberResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_vrf_fallback_route_group_member")
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
	tflog.Debug(ctx, "End configure of resource: aci_vrf_fallback_route_group_member")
}

func (r *FvFBRMemberResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_vrf_fallback_route_group_member")

	var data *FvFBRMemberResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setFvFBRMemberId(ctx, data)

	tflog.Debug(ctx, fmt.Sprintf("Create of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))

	jsonPayload := getFvFBRMemberCreateJsonPayload(ctx, &resp.Diagnostics, data)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetFvFBRMemberAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRMemberResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_vrf_fallback_route_group_member")
	var data *FvFBRMemberResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))

	getAndSetFvFBRMemberAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *FvFBRMemberResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRMemberResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_vrf_fallback_route_group_member")
	var data *FvFBRMemberResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))

	jsonPayload := getFvFBRMemberCreateJsonPayload(ctx, &resp.Diagnostics, data)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetFvFBRMemberAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRMemberResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_vrf_fallback_route_group_member")
	var data *FvFBRMemberResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "fvFBRMember", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_vrf_fallback_route_group_member with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRMemberResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_vrf_fallback_route_group_member")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *FvFBRMemberResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_vrf_fallback_route_group_member with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_vrf_fallback_route_group_member")
}

func getAndSetFvFBRMemberAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *FvFBRMemberResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("fvFBRMember").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("fvFBRMember").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setFvFBRMemberParentDn(ctx, attributeValue.(string), data)
				}
				if attributeName == "annotation" {
					data.Annotation = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "descr" {
					data.Descr = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "name" {
					data.Name = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "nameAlias" {
					data.NameAlias = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "rnhAddr" {
					data.RnhAddr = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			if data.Annotation.IsUnknown() {
				data.Annotation = types.StringNull()
			}
			if data.Descr.IsUnknown() {
				data.Descr = types.StringNull()
			}
			if data.Name.IsUnknown() {
				data.Name = types.StringNull()
			}
			if data.NameAlias.IsUnknown() {
				data.NameAlias = types.StringNull()
			}
			if data.RnhAddr.IsUnknown() {
				data.RnhAddr = types.StringNull()
			}
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'fvFBRMember'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getFvFBRMemberRn(ctx context.Context, data *FvFBRMemberResourceModel) string {
	rn := "nexthop-[{rnhAddr}]"
	for _, identifier := range []string{"rnhAddr"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setFvFBRMemberParentDn(ctx context.Context, dn string, data *FvFBRMemberResourceModel) {
	bracketIndex := 0
	rnIndex := 0
	for i := len(dn) - 1; i >= 0; i-- {
		if string(dn[i]) == "]" {
			bracketIndex = bracketIndex + 1
		} else if string(dn[i]) == "[" {
			bracketIndex = bracketIndex - 1
		} else if string(dn[i]) == "/" && bracketIndex == 0 {
			rnIndex = i
			break
		}
	}
	data.ParentDn = basetypes.NewStringValue(dn[:rnIndex])
}

func setFvFBRMemberId(ctx context.Context, data *FvFBRMemberResourceModel) {
	rn := getFvFBRMemberRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getFvFBRMemberCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, data *FvFBRMemberResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	if !data.Annotation.IsNull() && !data.Annotation.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["annotation"] = data.Annotation.ValueString()
	}
	if !data.Descr.IsNull() && !data.Descr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["descr"] = data.Descr.ValueString()
	}
	if !data.Name.IsNull() && !data.Name.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["name"] = data.Name.ValueString()
	}
	if !data.NameAlias.IsNull() && !data.NameAlias.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["nameAlias"] = data.NameAlias.ValueString()
	}
	if !data.RnhAddr.IsNull() && !data.RnhAddr.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["rnhAddr"] = data.RnhAddr.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"fvFBRMember": payloadMap})
	if err != nil {
		diags.AddError(
			"Marshalling of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}

	jsonPayload, err := container.ParseJSON(payload)

	if err != nil {
		diags.AddError(
			"Construction of json payload failed",
			fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err),
		)
		return nil
	}
	return jsonPayload
}
