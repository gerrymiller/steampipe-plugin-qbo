package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/Cloudticity/steampipe-plugin-qbo/qbo"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: qbo.Plugin})
}