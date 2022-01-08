package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

//Gerar retorna ao app de linha de comando
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca por ips e servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "www.atoscapital.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca por ips na internet",
			Flags:  flags,
			Action: buscarPorIp,
		},
		{
			Name:   "servidores",
			Usage:  "Busca por Servidores na internet",
			Flags:  flags,
			Action: buscarPorServidores,
		},
	}
	return app
}

func buscarPorIp(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {

		fmt.Println(ip)
	}
}

func buscarPorServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host) // ns Nome servidor:
	if erro != nil {
		log.Fatal(erro)
	}
	for _, servidor := range servidores {

		fmt.Println(servidor.Host)
	}
}
