package database

_ "github.com/lib/pq"

type Sql struct{
	Db*sql.DB
Host string
Porr string
Username string
Password string
DbName string

}
func(s**Sql)Connect(){
	dataSoure := fmt.Sprint("host=%s")
}
