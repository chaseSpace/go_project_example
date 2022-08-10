package config

type AppConf struct {
	MysqlDsn string
}

func NewConf(fpath string) (*AppConf, error) {
	// TODO: use fpath to parse conf
	return new(AppConf), nil
}
