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

	conta1 := contaCorrente{
		Titular: "Matheus",
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	conta2 := contaCorrente{
		Titular: "Marcela",
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	fmt.Println(conta1.Titular)
	fmt.Println(conta1.Saldo)
	fmt.Println(conta2.Titular)
	fmt.Println(conta2.Saldo)
	fmt.Println(conta1.transferir(100, &conta2))
	fmt.Println(conta1.Titular)
	fmt.Println(conta1.Saldo)
	fmt.Println(conta2.Titular)
	fmt.Println(conta2.Saldo)

}

func (c *contaCorrente) transferir(valorTransferencia float64, contaDestino *contaCorrente) string {

	podeTransferir := valorTransferencia <= c.Saldo && valorTransferencia > 0

	if podeTransferir {
		c.Saldo -= valorTransferencia
		contaDestino.Saldo += valorTransferencia
		return "Transferencia realizada com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *contaCorrente) depositar(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.Saldo += valor
		return "Deposito realizado com sucesso"
	} else {
		return "Valor invalido"
	}
}

func (c *contaCorrente) sacar(valorSaque float64) (string, float64) {

	podeSacar := valorSaque <= c.Saldo && valorSaque > 0

	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}

}
