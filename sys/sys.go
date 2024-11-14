package sys

type SysInterface interface {
	GetApps()
}

// SysInfoService is the implementation of SysInfo Service Interface
type SysService struct{}

func (r *SysService) GetApps() {
	GetApps()
}
