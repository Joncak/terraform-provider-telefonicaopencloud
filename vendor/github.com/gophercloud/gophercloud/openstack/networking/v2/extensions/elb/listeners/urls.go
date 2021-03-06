package listeners

import "github.com/gophercloud/gophercloud"

const (
	rootPath     = "elbaas"
	resourcePath = "listeners"
)

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath)
}

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(c.ProjectID, rootPath, resourcePath, id)
}
