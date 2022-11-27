package provider

// import (
// 	"context"

// 	// "github.com/hashicorp/terraform-plugin-log/tflog"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
// 	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
// )

// const VALID_TYPES = "primary secondary forward"

// func resourceZone() *schema.Resource {
// 	return &schema.Resource{
// 		Description: "Represents a zone in host_vars",

// 		CreateContext: Noop,
// 		ReadContext:   Noop,
// 		UpdateContext: Noop,
// 		DeleteContext: Noop,

// 		Schema: map[string]*schema.Schema{
// 			"allow_update": {
// 				Description: "A list of hosts that are allowed to dynamically update this DNS zone",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"also_notify": {
// 				Description: "A list of servers that will receive a notification when the primary zone file is reloaded",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"create_forward_zone": {
// 				Description: "When initialized and set to false, creation of forward zones will be skipped (resulting in a reverse only zone)",
// 				Type:        schema.TypeBool,
// 				Optional:    true,
// 			},
// 			"create_reverse_zone": {
// 				Description: "When initialized and set to false, creation of reverse zones will be skipped (resulting in a forward only zone)",
// 				Type:        schema.TypeBool,
// 				Optional:    true,
// 			},
// 			"delegate": {
// 				Description: "Zone delegation servers",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"forwarders": {
// 				Description: "List of forwarders for for the forward type zone",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"hostmaster_email": {
// 				Description: "The e-mail address of the system administrator for the zone",
// 				Type:        schema.TypeString,
// 				Optional:    true,
// 			},
// 			"ipv6_networks": {
// 				Description: "A list of the IPv6 networks that are part of the domain, in CIDR notation (e.g. 2001:db8::/48)",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			// "mail_servers": {
// 			// 	Description: "A list of mappings (with keys name: and preference:) specifying the mail servers for this domain.",
// 			// 	Type:        schema.TypeList,
// 			// 	Optional:    true,
// 			// 	Elem: &schema.Resource{
// 			// 		"name": {
// 			// 			Description: "Hostname of the Mail Server",
// 			// 			Type:        schema.TypeString,
// 			// 			Required:    true,
// 			// 		},
// 			// 		"preference": {
// 			// 			Description: "The priority for the Mail Server",
// 			// 			Type:        schema.TypeInt,
// 			// 			Required:    true,
// 			// 		},
// 			// 	},
// 			},
// 			"name_servers": {
// 				Description: "A list of the nameservers for this zone",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"name": {
// 				Description: "The domain name",
// 				Type:        schema.TypeString,
// 				Required:    true,
// 			},
// 			"naptr": {
// 				Description: "A list of mappings with keys name:, order:, pref:, flags:, service:, regex: and replacement: specifying NAPTR records.",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					"name": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"order": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"pref": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"flags": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"service": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"regex": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"replacement": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 				},
// 			},
// 			"networks": {
// 				Description: "A list of networks that are part of this domain",
// 				Type:        schema.TypeList,
// 				Required:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"other_name_servers": {
// 				Description: "A list of the DNS servers outside of this domain.",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"primaries": {
// 				Description: "A list of primary DNS servers for this zone.",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"services": {
// 				Description: "A list of services to be advertised by SRV records",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					Type: schema.TypeString,
// 				},
// 			},
// 			"text": {
// 				Description: "A list of mappings with keys name: and text:, specifying TXT records. text: can be a list or string.",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					"name": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"text": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 				},
// 			},
// 			"caa": {
// 				Description: "A list of mappings with keys name: and text:, specifying CAA records. text: can be a list or string",
// 				Type:        schema.TypeList,
// 				Optional:    true,
// 				Elem: &schema.Schema{
// 					"name": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 					"text": {
// 						Type:     schema.TypeString,
// 						Required: true,
// 					},
// 				},
// 			},
// 		},
// 	}
// }

// func Noop(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
// 	return nil
// }
