package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var tagName string
	fmt.Scan(&tagName)

	// fmt.Println(tagName)
	refsTag := fmt.Sprintf("refs/tags/%s", tagName)
	branchName := fmt.Sprintf("tag-%s", tagName)
	cmd := exec.Command("git", "checkout", "-b", branchName, refsTag)
	cmd.Run()
}
