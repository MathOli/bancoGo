package main

import (
	"banco/clientes"
	"banco/contas"
	// "fmt"
)

type verificaConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificaConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

func main() {

	joao := clientes.Titular{Nome: "João Bonifácio", CPF: "123.123.123-12", Profissao: "Engenheiro de Software"}
	conta1 := contas.ContaCorrente{
		Titular: joao,
		Agencia: 589,
		Conta:   123456,
	}

	marcela := clientes.Titular{Nome: "Marcela Lima", CPF: "321.321.321-32", Profissao: "Engenheira de Software"}
	conta2 := contas.ContaPoupanca{
		Titular: marcela,
		Agencia: 589,
		Conta:   123456,
	}

	conta1.Depositar(1000)
	conta2.Depositar(1000)

	// fmt.Println(conta1.Titular.Nome)
	// fmt.Println(conta1.VerificarSaldo())
	// fmt.Println(conta2.Titular.Nome)
	// fmt.Println(conta2.VerificarSaldo())
	// fmt.Println(conta1.Transferir(100, &conta2))
	// fmt.Println(conta1.Titular.Nome)
	// fmt.Println(conta1.VerificarSaldo())
	// fmt.Println(conta2.Titular.Nome)
	// fmt.Println(conta2.VerificarSaldo())

}
