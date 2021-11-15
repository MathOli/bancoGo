package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular   clientes.Titular
	Agencia   int
	Conta     int
	saldo     float64
	Operacoes int
}

func (c *ContaPoupanca) VerificarSaldo() float64 {
	return c.saldo
}

func (c *ContaPoupanca) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) string {

	podeTransferir := valorTransferencia <= c.saldo && valorTransferencia > 0

	if podeTransferir {
		c.saldo -= valorTransferencia
		contaDestino.saldo += valorTransferencia
		return "Transferencia realizada com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valor float64) string {
	podeDepositar := valor > 0
	if podeDepositar {
		c.saldo += valor
		return "Deposito realizado com sucesso"
	} else {
		return "Valor invalido"
	}
}

func (c *ContaPoupanca) Sacar(valorSaque float64) (string, float64) {

	podeSacar := valorSaque <= c.saldo && valorSaque > 0

	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso", c.saldo
	} else {
		return "saldo insuficiente", c.saldo
	}

}
