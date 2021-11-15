package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	joao := clientes.Titular{Nome: "João Bonifácio", CPF: "123.123.123-12", Profissao: "Engenheiro de Software"}
	conta1 := contas.ContaCorrente{
		Titular: joao,
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	marcela := clientes.Titular{Nome: "Marcela Lima", CPF: "321.321.321-32", Profissao: "Engenheira de Software"}
	conta2 := contas.ContaCorrente{
		Titular: marcela,
		Agencia: 589,
		Conta:   123456,
		Saldo:   1000.00,
	}

	fmt.Println(conta1.Titular.Nome)
	fmt.Println(conta1.Saldo)
	fmt.Println(conta2.Titular.Nome)
	fmt.Println(conta2.Saldo)
	fmt.Println(conta1.Transferir(100, &conta2))
	fmt.Println(conta1.Titular.Nome)
	fmt.Println(conta1.Saldo)
	fmt.Println(conta2.Titular.Nome)
	fmt.Println(conta2.Saldo)

}
