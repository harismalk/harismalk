// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/ciscoecosystem/aci-go-client/v2/client"
	"github.com/ciscoecosystem/aci-go-client/v2/container"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
var _ resource.Resource = &L3extConsLblResource{}
var _ resource.ResourceWithImportState = &L3extConsLblResource{}

func NewL3extConsLblResource() resource.Resource {
	return &L3extConsLblResource{}
}

// L3extConsLblResource defines the resource implementation.
type L3extConsLblResource struct {
	client *client.Client
}

// L3extConsLblResourceModel describes the resource data model.
type L3extConsLblResourceModel struct {
	Id            types.String `tfsdk:"id"`
	ParentDn      types.String `tfsdk:"parent_dn"`
	Annotation    types.String `tfsdk:"annotation"`
	Descr         types.String `tfsdk:"description"`
	Name          types.String `tfsdk:"name"`
	NameAlias     types.String `tfsdk:"name_alias"`
	Owner         types.String `tfsdk:"owner"`
	OwnerKey      types.String `tfsdk:"owner_key"`
	OwnerTag      types.String `tfsdk:"owner_tag"`
	Tag           types.String `tfsdk:"tag"`
	TagAnnotation types.Set    `tfsdk:"annotations"`
}

// TagAnnotationL3extConsLblResourceModel describes the resource data model for the children without relation ships.
type TagAnnotationL3extConsLblResourceModel struct {
	Key   types.String `tfsdk:"key"`
	Value types.String `tfsdk:"value"`
}

type L3extConsLblIdentifier struct {
	Name types.String
}

func (r *L3extConsLblResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	tflog.Trace(ctx, "start schema of resource: aci_l3out_consumer_label")
	resp.TypeName = req.ProviderTypeName + "_l3out_consumer_label"
	tflog.Trace(ctx, "end schema of resource: aci_l3out_consumer_label")
}

func (r *L3extConsLblResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "The l3out_consumer_label resource for the 'l3extConsLbl' class",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The distinguished name (DN) of the L3out Consumer Label object.",
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
				},
				Default:             stringdefault.StaticString(globalAnnotation),
				MarkdownDescription: `The annotation of the L3out Consumer Label object.`,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The description of the L3out Consumer Label object.`,
			},
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: `The name of the L3out Consumer Label object.`,
			},
			"name_alias": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The name alias of the L3out Consumer Label object.`,
			},
			"owner": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("infra", "tenant"),
				},
				MarkdownDescription: `The owner of the target relay. The DHCP relay is any host that forwards DHCP packets between clients and servers. The relays are used to forward requests and replies between clients and servers when they are not on the same physical subnet.`,
			},
			"owner_key": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `The key for enabling clients to own their data for entity correlation.`,
			},
			"owner_tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `A tag for enabling clients to add their own data. For example, to indicate who created this object.`,
			},
			"tag": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("alice-blue", "antique-white", "aqua", "aquamarine", "azure", "beige", "bisque", "black", "blanched-almond", "blue", "blue-violet", "brown", "burlywood", "cadet-blue", "chartreuse", "chocolate", "coral", "cornflower-blue", "cornsilk", "crimson", "cyan", "dark-blue", "dark-cyan", "dark-goldenrod", "dark-gray", "dark-green", "dark-khaki", "dark-magenta", "dark-olive-green", "dark-orange", "dark-orchid", "dark-red", "dark-salmon", "dark-sea-green", "dark-slate-blue", "dark-slate-gray", "dark-turquoise", "dark-violet", "deep-pink", "deep-sky-blue", "dim-gray", "dodger-blue", "fire-brick", "floral-white", "forest-green", "fuchsia", "gainsboro", "ghost-white", "gold", "goldenrod", "gray", "green", "green-yellow", "honeydew", "hot-pink", "indian-red", "indigo", "ivory", "khaki", "lavender", "lavender-blush", "lawn-green", "lemon-chiffon", "light-blue", "light-coral", "light-cyan", "light-goldenrod-yellow", "light-gray", "light-green", "light-pink", "light-salmon", "light-sea-green", "light-sky-blue", "light-slate-gray", "light-steel-blue", "light-yellow", "lime", "lime-green", "linen", "magenta", "maroon", "medium-aquamarine", "medium-blue", "medium-orchid", "medium-purple", "medium-sea-green", "medium-slate-blue", "medium-spring-green", "medium-turquoise", "medium-violet-red", "midnight-blue", "mint-cream", "misty-rose", "moccasin", "navajo-white", "navy", "old-lace", "olive", "olive-drab", "orange", "orange-red", "orchid", "pale-goldenrod", "pale-green", "pale-turquoise", "pale-violet-red", "papaya-whip", "peachpuff", "peru", "pink", "plum", "powder-blue", "purple", "red", "rosy-brown", "royal-blue", "saddle-brown", "salmon", "sandy-brown", "sea-green", "seashell", "sienna", "silver", "sky-blue", "slate-blue", "slate-gray", "snow", "spring-green", "steel-blue", "tan", "teal", "thistle", "tomato", "turquoise", "violet", "wheat", "white", "white-smoke", "yellow", "yellow-green"),
				},
				MarkdownDescription: `Specifies the color of a policy label.`,
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
							},
							MarkdownDescription: `The key or password used to uniquely identify this configuration object.`,
						},
						"value": schema.StringAttribute{
							Required: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `The value of the property.`,
						},
					},
				},
			},
		},
	}
}

func (r *L3extConsLblResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	tflog.Trace(ctx, "start configure of resource: aci_l3out_consumer_label")
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
	tflog.Trace(ctx, "end configure of resource: aci_l3out_consumer_label")
}

func (r *L3extConsLblResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	tflog.Trace(ctx, "start create of resource: aci_l3out_consumer_label")
	// On create retrieve information on current state prior to making any changes in order to determine child delete operations
	var stateData *L3extConsLblResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &stateData)...)
	setL3extConsLblId(ctx, stateData)
	messageMap := setL3extConsLblAttributes(ctx, r.client, stateData)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	var data *L3extConsLblResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	setL3extConsLblId(ctx, data)

	tflog.Trace(ctx, fmt.Sprintf("create of resource aci_l3out_consumer_label with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extConsLblResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload, message, messageDetail := getL3extConsLblCreateJsonPayload(ctx, data, tagAnnotationPlan, tagAnnotationState)

	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	requestData, message, messageDetail := doL3extConsLblRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	messageMap = setL3extConsLblAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end create of resource: aci_l3out_consumer_label")
}

func (r *L3extConsLblResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	tflog.Trace(ctx, "start read of resource: aci_l3out_consumer_label")
	var data *L3extConsLblResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("read of resource aci_l3out_consumer_label with id '%s'", data.Id.ValueString()))

	messageMap := setL3extConsLblAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end read of resource: aci_l3out_consumer_label")
}

func (r *L3extConsLblResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	tflog.Trace(ctx, "start update of resource: aci_l3out_consumer_label")
	var data *L3extConsLblResourceModel
	var stateData *L3extConsLblResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &stateData)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("update of resource aci_l3out_consumer_label with id '%s'", data.Id.ValueString()))

	var tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extConsLblResourceModel
	data.TagAnnotation.ElementsAs(ctx, &tagAnnotationPlan, false)
	stateData.TagAnnotation.ElementsAs(ctx, &tagAnnotationState, false)
	jsonPayload, message, messageDetail := getL3extConsLblCreateJsonPayload(ctx, data, tagAnnotationPlan, tagAnnotationState)

	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	requestData, message, messageDetail := doL3extConsLblRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}

	messageMap := setL3extConsLblAttributes(ctx, r.client, data)
	if messageMap != nil {
		resp.Diagnostics.AddError(messageMap.(map[string]string)["message"], messageMap.(map[string]string)["messageDetail"])
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
	tflog.Trace(ctx, "end update of resource: aci_l3out_consumer_label")
}

func (r *L3extConsLblResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	tflog.Trace(ctx, "start delete of resource: aci_l3out_consumer_label")
	var data *L3extConsLblResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Trace(ctx, fmt.Sprintf("delete of resource aci_l3out_consumer_label with id '%s'", data.Id.ValueString()))
	jsonPayload, message, messageDetail := getL3extConsLblDeleteJsonPayload(ctx, data)
	if jsonPayload == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}
	requestData, message, messageDetail := doL3extConsLblRequest(ctx, r.client, fmt.Sprintf("api/mo/%s.json", data.Id.ValueString()), "POST", jsonPayload)
	if requestData == nil {
		resp.Diagnostics.AddError(message, messageDetail)
		return
	}
	tflog.Trace(ctx, "end delete of resource: aci_l3out_consumer_label")
}

func (r *L3extConsLblResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func setL3extConsLblAttributes(ctx context.Context, client *client.Client, data *L3extConsLblResourceModel) interface{} {
	requestData, message, messageDetail := doL3extConsLblRequest(ctx, client, fmt.Sprintf("api/mo/%s.json?rsp-subtree=children&rsp-subtree-class=%s", data.Id.ValueString(), "l3extConsLbl,tagAnnotation"), "GET", nil)

	if requestData == nil {
		return map[string]string{"message": message, "messageDetail": messageDetail}
	}
	if requestData.Search("imdata").Search("l3extConsLbl").Data() != nil {
		classReadInfo := requestData.Search("imdata").Search("l3extConsLbl").Data().([]interface{})
		if len(classReadInfo) == 1 {
			attributes := classReadInfo[0].(map[string]interface{})["attributes"].(map[string]interface{})
			for attributeName, attributeValue := range attributes {
				if attributeName == "dn" {
					data.Id = basetypes.NewStringValue(attributeValue.(string))
					setL3extConsLblParentDn(ctx, attributeValue.(string), data)
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
				if attributeName == "owner" {
					data.Owner = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerKey" {
					data.OwnerKey = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "ownerTag" {
					data.OwnerTag = basetypes.NewStringValue(attributeValue.(string))
				}
				if attributeName == "tag" {
					data.Tag = basetypes.NewStringValue(attributeValue.(string))
				}
			}
			TagAnnotationL3extConsLblList := make([]TagAnnotationL3extConsLblResourceModel, 0)
			_, ok := classReadInfo[0].(map[string]interface{})["children"]
			if ok {
				children := classReadInfo[0].(map[string]interface{})["children"].([]interface{})
				for _, child := range children {
					for childClassName, childClassDetails := range child.(map[string]interface{}) {
						childAttributes := childClassDetails.(map[string]interface{})["attributes"].(map[string]interface{})
						if childClassName == "tagAnnotation" {
							TagAnnotationL3extConsLbl := TagAnnotationL3extConsLblResourceModel{}
							for childAttributeName, childAttributeValue := range childAttributes {
								if childAttributeName == "key" {
									TagAnnotationL3extConsLbl.Key = basetypes.NewStringValue(childAttributeValue.(string))
								}
								if childAttributeName == "value" {
									TagAnnotationL3extConsLbl.Value = basetypes.NewStringValue(childAttributeValue.(string))
								}
							}
							TagAnnotationL3extConsLblList = append(TagAnnotationL3extConsLblList, TagAnnotationL3extConsLbl)
						}
					}
				}
			}
			if len(TagAnnotationL3extConsLblList) > 0 {
				tagAnnotationSet, _ := types.SetValueFrom(ctx, data.TagAnnotation.ElementType(ctx), TagAnnotationL3extConsLblList)
				data.TagAnnotation = tagAnnotationSet
			}
		} else {
			return map[string]string{
				"message":       "too many results in response",
				"messageDetail": fmt.Sprintf("%v matches returned for class 'l3extConsLbl'. Please report this issue to the provider developers.", len(classReadInfo)),
			}
		}
	}
	return nil
}

func getL3extConsLblRn(ctx context.Context, data *L3extConsLblResourceModel) string {
	rn := "conslbl-{name}"
	for _, identifier := range []string{"name"} {
		fieldName := fmt.Sprintf("%s%s", strings.ToUpper(identifier[:1]), identifier[1:])
		fieldValue := reflect.ValueOf(data).Elem().FieldByName(fieldName).Interface().(basetypes.StringValue).ValueString()
		rn = strings.ReplaceAll(rn, fmt.Sprintf("{%s}", identifier), fieldValue)
	}
	return rn
}

func setL3extConsLblParentDn(ctx context.Context, dn string, data *L3extConsLblResourceModel) {
	rn := getL3extConsLblRn(ctx, data)
	data.ParentDn = basetypes.NewStringValue(strings.ReplaceAll(dn, fmt.Sprintf("/%s", rn), ""))
}

func setL3extConsLblId(ctx context.Context, data *L3extConsLblResourceModel) {
	rn := getL3extConsLblRn(ctx, data)
	data.Id = types.StringValue(fmt.Sprintf("%s/%s", data.ParentDn.ValueString(), rn))
}

func getL3extConsLblTagAnnotationChildPayloads(ctx context.Context, data *L3extConsLblResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extConsLblResourceModel) ([]map[string]interface{}, string, string) {

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

	return childPayloads, "", ""
}

func getL3extConsLblCreateJsonPayload(ctx context.Context, data *L3extConsLblResourceModel, tagAnnotationPlan, tagAnnotationState []TagAnnotationL3extConsLblResourceModel) (*container.Container, string, string) {
	payloadMap := map[string]interface{}{}
	payloadMap["attributes"] = map[string]string{}
	childPayloads := []map[string]interface{}{}

	TagAnnotationchildPayloads, errMessage, errMessageDetail := getL3extConsLblTagAnnotationChildPayloads(ctx, data, tagAnnotationPlan, tagAnnotationState)
	if TagAnnotationchildPayloads == nil {
		return nil, errMessage, errMessageDetail
	}
	childPayloads = append(childPayloads, TagAnnotationchildPayloads...)

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
	if !data.Owner.IsNull() && !data.Owner.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["owner"] = data.Owner.ValueString()
	}
	if !data.OwnerKey.IsNull() && !data.OwnerKey.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerKey"] = data.OwnerKey.ValueString()
	}
	if !data.OwnerTag.IsNull() && !data.OwnerTag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["ownerTag"] = data.OwnerTag.ValueString()
	}
	if !data.Tag.IsNull() && !data.Tag.IsUnknown() {
		payloadMap["attributes"].(map[string]string)["tag"] = data.Tag.ValueString()
	}

	payload, err := json.Marshal(map[string]interface{}{"l3extConsLbl": payloadMap})
	if err != nil {
		errMessage := "marshalling of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}

	jsonPayload, err := container.ParseJSON(payload)

	if err != nil {
		errMessage := "construction of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}
	return jsonPayload, "", ""
}

func getL3extConsLblDeleteJsonPayload(ctx context.Context, data *L3extConsLblResourceModel) (*container.Container, string, string) {

	jsonString := fmt.Sprintf(`{"l3extConsLbl":{"attributes":{"dn": "%s","status": "deleted"}}}`, data.Id.ValueString())
	jsonPayload, err := container.ParseJSON([]byte(jsonString))
	if err != nil {
		errMessage := "construction of json payload failed"
		errMessageDetail := fmt.Sprintf("err: %s. Please report this issue to the provider developers.", err)
		return nil, errMessage, errMessageDetail
	}
	return jsonPayload, "", ""
}

func doL3extConsLblRequest(ctx context.Context, client *client.Client, path, method string, payload *container.Container) (*container.Container, string, string) {

	restRequest, err := client.MakeRestRequest(method, path, payload, true)
	if err != nil {
		message := fmt.Sprintf("creation of %s rest request failed", strings.ToLower(method))
		messageDetail := fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err)
		return nil, message, messageDetail
	}

	cont, restResponse, err := client.Do(restRequest)

	if restResponse != nil && restResponse.StatusCode != 200 {
		message := fmt.Sprintf("%s rest request failed", strings.ToLower(method))
		messageDetail := fmt.Sprintf("Response: %s, err: %s. Please report this issue to the provider developers.", cont.Data().(map[string]interface{})["imdata"], err)
		return nil, message, messageDetail
	} else if err != nil {
		message := fmt.Sprintf("%s rest request failed", strings.ToLower(method))
		messageDetail := fmt.Sprintf("Err: %s. Please report this issue to the provider developers.", err)
		return nil, message, messageDetail
	}

	return cont, "", ""
}
