package main

import (
	"fmt"

	"github.com/karchx/nrs/cmd"
	"github.com/karchx/nrs/utils"
)

func main() {
  project := utils.GetProject(".")
  fmt.Printf("%v: ", project)
	cmd.Execute()
}
