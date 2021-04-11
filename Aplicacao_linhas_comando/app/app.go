package app

import "github.com/urfave/cli"


//Gerar retorna a aplicação pronta pra ser executada
func Gerar() *cli.App  {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca de IPs e nomes de servidor na internet"

	return app
}
