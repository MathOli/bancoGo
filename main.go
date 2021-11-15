package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	conta1 := contas.ContaCorrente{
		Titular: "Matheus",
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	conta2 := contas.ContaCorrente{
		Titular: "Marcela",
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	fmt.Println(conta1.Titular)
	fmt.Println(conta1.Saldo)
	fmt.Println(conta2.Titular)
	fmt.Println(conta2.Saldo)
	fmt.Println(conta1.Transferir(100, &conta2))
	fmt.Println(conta1.Titular)
	fmt.Println(conta1.Saldo)
	fmt.Println(conta2.Titular)
	fmt.Println(conta2.Saldo)

}
