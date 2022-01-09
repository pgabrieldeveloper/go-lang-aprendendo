package endereco

import "testing"

type cenarioTest struct {
	enderecoRecebido  string
	resultadoEsperado string
}

func TestTipeDeEndereco(t *testing.T) {

	//Estando teste com apenas 1 cenario

	//enderecoTest := "RUA DOS BOBOS NUMERO 0"
	//resultadoEsperado := "Rua"
	//resultadoRecebito := TipeDeEndereco(enderecoTest)
	//
	//if resultadoRecebito != resultadoEsperado {
	//	t.Error("Resultado esperado é diferente do rececbido")
	//}

	//Testando com varios cenarios
	cenarios := []cenarioTest{
		{"Rua dos bobos numero 0", "Rua"},
		{"Avenida Paulias", "Avenida"},
		{"Estrada Pinheiro Campo", "Estrada"},
		{"", "tipoInvalido"},
		{"Rodovia alameda spring", "Rodovia"},
	}
	for _, cenario := range cenarios {
		resultadoRecebido := TipeDeEndereco(cenario.enderecoRecebido)
		if resultadoRecebido != cenario.resultadoEsperado {
			t.Errorf("Resultado Esperado: %s Resultado Recebido: %s", cenario.resultadoEsperado, resultadoRecebido)
		}
	}

}
