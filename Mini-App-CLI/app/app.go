package app

import "github.com/urfave/cli"

//Gerar retorna ao app de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca por ips e servidores na internet"
	return app
}
