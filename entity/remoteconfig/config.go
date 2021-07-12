package remoteconfig

const DatabaseConfigPath = "/main/db"

type Database struct {
	User   string            `tars:"/main/db<user>"`
	Pass   string            `tars:"/main/db<pass>"`
	Host   string            `tars:"/main/db<host>"`
	DBName string            `tars:"/main/db<dbname>"`
	Params map[string]string `tars:"/main/db<param>"`
}
