package provider

import (
	"bytes"
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
)

func resourceKey() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceKeyCreate,
		ReadContext:   resourceKeyRead,
		DeleteContext: resourceKeyDelete,
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"comment": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key_armor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_key_armor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKeyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {

	email := d.Get("email").(string)
	name := d.Get("name").(string)
	comment := d.Get("comment").(string)

	ent, err := openpgp.NewEntity(name, comment, email, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	buf := &bytes.Buffer{}

	w, err := armor.Encode(buf, openpgp.PrivateKeyType, make(map[string]string))
	if err != nil {
		return diag.FromErr(err)
	}
	if err := ent.PrivateKey.Serialize(w); err != nil {
		return diag.FromErr(err)
	}
	d.Set("private_key_armor", buf.String())

	buf.Reset()

	w, err = armor.Encode(buf, openpgp.PublicKeyType, make(map[string]string))
	if err != nil {
		return diag.FromErr(err)
	}
	if err := ent.PrimaryKey.Serialize(w); err != nil {
		return diag.FromErr(err)
	}
	d.Set("public_key_armor", buf.String())

	d.Set("id", ent.PrimaryKey.KeyIdString())
	d.SetId(ent.PrimaryKey.KeyIdString())

	return diags
}

func resourceKeyRead(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	return
}

func resourceKeyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	d.SetId("")
	return
}
