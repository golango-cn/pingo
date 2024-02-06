package main

import (
	"log"
	"testing"

	"github.com/golango-cn/pingo"
)

func TestMain(t *testing.T) {
	// Make a new plugin from the executable we created. Connect to it via TCP
	p := pingo.NewPlugin("tcp", "/mnt/git/github.com/pingo/examples/pingo-hello-world/pingo-hello-world")
	// Actually start the plugin
	p.Start()
	// Remember to stop the plugin when done using it
	defer p.Stop()

	var resp string

	// Call a function from the object we created previously
	if err := p.Call("Plugin.SayHello", "Go developer", &resp); err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}
