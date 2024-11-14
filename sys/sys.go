package sys

import "agent/impl/sysimpl"
import "fmt"

func GetApps() {
	fmt.Println("Get Apps Common Code")
	sysimpl.GetAppsImpl()
}
