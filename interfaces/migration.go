package interfaces

type Migration interface {
	AddSql(sql string)
}
