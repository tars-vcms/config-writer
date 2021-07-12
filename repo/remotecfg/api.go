package remotecfg

type RemoteConfig interface {
	GetDatabaseDSN() (string, error)
}

func New() RemoteConfig {
	return newRemoteConfigImpl()
}
