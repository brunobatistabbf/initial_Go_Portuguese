package banco

import "database/sql"

//Conectar abre a conexão com banco de dados
func Conectar() (*sql.DB, error) {
	stringConexao := "golang:golang@/devbook?charset=utf8f&parseTIme=True&loc=Local"

	db, erro := sql.Open("postgres", stringConexao)
	if erro != nil {
		return nil, erro
	}
	if erro = db.Ping(); erro != nil {
		return nil, erro
	}
	return db, nil

}
