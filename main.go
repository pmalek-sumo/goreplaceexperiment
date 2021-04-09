package main

import (
	"fmt"

	"go.opentelemetry.io/collector/processor/filterprocessor"
)

func main() {
	f := filterprocessor.NewFactory()
	_ = f
	fmt.Printf("name %#v\n", f)
	// Doesn't compile because the local version of factory doesn't implement
	// CreateDefaultConfig()
	// fmt.Printf("name %v\n", f.CreateDefaultConfig().Name())
}
