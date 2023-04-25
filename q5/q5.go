package q5

import (
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	// Seu código aqui
	correta := ""
	correta = strings.ReplaceAll(s, "A", "")
	correta = strings.ReplaceAll(correta, "E", "")
	correta = strings.ReplaceAll(correta, "I", "")
	correta = strings.ReplaceAll(correta, "O", "")
	correta = strings.ReplaceAll(correta, "U", "")
	correta = strings.ReplaceAll(correta, "a", "")
	correta = strings.ReplaceAll(correta, "e", "")
	correta = strings.ReplaceAll(correta, "i", "")
	correta = strings.ReplaceAll(correta, "o", "")
	correta = strings.ReplaceAll(correta, "u", "")

	var nova string
	for i := 0; i < len(correta); i++ {
		nova += "."
		nova += string(correta[i])

	}
	minuscula := strings.ToLower(nova)
	return minuscula
}
