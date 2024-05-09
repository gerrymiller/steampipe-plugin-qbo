package main

import (
	"github.com/Cloudticity/steampipe-plugin-qbo/qbo"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: qbo.Plugin})
}
