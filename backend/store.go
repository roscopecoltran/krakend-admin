package backend

type StoreConfig struct {
	Admin   AdminConfig
	Bolt    BoltStorageConfig
	Storm   StormStorageConfig
	Gorm    GormStorageConfig
	IsReady bool
}

var Store StoreConfig
