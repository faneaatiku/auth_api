package connector

import "database/sql"

type Connections struct {
	Mysql *sql.DB
}

func NewConnections() (*Connections, error) {
	sqlConn, err := GetMysqlConnection()
	if err != nil {
		return nil, err
	}

	return &Connections{Mysql: sqlConn}, nil
}
