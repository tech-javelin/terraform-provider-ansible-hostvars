package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceHost() *schema.Resource {
	return &schema.Resource{
		Description: "Represents a host entry in a domain as a hostvar",

		CreateContext: resourceHostCreate,
		ReadContext:   resourceHostRead,
		UpdateContext: resourceHostUpdate,
		DeleteContext: resourceHostDelete,

		Schema: map[string]*schema.Schema{
			"domain": {
				Description: "The fqn of the domain this host is in",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"hostname": {
				Description: "The hostname (only) for this host",
				Type:        schema.TypeString,
				Optional:    false,
			},
			"ip": {
				Description: "The IP Address for this host",
				Type:        schema.TypeString,
				Optional:    false,
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
	return diag.Errorf("not implemented")
}

func resourceHostRead(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return diag.Errorf("not implemented")
}

func resourceHostUpdate(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return diag.Errorf("not implemented")
}

func resourceHostDelete(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return diag.Errorf("not implemented")
}
