// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"fmt"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iamidentityv1"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

// Data source to find all the policies for a serviceID
func DataSourceIBMIAMServicePolicy() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMIAMServicePolicyRead,

		Schema: map[string]*schema.Schema{
			"iam_service_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"iam_service_id", "iam_id"},
				Description:  "UUID of ServiceID",
			},
			"iam_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"iam_service_id", "iam_id"},
				Description:  "IAM ID of ServiceID",
			},
			"sort": {
				Description: "Sort query for policies",
				Type:        schema.TypeString,
				Optional:    true,
			},
			"transaction_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Set transactionID for debug",
			},
			"policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"roles": {
							Type:        schema.TypeList,
							Computed:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Role names of the policy definition",
						},
						"resources": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Service name of the policy definition",
									},
									"resource_instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "ID of resource instance of the policy definition",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Region of the policy definition",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource type of the policy definition",
									},
									"resource": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Resource of the policy definition",
									},
									"resource_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the resource group.",
									},
									"service_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Service type of the policy definition",
									},
									"attributes": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Set resource attributes in the form of 'name=value,name=value....",
										Elem:        schema.TypeString,
									},
								},
							},
						},
						"resource_tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Set access management tags.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of attribute.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of attribute.",
									},
									"operator": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Operator of attribute.",
									},
								},
							},
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the Policy",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMIAMServicePolicyRead(d *schema.ResourceData, meta interface{}) error {

	var iamID string
	if v, ok := d.GetOk("iam_service_id"); ok && v != nil {

		serviceIDUUID := v.(string)
		iamClient, err := meta.(conns.ClientSession).IAMIdentityV1API()
		if err != nil {
			return err
		}
		getServiceIDOptions := iamidentityv1.GetServiceIDOptions{
			ID: &serviceIDUUID,
		}
		serviceID, resp, err := iamClient.GetServiceID(&getServiceIDOptions)
		if err != nil || resp == nil {
			return fmt.Errorf("[ERROR] Error] Error Getting Service Id %s %s", err, resp)
		}
		iamID = *serviceID.IamID
	}
	if v, ok := d.GetOk("iam_id"); ok && v != nil {
		iamID = v.(string)
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}

	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		return err
	}

	listPoliciesOptions := &iampolicymanagementv1.ListPoliciesOptions{
		AccountID: core.StringPtr(userDetails.UserAccount),
		IamID:     core.StringPtr(iamID),
		Type:      core.StringPtr("access"),
	}

	if v, ok := d.GetOk("sort"); ok {
		listPoliciesOptions.Sort = core.StringPtr(v.(string))
	}

	if transactionID, ok := d.GetOk("transaction_id"); ok {
		listPoliciesOptions.SetHeaders(map[string]string{"Transaction-Id": transactionID.(string)})
	}

	policyList, resp, err := iamPolicyManagementClient.ListPolicies(listPoliciesOptions)

	if err != nil {
		return fmt.Errorf("Error listing service policies: %s, %s", err, resp)
	}

	policies := policyList.Policies
	servicePolicies := make([]map[string]interface{}, 0, len(policies))
	for _, policy := range policies {
		roles := make([]string, len(policy.Roles))
		for i, role := range policy.Roles {
			roles[i] = *role.DisplayName
		}
		resources := flex.FlattenPolicyResource(policy.Resources)
		p := map[string]interface{}{
			"roles":         roles,
			"resources":     resources,
			"resource_tags": flex.FlattenPolicyResourceTags(policy.Resources),
		}
		if v, ok := d.GetOk("iam_service_id"); ok && v != nil {
			serviceIDUUID := v.(string)
			p["id"] = fmt.Sprintf("%s/%s", serviceIDUUID, *policy.ID)
		} else if v, ok := d.GetOk("iam_id"); ok && v != nil {
			iamID := v.(string)
			p["id"] = fmt.Sprintf("%s/%s", iamID, *policy.ID)
		}
		if policy.Description != nil {
			p["description"] = policy.Description
		}
		servicePolicies = append(servicePolicies, p)
	}

	if v, ok := d.GetOk("iam_service_id"); ok && v != nil {
		serviceIDUUID := v.(string)
		d.SetId(serviceIDUUID)
	} else if v, ok := d.GetOk("iam_id"); ok && v != nil {
		iamID := v.(string)
		d.SetId(iamID)
	}
	if len(resp.Headers["Transaction-Id"]) > 0 && resp.Headers["Transaction-Id"][0] != "" {
		d.Set("transaction_id", resp.Headers["Transaction-Id"][0])
	}
	d.Set("policies", servicePolicies)
	return nil
}
