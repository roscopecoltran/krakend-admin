package gateway

//import (
//	"github.com/jinzhu/gorm"
//)

type EndpointTag struct {
	EndpointID uint `gorm:"unique_index:idx_endpoint_id_tag_id"`
	TagID      uint `gorm:"unique_index:idx_endpoint_id_tag_id"`
}

type BackendTag struct {
	BackendID uint `gorm:"unique_index:idx_backend_id_tag_id"`
	TagID     uint `gorm:"unique_index:idx_backend_id_tag_id"`
}

type HostTag struct {
	HostID uint `gorm:"unique_index:idx_host_id_tag_id"`
	TagID  uint `gorm:"unique_index:idx_host_id_tag_id"`
}

type ServiceTag struct {
	ServiceID uint `gorm:"unique_index:idx_service_id_tag_id"`
	TagID     uint `gorm:"unique_index:idx_service_id_tag_id"`
}

type PathMXJTag struct {
	PathMxjID uint `gorm:"unique_index:idx_path_mxj_id_tag_id"`
	TagID     uint `gorm:"unique_index:idx_path_mxj_id_tag_id"`
}
