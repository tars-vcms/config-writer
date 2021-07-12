package remotecfg

import (
	"errors"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/conf"
	"github.com/tars-vcms/config-writer/entity/remoteconfig"
	"net/url"
	"reflect"
)

type RemoteConfigImpl struct {
	remoteConf *conf.Conf
}

func (r RemoteConfigImpl) GetDatabaseDSN() (string, error) {
	database := &remoteconfig.Database{
		User:   "",
		Pass:   "",
		Host:   "",
		DBName: "",
		Params: map[string]string{},
	}
	t := reflect.TypeOf(database).Elem()
	v := reflect.ValueOf(database).Elem()
	for i := 0; i < t.NumField(); i++ {
		tf := t.Field(i)
		vf := v.Field(i)
		path := tf.Tag.Get("tars")
		switch tf.Type.Kind() {
		case reflect.String:
			vf.SetString(r.remoteConf.GetString(path))
			break
		case reflect.Map:
			for k, v := range r.remoteConf.GetMap(path) {
				vf.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
			}
			break
		default:
			return "", errors.New("unknown database params type")
		}

	}
	q := url.Values{}
	for k, v := range database.Params {
		q.Set(k, v)
	}
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?%v", database.User, database.Pass, database.Host, database.DBName, q.Encode())
	return dsn, nil
}

func newRemoteConfigImpl() *RemoteConfigImpl {
	//确保tars已经初始化
	tars.GetServerConfig()
	tarsConf, err := conf.NewConf(tars.ServerConfigPath)
	if err != nil {
		panic("[Tars配置获取失败]" + err.Error())
	}
	return &RemoteConfigImpl{
		remoteConf: tarsConf,
	}
}
