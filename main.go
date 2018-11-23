package main

import (
	"github.com/Joncak/terraform-provider-telefonicaopencloud/telefonicaopencloud"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: telefonicaopencloud.Provider})
}
