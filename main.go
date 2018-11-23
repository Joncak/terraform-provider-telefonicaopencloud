package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/joncak/terraform-provider-telefonicaopencloud/telefonicaopencloud"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: telefonicaopencloud.Provider})
}
