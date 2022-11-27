package provider

import (
	"context"
	"fmt"

	// "github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Description: "Represents a host entry in a domain as a hostvar",

		CreateContext: resourceHostCreate,
		ReadContext:   Noop,
		UpdateContext: Noop,
		DeleteContext: Noop,

		Schema: map[string]*schema.Schema{
			"inventory_hostname": {
				Description: "The inventory hostname defined using nbering/ansible, or manually in a static inventory",
				Type	   : schema.TypeString,
				Required   : true,
			},
			"domain": {
				Description: "The fqn of the domain this host is in",
				Type:        schema.TypeString,
				Required:    true,
			},
			"hostname": {
				Description: "The hostname (only) for this host",
				Type:        schema.TypeString,
				Required:    true,
			},
			"ip": {
				Description: "The IP Address for this host",
				Type:        schema.TypeString,
				Required:    true,
			},
			"aliases": {
				Description: "A list of aliases for this host",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func resourceHostCreate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	var id = fmt.Sprintf("%s.%s", d.Get("hostname").(string), d.Get("domain").(string))
	d.SetId(id)
	return nil
}

func Noop(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}