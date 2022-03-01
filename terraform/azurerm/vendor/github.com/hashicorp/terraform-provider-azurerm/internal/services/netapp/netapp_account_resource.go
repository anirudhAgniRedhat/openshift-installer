package netapp

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/netapp/mgmt/2021-06-01/netapp"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/azure"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/tf"
	"github.com/hashicorp/terraform-provider-azurerm/helpers/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/clients"
	"github.com/hashicorp/terraform-provider-azurerm/internal/services/netapp/parse"
	netAppValidate "github.com/hashicorp/terraform-provider-azurerm/internal/services/netapp/validate"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tags"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azurerm/internal/tf/validation"
	"github.com/hashicorp/terraform-provider-azurerm/internal/timeouts"
	"github.com/hashicorp/terraform-provider-azurerm/utils"
)

func resourceNetAppAccount() *pluginsdk.Resource {
	return &pluginsdk.Resource{
		Create: resourceNetAppAccountCreateUpdate,
		Read:   resourceNetAppAccountRead,
		Update: resourceNetAppAccountCreateUpdate,
		Delete: resourceNetAppAccountDelete,

		Timeouts: &pluginsdk.ResourceTimeout{
			Create: pluginsdk.DefaultTimeout(30 * time.Minute),
			Read:   pluginsdk.DefaultTimeout(5 * time.Minute),
			Update: pluginsdk.DefaultTimeout(30 * time.Minute),
			Delete: pluginsdk.DefaultTimeout(30 * time.Minute),
		},
		Importer: pluginsdk.ImporterValidatingResourceId(func(id string) error {
			_, err := parse.AccountID(id)
			return err
		}),

		Schema: map[string]*pluginsdk.Schema{
			"name": {
				Type:         pluginsdk.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: netAppValidate.AccountName,
			},

			"resource_group_name": azure.SchemaResourceGroupName(),

			"location": azure.SchemaLocation(),

			"active_directory": {
				Type:     pluginsdk.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &pluginsdk.Resource{
					Schema: map[string]*pluginsdk.Schema{
						"dns_servers": {
							Type:     pluginsdk.TypeList,
							Required: true,
							Elem: &pluginsdk.Schema{
								Type:         pluginsdk.TypeString,
								ValidateFunc: validate.IPv4Address,
							},
						},
						"domain": {
							Type:     pluginsdk.TypeString,
							Required: true,
							ValidateFunc: validation.StringMatch(
								regexp.MustCompile(`^[(\da-zA-Z-).]{1,255}$`),
								`The domain name must end with a letter or number before dot and start with a letter or number after dot and can not be longer than 255 characters in length.`,
							),
						},
						"smb_server_name": {
							Type:     pluginsdk.TypeString,
							Required: true,
							ValidateFunc: validation.StringMatch(
								regexp.MustCompile(`^[\da-zA-Z]{1,10}$`),
								`The smb server name can not be longer than 10 characters in length.`,
							),
						},
						"username": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"password": {
							Type:         pluginsdk.TypeString,
							Required:     true,
							Sensitive:    true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"organizational_unit": {
							Type:     pluginsdk.TypeString,
							Optional: true,
						},
					},
				},
			},

			"tags": tags.Schema(),
		},
	}
}

func resourceNetAppAccountCreateUpdate(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).NetApp.AccountClient
	subscriptionId := meta.(*clients.Client).Account.SubscriptionId
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id := parse.NewAccountID(subscriptionId, d.Get("resource_group_name").(string), d.Get("name").(string))
	if d.IsNewResource() {
		existing, err := client.Get(ctx, id.ResourceGroup, id.NetAppAccountName)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing %s: %+v", id, err)
			}
		}
		if !utils.ResponseWasNotFound(existing.Response) {
			return tf.ImportAsExistsError("azurerm_netapp_account", id.ID())
		}
	}

	location := azure.NormalizeLocation(d.Get("location").(string))
	activeDirectories := d.Get("active_directory").([]interface{})

	accountParameters := netapp.Account{
		Location: utils.String(location),
		AccountProperties: &netapp.AccountProperties{
			ActiveDirectories: expandNetAppActiveDirectories(activeDirectories),
		},
		Tags: tags.Expand(d.Get("tags").(map[string]interface{})),
	}

	future, err := client.CreateOrUpdate(ctx, accountParameters, id.ResourceGroup, id.NetAppAccountName)
	if err != nil {
		return fmt.Errorf("creating %s: %+v", id, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for the creation of %s: %+v", id, err)
	}

	d.SetId(id.ID())
	return resourceNetAppAccountRead(d, meta)
}

func resourceNetAppAccountRead(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).NetApp.AccountClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AccountID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceGroup, id.NetAppAccountName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			log.Printf("[INFO] %s does not exist - removing from state", *id)
			d.SetId("")
			return nil
		}
		return fmt.Errorf("retrieving %s: %+v", *id, err)
	}

	d.Set("name", id.NetAppAccountName)
	d.Set("resource_group_name", id.ResourceGroup)

	if location := resp.Location; location != nil {
		d.Set("location", azure.NormalizeLocation(*location))
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func resourceNetAppAccountDelete(d *pluginsdk.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).NetApp.AccountClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AccountID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceGroup, id.NetAppAccountName)
	if err != nil {
		return fmt.Errorf("deleting %s: %+v", *id, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for deletion of %s: %+v", *id, err)
	}

	return nil
}

func expandNetAppActiveDirectories(input []interface{}) *[]netapp.ActiveDirectory {
	results := make([]netapp.ActiveDirectory, 0)
	for _, item := range input {
		v := item.(map[string]interface{})
		dns := strings.Join(*utils.ExpandStringSlice(v["dns_servers"].([]interface{})), ",")

		result := netapp.ActiveDirectory{
			DNS:                utils.String(dns),
			Domain:             utils.String(v["domain"].(string)),
			OrganizationalUnit: utils.String(v["organizational_unit"].(string)),
			Password:           utils.String(v["password"].(string)),
			SmbServerName:      utils.String(v["smb_server_name"].(string)),
			Username:           utils.String(v["username"].(string)),
		}

		results = append(results, result)
	}
	return &results
}
