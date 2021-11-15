package contas

import (
	"banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	Agencia int
	Conta   int
	Saldo   float64
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) string {

	podeTransferir := valorTransferencia <= c.Saldo && valorTransferencia > 0

	if podeTransferir {
		c.Saldo -= valorTransferencia
		contaDestino.Saldo += valorTransferencia
		return "Transferencia realizada com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.Saldo += valor
		return "Deposito realizado com sucesso"
	} else {
		return "Valor invalido"
	}
}

func (c *ContaCorrente) Sacar(valorSaque float64) (string, float64) {

	podeSacar := valorSaque <= c.Saldo && valorSaque > 0

	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso", c.Saldo
	} else {
		return "Saldo insuficiente", c.Saldo
	}

}
