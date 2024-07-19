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
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &FvFBRGroupResource{}
var _ resource.ResourceWithImportState = &FvFBRGroupResource{}

func NewFvFBRGroupResource() resource.Resource {
	return &FvFBRGroupResource{}
}

// FvFBRGroupResource defines the resource implementation.
type FvFBRGroupResource struct {
	client *client.Client
}

// FvFBRGroupResourceModel describes the resource data model.
type FvFBRGroupResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	FvFBRMember   types.Set    `tfsdk:"vrf_fallback_route_group_members"`
	FvFBRoute     types.Set    `tfsdk:"vrf_fallback_routes"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
	TagTag        types.Set    `tfsdk:"tags"`
}

// FvFBRMemberFvFBRGroupResourceModel describes the resource data model for the children without relation ships.
type FvFBRMemberFvFBRGroupResourceModel struct {
	Annotation types.String `tfsdk:"annotation"`
	Descr      types.String `tfsdk:"description"`
	Name       types.String `tfsdk:"name"`
	NameAlias  types.String `tfsdk:"name_alias"`
	RnhAddr    types.String `tfsdk:"fallback_member"`
}

// FvFBRouteFvFBRGroupResourceModel describes the resource data model for the children without relation ships.
type FvFBRouteFvFBRGroupResourceModel struct {
	Annotation types.String `tfsdk:"annotation"`
	Descr      types.String `tfsdk:"description"`
	FbrPrefix  types.String `tfsdk:"prefix_address"`
	Name       types.String `tfsdk:"name"`
	NameAlias  types.String `tfsdk:"name_alias"`
}

// TagAnnotationFvFBRGroupResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationFvFBRGroupResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

// TagTagFvFBRGroupResourceModel describes the resource data model for the children without relation ships.
type TagTagFvFBRGroupResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type FvFBRGroupIdentifier struct {
	Name types.String
}

func (r *FvFBRGroupResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if !req.Plan.Raw.IsNull() {
		var planData, stateData *FvFBRGroupResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &planData)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

		if resp.Diagnostics.HasError() {
			return
		}

		if (planData.Id.IsUnknown() || planData.Id.IsNull()) && !planData.ParentDn.IsUnknown() && !planData.Name.IsUnknown() {
			setFvFBRGroupId(ctx, planData)
		}

		if stateData == nil && !globalAllowExistingOnCreate && !planData.Id.IsUnknown() && !planData.Id.IsNull() {
			CheckDn(ctx, &resp.Diagnostics, r.client, "fvFBRGroup", planData.Id.ValueString())
			if resp.Diagnostics.HasError() {
				return
			}
		}

		resp.Diagnostics.Append(resp.Plan.Set(ctx, &planData)...)
	}
}

func (r *FvFBRGroupResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Debug(ctx, "Start metadata of resource: aci_vrf_fallback_route_group")
	resp.TypeName = req.ProviderTypeName + "_vrf_fallback_route_group"
	tflog.Debug(ctx, "End metadata of resource: aci_vrf_fallback_route_group")
}

func (r *FvFBRGroupResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	tflog.Debug(ctx, "Start schema of resource: aci_vrf_fallback_route_group")
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The vrf_fallback_route_group resource for the 'fvFBRGroup' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the VRF Fallback Route Group object.",
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
				MarkdownDescription: `The annotation of the VRF Fallback Route Group object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The description of the VRF Fallback Route Group object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the VRF Fallback Route Group object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
				},
				MarkdownDescription: `The name alias of the VRF Fallback Route Group object.`,
			},
			"vrf_fallback_route_group_members": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
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
							},
							MarkdownDescription: `The address of the VRF Fallback Route Group Member object.`,
						},
					},
				},
			},
			"vrf_fallback_routes": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Set{
					setvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"annotation": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The annotation of the VRF Fallback Route object.`,
						},
						"description": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The description of the VRF Fallback Route object.`,
						},
						"prefix_address": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The prefix address of the VRF Fallback Route object.`,
						},
						"name": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The name of the VRF Fallback Route object.`,
						},
						"name_alias": schema.StringAttribute{
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The name alias of the VRF Fallback Route object.`,
						},
					},
				},
			},
			"annotations": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: ``,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The key used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
								SetToStringNullWhenStateIsNullPlanIsUnknownDuringUpdate(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
	tflog.Debug(ctx, "End schema of resource: aci_vrf_fallback_route_group")
}

func (r *FvFBRGroupResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Debug(ctx, "Start configure of resource: aci_vrf_fallback_route_group")
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
	tflog.Debug(ctx, "End configure of resource: aci_vrf_fallback_route_group")
}

func (r *FvFBRGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Debug(ctx, "Start create of resource: aci_vrf_fallback_route_group")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *FvFBRGroupResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	if stateData.Id.IsUnknown() || stateData.Id.IsNull() {
		setFvFBRGroupId(ctx, stateData)
	}
	getAndSetFvFBRGroupAttributes(ctx, &resp.Diagnostics, r.client, stateData)
	if !globalAllowExistingOnCreate && !stateData.Id.IsNull() {
		resp.Diagnostics.AddError(
			"Object Already Exists",
			fmt.Sprintf("The fvFBRGroup object with DN '%s' already exists.", stateData.Id.ValueString()),
		)
		return
	}

	var data *FvFBRGroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	if data.Id.IsUnknown() || data.Id.IsNull() {
		setFvFBRGroupId(ctx, data)
	}

	setFvFBRGroupId(ctx, data)

	var fvFBRMemberPlan, fvFBRMemberState []FvFBRMemberFvFBRGroupResourceModel
	data.FvFBRMember.ElementsAs(ctx, &fvFBRMemberPlan, false)
	stateData.FvFBRMember.ElementsAs(ctx, &fvFBRMemberState, false)
	var fvFBRoutePlan, fvFBRouteState []FvFBRouteFvFBRGroupResourceModel
	data.FvFBRoute.ElementsAs(ctx, &fvFBRoutePlan, false)
	stateData.FvFBRoute.ElementsAs(ctx, &fvFBRouteState, false)
	var tagAnnotationPlan, tagAnnotationState []TagAnnotationFvFBRGroupResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagFvFBRGroupResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getFvFBRGroupCreateJsonPayload(ctx, &resp.Diagnostics, true, data, fvFBRMemberPlan, fvFBRMemberState, fvFBRoutePlan, fvFBRouteState, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetFvFBRGroupAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End create of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Debug(ctx, "Start read of resource: aci_vrf_fallback_route_group")
	var data *FvFBRGroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Read of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))

	getAndSetFvFBRGroupAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	if data.Id.IsNull() {
		var emptyData *FvFBRGroupResourceModel
		resp.Diagnostics.Append(resp.State.Set(ctx, &emptyData)...)
	} else {
		resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	}

	tflog.Debug(ctx, fmt.Sprintf("End read of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Debug(ctx, "Start update of resource: aci_vrf_fallback_route_group")
	var data *FvFBRGroupResourceModel
	var stateData *FvFBRGroupResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))

	var fvFBRMemberPlan, fvFBRMemberState []FvFBRMemberFvFBRGroupResourceModel
	data.FvFBRMember.ElementsAs(ctx, &fvFBRMemberPlan, false)
	stateData.FvFBRMember.ElementsAs(ctx, &fvFBRMemberState, false)
	var fvFBRoutePlan, fvFBRouteState []FvFBRouteFvFBRGroupResourceModel
	data.FvFBRoute.ElementsAs(ctx, &fvFBRoutePlan, false)
	stateData.FvFBRoute.ElementsAs(ctx, &fvFBRouteState, false)
	var tagAnnotationPlan, tagAnnotationState []TagAnnotationFvFBRGroupResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	var tagTagPlan, tagTagState []TagTagFvFBRGroupResourceModel
	data.TagTag.ElementsAs(ctx, &tagTagPlan, false)
	stateData.TagTag.ElementsAs(ctx, &tagTagState, false)
	jsonPayload := getFvFBRGroupCreateJsonPayload(ctx, &resp.Diagnostics, false, data, fvFBRMemberPlan, fvFBRMemberState, fvFBRoutePlan, fvFBRouteState, tagAnnotationPlan, tagAnnotationState, tagTagPlan, tagTagState)

	if resp.Diagnostics.HasError() {
		return
	}

	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Update of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))

	if resp.Diagnostics.HasError() {
		return
	}

	getAndSetFvFBRGroupAttributes(ctx, &resp.Diagnostics, r.client, data)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Debug(ctx, fmt.Sprintf("End update of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Debug(ctx, "Start delete of resource: aci_vrf_fallback_route_group")
	var data *FvFBRGroupResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("Delete of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))
	jsonPayload := GetDeleteJsonPayload(ctx, &resp.Diagnostics, "fvFBRGroup", data.Id.ValueString())
	if resp.Diagnostics.HasError() {
		return
	}
	DoRestRequest(ctx, &resp.Diagnostics, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, fmt.Sprintf("End delete of resource aci_vrf_fallback_route_group with id '%s'", data.Id.ValueString()))
}

func (r *FvFBRGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	tflog.Debug(ctx, "Start import state of resource: aci_vrf_fallback_route_group")
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)

	var stateData *FvFBRGroupResourceModel
	resp.Diagnostics.Append(resp.State.Get(ctx, &stateData)...)
	tflog.Debug(ctx, fmt.Sprintf("Import state of resource aci_vrf_fallback_route_group with id '%s'", stateData.Id.ValueString()))

	tflog.Debug(ctx, "End import of state resource: aci_vrf_fallback_route_group")
}

func getAndSetFvFBRGroupAttributes(ctx context.Context, diags *diag.Diagnostics, client *client.Client, data *FvFBRGroupResourceModel) {
	requestData := DoRestRequest(ctx, diags, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "fvFBRGroup,fvFBRMember,fvFBRoute,tagAnnotation,tagTag"), "GET", nil)

	if diags.HasError() {
		return
	}
	if requestData.Search("imdata").Search("fvFBRGroup").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("fvFBRGroup").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setFvFBRGroupParentDn(ctx, attributeValue.(string), data)
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
			FvFBRMemberFvFBRGroupList := make([]FvFBRMemberFvFBRGroupResourceModel, 0)
			FvFBRouteFvFBRGroupList := make([]FvFBRouteFvFBRGroupResourceModel, 0)
			TagAnnotationFvFBRGroupList := make([]TagAnnotationFvFBRGroupResourceModel, 0)
			TagTagFvFBRGroupList := make([]TagTagFvFBRGroupResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "fvFBRMember" {
							FvFBRMemberFvFBRGroup := FvFBRMemberFvFBRGroupResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "annotation" {
									FvFBRMemberFvFBRGroup.Annotation = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "descr" {
									FvFBRMemberFvFBRGroup.Descr = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "name" {
									FvFBRMemberFvFBRGroup.Name = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "nameAlias" {
									FvFBRMemberFvFBRGroup.NameAlias = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "rnhAddr" {
									FvFBRMemberFvFBRGroup.RnhAddr = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							FvFBRMemberFvFBRGroupList = append(FvFBRMemberFvFBRGroupList, FvFBRMemberFvFBRGroup)
						}
						if childClassName == "fvFBRoute" {
							FvFBRouteFvFBRGroup := FvFBRouteFvFBRGroupResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "annotation" {
									FvFBRouteFvFBRGroup.Annotation = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "descr" {
									FvFBRouteFvFBRGroup.Descr = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "fbrPrefix" {
									FvFBRouteFvFBRGroup.FbrPrefix = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "name" {
									FvFBRouteFvFBRGroup.Name = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "nameAlias" {
									FvFBRouteFvFBRGroup.NameAlias = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							FvFBRouteFvFBRGroupList = append(FvFBRouteFvFBRGroupList, FvFBRouteFvFBRGroup)
						}
						if childClassName == "tagAnnotation" {
							TagAnnotationFvFBRGroup := TagAnnotationFvFBRGroupResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationFvFBRGroup.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationFvFBRGroup.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationFvFBRGroupList = append(TagAnnotationFvFBRGroupList, TagAnnotationFvFBRGroup)
						}
						if childClassName == "tagTag" {
							TagTagFvFBRGroup := TagTagFvFBRGroupResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagTagFvFBRGroup.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagTagFvFBRGroup.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagTagFvFBRGroupList = append(TagTagFvFBRGroupList, TagTagFvFBRGroup)
						}
					}
				}
			}
			fvFBRMemberSet, _ := types.SetValueFrom(ctx, data.FvFBRMember.ElementType(ctx), FvFBRMemberFvFBRGroupList)
			data.FvFBRMember = fvFBRMemberSet
			fvFBRouteSet, _ := types.SetValueFrom(ctx, data.FvFBRoute.ElementType(ctx), FvFBRouteFvFBRGroupList)
			data.FvFBRoute = fvFBRouteSet
			tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationFvFBRGroupList)
			data.TagAnnotation = tagAnnotationSet
			tagTagSet, _ := types.SetValueFrom(ctx, data.TagTag.ElementType(ctx), TagTagFvFBRGroupList)
			data.TagTag = tagTagSet
		} else {
			diags.AddError(
				"too many results in response",
				fmt.Sprintf("%v matches returned for class 'fvFBRGroup'. Please report this issue to the provider developers.", len(classReadInfo)),
			)
		}
	} else {
		data.Id = basetypes.NewStringNull()
	}
}

func getFvFBRGroupRn(ctx context.Context, data *FvFBRGroupResourceModel) string {
	rn := "fbrg-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setFvFBRGroupParentDn(ctx context.Context, dn string, data *FvFBRGroupResourceModel) {
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

func setFvFBRGroupId(ctx context.Context, data *FvFBRGroupResourceModel) {
	rn := getFvFBRGroupRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getFvFBRGroupFvFBRMemberChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *FvFBRGroupResourceModel, fvFBRMemberPlan, fvFBRMemberState []FvFBRMemberFvFBRGroupResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.FvFBRMember.IsUnknown() {
		fvFBRMemberIdentifiers := []FvFBRMemberIdentifier{}
		for _, fvFBRMember := range fvFBRMemberPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !fvFBRMember.Annotation.IsUnknown() {
				childMap["attributes"]["annotation"] = fvFBRMember.Annotation.ValueString()
			} else {
				childMap["attributes"]["annotation"] = globalAnnotation
			}
			if !fvFBRMember.Descr.IsUnknown() {
				childMap["attributes"]["descr"] = fvFBRMember.Descr.ValueString()
			}
			if !fvFBRMember.Name.IsUnknown() {
				childMap["attributes"]["name"] = fvFBRMember.Name.ValueString()
			}
			if !fvFBRMember.NameAlias.IsUnknown() {
				childMap["attributes"]["nameAlias"] = fvFBRMember.NameAlias.ValueString()
			}
			if !fvFBRMember.RnhAddr.IsUnknown() {
				childMap["attributes"]["rnhAddr"] = fvFBRMember.RnhAddr.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"fvFBRMember": childMap})
			fvFBRMemberIdentifier := FvFBRMemberIdentifier{}
			fvFBRMemberIdentifier.RnhAddr = fvFBRMember.RnhAddr
			fvFBRMemberIdentifiers = append(fvFBRMemberIdentifiers, fvFBRMemberIdentifier)
		}
		for _, fvFBRMember := range fvFBRMemberState {
			delete := true
			for _, fvFBRMemberIdentifier := range fvFBRMemberIdentifiers {
				if fvFBRMemberIdentifier.RnhAddr == fvFBRMember.RnhAddr {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["rnhAddr"] = fvFBRMember.RnhAddr.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"fvFBRMember": childMap})
			}
		}
	} else {
		data.FvFBRMember = types.SetNull(data.FvFBRMember.ElementType(ctx))
	}

	return childPayloads
}
func getFvFBRGroupFvFBRouteChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *FvFBRGroupResourceModel, fvFBRoutePlan, fvFBRouteState []FvFBRouteFvFBRGroupResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.FvFBRoute.IsUnknown() {
		fvFBRouteIdentifiers := []FvFBRouteIdentifier{}
		for _, fvFBRoute := range fvFBRoutePlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !fvFBRoute.Annotation.IsUnknown() {
				childMap["attributes"]["annotation"] = fvFBRoute.Annotation.ValueString()
			} else {
				childMap["attributes"]["annotation"] = globalAnnotation
			}
			if !fvFBRoute.Descr.IsUnknown() {
				childMap["attributes"]["descr"] = fvFBRoute.Descr.ValueString()
			}
			if !fvFBRoute.FbrPrefix.IsUnknown() {
				childMap["attributes"]["fbrPrefix"] = fvFBRoute.FbrPrefix.ValueString()
			}
			if !fvFBRoute.Name.IsUnknown() {
				childMap["attributes"]["name"] = fvFBRoute.Name.ValueString()
			}
			if !fvFBRoute.NameAlias.IsUnknown() {
				childMap["attributes"]["nameAlias"] = fvFBRoute.NameAlias.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"fvFBRoute": childMap})
			fvFBRouteIdentifier := FvFBRouteIdentifier{}
			fvFBRouteIdentifier.FbrPrefix = fvFBRoute.FbrPrefix
			fvFBRouteIdentifiers = append(fvFBRouteIdentifiers, fvFBRouteIdentifier)
		}
		for _, fvFBRoute := range fvFBRouteState {
			delete := true
			for _, fvFBRouteIdentifier := range fvFBRouteIdentifiers {
				if fvFBRouteIdentifier.FbrPrefix == fvFBRoute.FbrPrefix {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["fbrPrefix"] = fvFBRoute.FbrPrefix.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"fvFBRoute": childMap})
			}
		}
	} else {
		data.FvFBRoute = types.SetNull(data.FvFBRoute.ElementType(ctx))
	}

	return childPayloads
}
func getFvFBRGroupTagAnnotationChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *FvFBRGroupResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationFvFBRGroupResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagAnnotation.IsUnknown() {
		tagAnnotationIdentifiers := []TagAnnotationIdentifier{}
		for _, tagAnnotation := range tagAnnotationPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagAnnotation.Key.IsUnknown() {
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
			}
			if !tagAnnotation.Value.IsUnknown() {
				childMap["attributes"]["value"] = tagAnnotation.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			tagAnnotationIdentifier := TagAnnotationIdentifier{}
			tagAnnotationIdentifier.Key = tagAnnotation.Key
			tagAnnotationIdentifiers = append(tagAnnotationIdentifiers, tagAnnotationIdentifier)
		}
		for _, tagAnnotation := range tagAnnotationState {
			delete := true
			for _, tagAnnotationIdentifier := range tagAnnotationIdentifiers {
				if tagAnnotationIdentifier.Key == tagAnnotation.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagAnnotation.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagAnnotation": childMap})
			}
		}
	} else {
		data.TagAnnotation = types.SetNull(data.TagAnnotation.ElementType(ctx))
	}

	return childPayloads
}
func getFvFBRGroupTagTagChildPayloads(ctx context.Context, diags *diag.Diagnostics, data *FvFBRGroupResourceModel, tagTagPlan, tagTagState []TagTagFvFBRGroupResourceModel) []map[string]interface{} {

	childPayloads := []map[string]interface{}{}
	if !data.TagTag.IsUnknown() {
		tagTagIdentifiers := []TagTagIdentifier{}
		for _, tagTag := range tagTagPlan {
			childMap := map[string]map[string]interface{}{"attributes": {}}
			if !tagTag.Key.IsUnknown() {
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
			}
			if !tagTag.Value.IsUnknown() {
				childMap["attributes"]["value"] = tagTag.Value.ValueString()
			}
			childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			tagTagIdentifier := TagTagIdentifier{}
			tagTagIdentifier.Key = tagTag.Key
			tagTagIdentifiers = append(tagTagIdentifiers, tagTagIdentifier)
		}
		for _, tagTag := range tagTagState {
			delete := true
			for _, tagTagIdentifier := range tagTagIdentifiers {
				if tagTagIdentifier.Key == tagTag.Key {
					delete = false
					break
				}
			}
			if delete {
				childMap := map[string]map[string]interface{}{"attributes": {}}
				childMap["attributes"]["status"] = "deleted"
				childMap["attributes"]["key"] = tagTag.Key.ValueString()
				childPayloads = append(childPayloads, map[string]interface{}{"tagTag": childMap})
			}
		}
	} else {
		data.TagTag = types.SetNull(data.TagTag.ElementType(ctx))
	}

	return childPayloads
}

func getFvFBRGroupCreateJsonPayload(ctx context.Context, diags *diag.Diagnostics, createType bool, data *FvFBRGroupResourceModel, fvFBRMemberPlan, fvFBRMemberState []FvFBRMemberFvFBRGroupResourceModel, fvFBRoutePlan, fvFBRouteState []FvFBRouteFvFBRGroupResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationFvFBRGroupResourceModel, tagTagPlan, tagTagState []TagTagFvFBRGroupResourceModel) *container.Container {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}

	if createType && !globalAllowExistingOnCreate {
		payloadMap["attributes"].(map[string]string)["status"] = "created"
	}
	childPayloads := []map[string]interface{}{}

	FvFBRMemberchildPayloads := getFvFBRGroupFvFBRMemberChildPayloads(ctx, diags, data, fvFBRMemberPlan, fvFBRMemberState)
	if FvFBRMemberchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, FvFBRMemberchildPayloads...)

	FvFBRoutechildPayloads := getFvFBRGroupFvFBRouteChildPayloads(ctx, diags, data, fvFBRoutePlan, fvFBRouteState)
	if FvFBRoutechildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, FvFBRoutechildPayloads...)

	TagAnnotationchildPayloads := getFvFBRGroupTagAnnotationChildPayloads(ctx, diags, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

	TagTagchildPayloads := getFvFBRGroupTagTagChildPayloads(ctx, diags, data, tagTagPlan, tagTagState)
	if TagTagchildPayloads == nil {
		return nil
	}
	childPayloads = append(childPayloads, TagTagchildPayloads...)

	payloadMap["children"] = childPayloads
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
	payload, err := json.Marshal(map[string]interface{}{"fvFBRGroup": payloadMap})
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
