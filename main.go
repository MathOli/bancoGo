package main

import (
	"fmt"
)

type contaCorrente struct {
	Titular string
	Agencia int
	Conta   int
	Saldo   float64
}

func main() {

	contaMatheus := contaCorrente{
		Titular: "Matheus",
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	conta := fmt.Sprintf("Nome: %s \nAg:%d \nCc: %d \nSaldo: %.2f", contaMatheus.Titular, contaMatheus.Agencia, contaMatheus.Conta, contaMatheus.Saldo)

	fmt.Println(conta)

}
