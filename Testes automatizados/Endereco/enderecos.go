package endereco

import "strings"

//TipeDeEndereco se o tipo de ennderço escrito é valido
func TipeDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	primeiraLetra := strings.Split(strings.ToLower(endereco), " ")[0]

	contem := false

	for _, tipo := range tiposValidos {
		if strings.ToLower(tipo) == primeiraLetra {
			contem = true
		}
	}
	if contem {
		return strings.Title(primeiraLetra)
	}
	return "tipoInvalido"

}
