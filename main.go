package main

import (
	"fmt"
	"my-budgety-pub-sub/utils"
)

func main() {
	envVariables := utils.GetEnvVariables()
	fmt.Println("Test Env Variable:", envVariables.TestVariable)
}
