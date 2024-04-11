package main

import (
	"errors"
	"fmt"
)

/*
Exercício 1: Nível Básico
Descrição:Escreva uma função chamada trocar que recebe dois ponteiros para inteiros e
		  troca os valores das variáveis apontadas pelos ponteiros.

*/

/*
Exercício 2: Nível Intermediário
Descrição:Crie uma estrutura chamada Pessoa que represente uma pessoa com os campos Nome (string) e Idade (int).
		  Escreva uma função chamada alterarIdade que recebe um ponteiro para uma pessoa e um valor inteiro,
		  e atualiza a idade da pessoa para o valor fornecido.
*/

// Definição da estrutura Pessoa
type Pessoa struct {
	Nome  string
	Idade int
}

// Função alterarIdade recebe um ponteiro para Pessoa e uma nova idade, atualizando o campo Idade
func alterarIdade(p *Pessoa, novaIdade int) {
	p.Idade = novaIdade
}

// Função trocar recebe dois ponteiros para inteiros e troca os valores das variáveis apontadas pelos ponteiros
// Realiza a troca de valores entre dois inteiros usando ponteiros.
// Se um ou ambos os ponteiros forem nulos, a função retorna um erro
// indicando a ocorrência de referência nula (nil pointer dereference).
// Para evitar esse erro, verificamos se os ponteiros são diferentes de nil
// antes de realizar a troca de valores.
func trocar(a, b *int) error {
	if a == nil || b == nil {
		return errors.New("Um ou ambos os ponteiros são nulos")
	}
	*a, *b = *b, *a
	return nil
}

func main() {
	fmt.Println("------------------------ EXERCÍCIO 1 ------------------------")

	// Declaração de duas variáveis inteiras x e y
	x := 10
	y := 20

	// Impressão dos valores antes da troca
	fmt.Println("Antes da troca")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	// Chamada da função trocar passando os endereços de x e y
	err := trocar(&x, nil)

	if err != nil {
		fmt.Println("\nErro ao trocar valores:", err, "\n")
	}

	// Impressão dos valores depois da troca
	fmt.Println("Depois da troca")
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("------------------------ EXERCÍCIO 2 ------------------------")

	// Declaração de uma variável do tipo Pessoa
	pessoa := Pessoa{
		Nome:  "Leonardo",
		Idade: 38,
	}

	// Impressão da Pessoa antes da alteração
	fmt.Println("Antes da alteração:", pessoa)

	// Chamada da função alterarIdade passando o endereço da pessoa e a nova idade
	alterarIdade(&pessoa, 40)

	// Impressão da Pessoa depois da alteração
	fmt.Println("Antes da alteração:", pessoa)

}
