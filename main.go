package main

import (
	"flag"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	cmdutil "k8s.io/kubectl/pkg/cmd/util"
)

func main() {

	var res string

	flag.StringVar(&res, "res", "", "Resource you want to interact with")
	flag.Parse()
	configFlag := genericclioptions.NewConfigFlags(true).WithDeprecatedPasswordFlag()

	matchVersionFlags := cmdutil.NewMatchVersionFlags(configFlag)

	mapper, err := cmdutil.NewFactory(matchVersionFlags).ToRESTMapper()

	if err != nil {
		fmt.Printf("error while creating rest mapper - %s", err.Error())
	}

	groupVersionResource, err := mapper.ResourceFor(schema.GroupVersionResource{
		Resource: res,
	})

	if err != nil {
		fmt.Printf("Error while getting GVR from RestMapper - %s", err.Error())
	}

	fmt.Printf("Complete GVR is Group : %s, Version: %s, Resource: %s\n ",
		groupVersionResource.Group, groupVersionResource.Version, groupVersionResource.Resource)

}
