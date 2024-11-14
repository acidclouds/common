package com

import "fmt"
import "agent/common/sys"

func CommFunc() {
	fmt.Println("Common Comm code")
	sys.GetApps()
}
