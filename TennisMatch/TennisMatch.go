package main

import (
	"fmt"
	"math/rand"
	"time"
)

var placar [3][2]int               //Placar do jogo
var points, games, sets = 60, 1, 1 //Regras de pontuação do jogo. Podem ser alteradas pelo usuário (ver funções tela e regras)

// Funções auxiliares
func tela() {
	//Essa função simula uma espécie de interface gráfica

	var resposta string

	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("                              T E N N I S        M A T C H")
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("                                        Bem-vindo!")
	fmt.Println("> Por padrão, será disputado 1 set, composto por 1 game, onde um jogador precisa alcançar\n" +
		"  a pontuação 60 para terminar o game. Você gostaria de mudar essas regras?")
	fmt.Scanln(&resposta)

	//Se o usuário desejar alterar as regras...
	if resposta == "s" || resposta == "S" || resposta == "sim" || resposta == "Sim" || resposta == "SIM" {
		regras()
	}

	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("                                   Começando novo jogo...")
	fmt.Println("-----------------------------------------------------------------------------------------")
}

func regras() {
	//Essa função altera as regras de pontos, games e sets necessários para uma partida de tênis. Por padrão, temos
	//definidos os valores 60, para pontos, 1, para games, e 1 para sets.

	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("Número de pontos para vencer o game: ")
	fmt.Scanln(&points)
	fmt.Println("Número mínimo de games para vencer o set: ")
	fmt.Scanln(&games)
	fmt.Println("Número mínimo de sets para vencer a partida: ")
	fmt.Scanln(&sets)
}

func marcou(player int) {
	//Essa função simplesmente informa que o PLAYER marcou um ponto e atualiza e imprime o placar

	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println(">>> PLAYER " + string(player) + " marcou ponto.")
	atualizaPlacar(player)
	fmt.Printf("   | PLAYER 1 | %d | %d | %d |\n", placar[0][0], placar[1][0], placar[2][0])
	fmt.Printf("   | PLAYER 2 | %d | %d | %d |\n", placar[0][1], placar[1][1], placar[2][1])
	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Println("									   NOVA JOGADA                                        ")
	fmt.Println("-----------------------------------------------------------------------------------------")
}

func atualizaPlacar(player int) bool {
	//Essa função atualiza a pontuação de uma player. Retorna um valor boolean que é positivo quando a partida é encerrada
	//e negativo, caso contrário.
	//Essa função é um megazord. Ela descreve todos os cenários de pontuação possíveis numa partida de tênis para dois
	//jogadores. Porém, para possibilitar a compreensão, ela está toda documentada com a descrição de cada cenário e ação,
	//e está minimizada, para facilitar a visualização.

	/* placar[0][x] : Equivale ao número de pontos do PLAYER x
	   placar[1][x] : Equivale ao número de games do PLAYER x
	   placar[2][x] : Equivale ao número de sets do PLAYER x
	*/

	if player == 1 { //CENÁRIOS PARA O PLAYER 1

		if placar[2][0] == games-1 {
			/* CENÁRIO DE SETS 1
			Caso ele esteja a um set da vitória...
			*/

			if /*CENÁRIO 1*/ placar[1][0] == games-1 && placar[1][1] < games-1 || /*CENÁRIO 2*/ placar[1][0] >= games && placar[1][1] < placar[1][0] {
				/* CENÁRIO DE GAMES 1
					Caso ele esteja a um game de abrir uma diferença de dois games para seu adversário e atingir o mínimo
					de games para fechar o set.

				   CENÁRIO DE GAMES 2
					Caso ele esteja a apenas um game de vantagem do adversário, mesmo já tendo atingido a quantidade
				    mínima de games para fechar o set. Pela regra do tênis, é necessário tem uma vantagem de,
				    pelo menos, dois games para fechar o set.
				*/

				if placar[0][0] == 45 && placar[0][1] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 1
					Caso ele esteja em vantagem no último ponto. Neste cenário, era um match point. Como ele marcou,
					venceu o jogo.
					*/
					placar[2][0]++ //Aumenta a contagem de sets em 1
					fmt.Println("-----------------------------------------------------------------------------------------")
					fmt.Println(">>> PLAYER 1 venceu o jogo por " + string(placar[2][0]) + "x" + string(placar[2][1]) + ".")
					fmt.Println("-----------------------------------------------------------------------------------------")
					return true

				} else if placar[0][0] == 45 && placar[0][1] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 2
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][0] += 5 //Aumenta a contagem de pontos em 5, simbolizando o ADVANTAGE

				} else if placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 3
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como estava a um game de fechar o set, também fecha o set. E, como estava também a um set de ganhar
					o jogo, ele vence a partida.
					*/
					placar[2][0]++ //Aumenta a contagem de sets em 1
					fmt.Println("-----------------------------------------------------------------------------------------")
					fmt.Println(">>> PLAYER 1 venceu o jogo por " + string(placar[2][0]) + "x" + string(placar[2][1]) + ".")
					fmt.Println("-----------------------------------------------------------------------------------------")
					return true

				} else if placar[0][0] == 45 && placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 4
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][1] -= 5 //Diminui a contagem de pontos do adversário em 5, pois perdeu o ADVANTAGE

				} else if placar[0][0] > 0 && placar[0][0] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][0] += 15 //Aumenta a contagem de pontos em 5
				}

			} else {
				/* CENÁRIO DE GAMES 3, 4 E 5
				Está cláusula "else" inclui os outros três cenários possíveis, descritos abaixo, que executam as mesmas ações

				CENÁRIO DE GAMES 3
				 Os PLAYERs estão empatados no número de games. Mesmo que vencer um game signifique que o PLAYER atingiu
				 o número mínimo de games necessários para fechar o set, pela regra do tênis, ele precisa de uma
				 vantagem de pelo menos dois games para fechar o set. Por isso, é apenas adicionada uma unidade a sua
				 contagem de sets.

				CENÁRIO DE GAMES 4
				 O PLAYER está em desvantagem no número de games, mesmo os dois já tendo atingido a quantidade mínima de
				 games. Pela mesma regra descrita acima, seu adversário necessitaria de dois games de vantagens para
				 fechar o set. Por isso, o set ainda não foi encerrado e uma unidade é adicionada a sua contagem de sets.

				CENÁRIO DE GAMES 5
				 Cenário normal, onde o PLAYER ainda não atingiu a quantidade necessária de games para fechar o set.
				*/

				if placar[0][0] == 45 && placar[0][1] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 5
					Caso ele esteja em vantagem no último ponto para fechar o game. Neste cenário, não era um match point,
					pois ganhando este game, o PLAYER não fecha o set. Portanto, foi apenas adicionada uma unidade a
					sua contagem de games. Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][0]++
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][0] == 45 && placar[0][1] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 6
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][0] += 5

				} else if placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 7
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como não estava a um game de fechar o set, apenas uma unidade é adicionada a sua contagem de games.
					Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][0]++
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][0] == 45 && placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 8
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][1] -= 5

				} else if placar[0][0] > 0 && placar[0][0] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][0] += 15
				}
			}

		} else if placar[2][0] < games-1 {
			/* CENÁRIO DE SETS 2
			Caso ele não esteja a um set da vitória...
			*/

			if /*CENÁRIO 1*/ placar[1][0] == games-1 && placar[1][1] < games-1 || /*CENÁRIO 2*/ placar[1][0] >= games && placar[1][1] < placar[1][0] {
				/* CENÁRIO DE GAMES 1
					Caso ele esteja a um game de abrir uma diferença de dois games para seu adversário e atingir o mínimo
					de games para fechar o set.

				   CENÁRIO DE GAMES 2
					Caso ele esteja a apenas um game de vantagem do adversário, mesmo já tendo atingido a quantidade
				    mínima de games para fechar o set. Pela regra do tênis, é necessário tem uma vantagem de,
				    pelo menos, dois games para fechar o set.
				*/

				if placar[0][0] == 45 && placar[0][1] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 1
					Caso ele esteja em vantagem no último ponto. Neste cenário, o número de sets é incrementado em 1 e
					os placares de games e pontos são zerados para o próximo set.
					*/
					placar[2][0]++ //Aumenta a contagem de sets em 1
					placar[1][0] = 0
					placar[1][1] = 0
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][0] == 45 && placar[0][1] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 2
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][0] += 5 //Aumenta a contagem de pontos em 5, simbolizando o ADVANTAGE

				} else if placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 3
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como estava a um game de fechar o set, também fecha o set. O número de sets é incrementado em 1 e
					os placares de games e pontos são zerados para o próximo set.
					*/
					placar[2][0]++ //Aumenta a contagem de sets em 1
					placar[1][0] = 0
					placar[1][1] = 0
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][0] == 45 && placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 4
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][1] -= 5 //Diminui a contagem de pontos do adversário em 5, pois perdeu o ADVANTAGE

				} else if placar[0][0] > 0 && placar[0][0] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][0] += 15 //Aumenta a contagem de pontos em 5
				}

			} else {
				/* CENÁRIO DE GAMES 3, 4 E 5
				Está cláusula "else" inclui os outros três cenários possíveis, descritos abaixo, que executam as mesmas ações

				CENÁRIO DE GAMES 3
				 Os PLAYERs estão empatados no número de games. Mesmo que vencer um game signifique que o PLAYER atingiu
				 o número mínimo de games necessários para fechar o set, pela regra do tênis, ele precisa de uma
				 vantagem de pelo menos dois games para fechar o set. Por isso, é apenas adicionada uma unidade a sua
				 contagem de sets.

				CENÁRIO DE GAMES 4
				 O PLAYER está em desvantagem no número de games, mesmo os dois já tendo atingido a quantidade mínima de
				 games. Pela mesma regra descrita acima, seu adversário necessitaria de dois games de vantagens para
				 fechar o set. Por isso, o set ainda não foi encerrado e uma unidade é adicionada a sua contagem de sets.

				CENÁRIO DE GAMES 5
				 Cenário normal, onde o PLAYER ainda não atingiu a quantidade necessária de games para fechar o set.
				*/

				if placar[0][0] == 45 && placar[0][1] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 5
					Caso ele esteja em vantagem no último ponto para fechar o game. Neste cenário, não era um match point,
					pois ganhando este game, o PLAYER não fecha o set. Portanto, foi apenas adicionada uma unidade a
					sua contagem de games. Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][0]++
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][0] == 45 && placar[0][1] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 6
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][0] += 5

				} else if placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 7
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como não estava a um game de fechar o set, apenas uma unidade é adicionada a sua contagem de games.
					Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][0]++
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][0] == 45 && placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 8
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][1] -= 5

				} else if placar[0][0] > 0 && placar[0][0] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][0] += 15
				}
			}
		}

	} else if player == 2 { //CENÁRIOS PARA O PLAYER 2

		if placar[2][1] == games-1 {
			/* CENÁRIO DE SETS 1
			Caso ele esteja a um set da vitória...
			*/

			if /*CENÁRIO 1*/ placar[1][1] == games-1 && placar[1][0] < games-1 || /*CENÁRIO 2*/ placar[1][1] >= games && placar[1][0] < placar[1][1] {
				/* CENÁRIO DE GAMES 1
					Caso ele esteja a um game de abrir uma diferença de dois games para seu adversário e atingir o mínimo
					de games para fechar o set.

				   CENÁRIO DE GAMES 2
					Caso ele esteja a apenas um game de vantagem do adversário, mesmo já tendo atingido a quantidade
				    mínima de games para fechar o set. Pela regra do tênis, é necessário tem uma vantagem de,
				    pelo menos, dois games para fechar o set.
				*/

				if placar[0][1] == 45 && placar[0][0] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 1
					Caso ele esteja em vantagem no último ponto. Neste cenário, era um match point. Como ele marcou,
					venceu o jogo.
					*/
					placar[2][1]++ //Aumenta a contagem de sets em 1
					fmt.Println("-----------------------------------------------------------------------------------------")
					fmt.Println(">>> PLAYER 2 venceu o jogo por " + string(placar[2][0]) + "x" + string(placar[2][1]) + ".")
					fmt.Println("-----------------------------------------------------------------------------------------")
					return true

				} else if placar[0][1] == 45 && placar[0][0] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 2
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][1] += 5 //Aumenta a contagem de pontos em 5, simbolizando o ADVANTAGE

				} else if placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 3
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como estava a um game de fechar o set, também fecha o set. E, como estava também a um set de ganhar
					o jogo, ele vence a partida.
					*/
					placar[2][1]++ //Aumenta a contagem de sets em 1
					fmt.Println("-----------------------------------------------------------------------------------------")
					fmt.Println(">>> PLAYER 2 venceu o jogo por " + string(placar[2][0]) + "x" + string(placar[2][1]) + ".")
					fmt.Println("-----------------------------------------------------------------------------------------")
					return true

				} else if placar[0][1] == 45 && placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 4
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][0] -= 5 //Diminui a contagem de pontos do adversário em 5, pois perdeu o ADVANTAGE

				} else if placar[0][1] > 0 && placar[0][1] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][1] += 15 //Aumenta a contagem de pontos em 5
				}

			} else {
				/* CENÁRIO DE GAMES 3, 4 E 5
				Está cláusula "else" inclui os outros três cenários possíveis, descritos abaixo, que executam as mesmas ações

				CENÁRIO DE GAMES 3
				 Os PLAYERs estão empatados no número de games. Mesmo que vencer um game signifique que o PLAYER atingiu
				 o número mínimo de games necessários para fechar o set, pela regra do tênis, ele precisa de uma
				 vantagem de pelo menos dois games para fechar o set. Por isso, é apenas adicionada uma unidade a sua
				 contagem de sets.

				CENÁRIO DE GAMES 4
				 O PLAYER está em desvantagem no número de games, mesmo os dois já tendo atingido a quantidade mínima de
				 games. Pela mesma regra descrita acima, seu adversário necessitaria de dois games de vantagens para
				 fechar o set. Por isso, o set ainda não foi encerrado e uma unidade é adicionada a sua contagem de sets.

				CENÁRIO DE GAMES 5
				 Cenário normal, onde o PLAYER ainda não atingiu a quantidade necessária de games para fechar o set.
				*/

				if placar[0][1] == 45 && placar[0][0] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 5
					Caso ele esteja em vantagem no último ponto para fechar o game. Neste cenário, não era um match point,
					pois ganhando este game, o PLAYER não fecha o set. Portanto, foi apenas adicionada uma unidade a
					sua contagem de games. Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][1]++
					placar[0][1] = 0
					placar[0][0] = 0

				} else if placar[0][1] == 45 && placar[0][0] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 6
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][1] += 5

				} else if placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 7
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como não estava a um game de fechar o set, apenas uma unidade é adicionada a sua contagem de games.
					Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][1]++
					placar[0][1] = 0
					placar[0][0] = 0

				} else if placar[0][1] == 45 && placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 8
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][0] -= 5

				} else if placar[0][1] > 0 && placar[0][1] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][1] += 15
				}
			}

		} else if placar[2][1] < games-1 {
			/* CENÁRIO DE SETS 2
			Caso ele não esteja a um set da vitória...
			*/

			if /*CENÁRIO 1*/ placar[1][1] == games-1 && placar[1][0] < games-1 || /*CENÁRIO 2*/ placar[1][1] >= games && placar[1][0] < placar[1][1] {
				/* CENÁRIO DE GAMES 1
					Caso ele esteja a um game de abrir uma diferença de dois games para seu adversário e atingir o mínimo
					de games para fechar o set.

				   CENÁRIO DE GAMES 2
					Caso ele esteja a apenas um game de vantagem do adversário, mesmo já tendo atingido a quantidade
				    mínima de games para fechar o set. Pela regra do tênis, é necessário tem uma vantagem de,
				    pelo menos, dois games para fechar o set.
				*/

				if placar[0][1] == 45 && placar[0][0] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 1
					Caso ele esteja em vantagem no último ponto. Neste cenário, o número de sets é incrementado em 1 e
					os placares de games e pontos são zerados para o próximo set.
					*/
					placar[2][1]++ //Aumenta a contagem de sets em 1
					placar[1][1] = 0
					placar[1][0] = 0
					placar[0][1] = 0
					placar[0][0] = 0

				} else if placar[0][1] == 45 && placar[0][0] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 2
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][1] += 5 //Aumenta a contagem de pontos em 5, simbolizando o ADVANTAGE

				} else if placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 3
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como estava a um game de fechar o set, também fecha o set. O número de sets é incrementado em 1 e
					os placares de games e pontos são zerados para o próximo set.
					*/
					placar[2][1]++ //Aumenta a contagem de sets em 1
					placar[1][1] = 0
					placar[1][0] = 0
					placar[0][1] = 0
					placar[0][0] = 0

				} else if placar[0][1] == 45 && placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 4
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][0] -= 5 //Diminui a contagem de pontos do adversário em 5, pois perdeu o ADVANTAGE

				} else if placar[0][1] > 0 && placar[0][1] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][1] += 15 //Aumenta a contagem de pontos em 5
				}

			} else {
				/* CENÁRIO DE GAMES 3, 4 E 5
				Está cláusula "else" inclui os outros três cenários possíveis, descritos abaixo, que executam as mesmas ações

				CENÁRIO DE GAMES 3
				 Os PLAYERs estão empatados no número de games. Mesmo que vencer um game signifique que o PLAYER atingiu
				 o número mínimo de games necessários para fechar o set, pela regra do tênis, ele precisa de uma
				 vantagem de pelo menos dois games para fechar o set. Por isso, é apenas adicionada uma unidade a sua
				 contagem de sets.

				CENÁRIO DE GAMES 4
				 O PLAYER está em desvantagem no número de games, mesmo os dois já tendo atingido a quantidade mínima de
				 games. Pela mesma regra descrita acima, seu adversário necessitaria de dois games de vantagens para
				 fechar o set. Por isso, o set ainda não foi encerrado e uma unidade é adicionada a sua contagem de sets.

				CENÁRIO DE GAMES 5
				 Cenário normal, onde o PLAYER ainda não atingiu a quantidade necessária de games para fechar o set.
				*/

				if placar[0][1] == 45 && placar[0][0] < 45 {
					/*CENÁRIO ESPECIAL DE PONTOS 5
					Caso ele esteja em vantagem no último ponto para fechar o game. Neste cenário, não era um match point,
					pois ganhando este game, o PLAYER não fecha o set. Portanto, foi apenas adicionada uma unidade a
					sua contagem de games. Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][1]++
					placar[0][0] = 0
					placar[0][1] = 0

				} else if placar[0][1] == 45 && placar[0][0] == 45 {
					/* CENÁRIO ESPECIAL DE PONTOS 6
					Caso ele esteja empatado com o adversário. Neste cenário, ganha 5 pontos que representam a regra
					ADVANTAGE do tênis
					*/
					placar[0][1] += 5

				} else if placar[0][1] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 7
					Caso ele tenha conquistado o ADVANTAGE e tenha marcado mais uma vez. Neste cenário, fecha o game.
					Como não estava a um game de fechar o set, apenas uma unidade é adicionada a sua contagem de games.
					Os placares de pontos são zerados para o próximo game.
					*/
					placar[1][1]++
					placar[0][1] = 0
					placar[0][0] = 0

				} else if placar[0][1] == 45 && placar[0][0] == 50 {
					/* CENÁRIO ESPECIAL DE PONTOS 8
					Caso o adversário tenha conquistado o ADVANTAGE. Neste cenário, como foi PLAYER que marcou,
					o adversário perde o ADVANTAGE e a disputa volta para o empate simples 45 a 45
					*/
					placar[0][0] -= 5

				} else if placar[0][1] > 0 && placar[0][1] < 45 {
					/* CENÁRIO NORMAL DE PONTOS
					Caso o PLAYER tenha menos de 45 pontos (0, 15 OU 30), ou seja, conquistou 15 pontos simples,
					sem cenário especial.
					*/
					placar[0][1] += 15
				}
			}
		}
	}

	return false
}

func zeraPlacar() {
	//Essa função inicializa o placar antes de uma partida

	/* placar[0][x] : Equivale ao número de pontos do PLAYER x
	   placar[1][x] : Equivale ao número de games do PLAYER x
	   placar[2][x] : Equivale ao número de sets do PLAYER x
	*/

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			placar[i][j] = 0
		}
	}
}

// Goroutines
func player1(quadra chan string) {
	action := <-quadra
	if action == "Devolveu" {
		fmt.Print(">>> PLAYER 1 está esperando a bola...")
		time.Sleep(time.Second)
		fails := rand.Intn(2)
		if fails == 0 {
			quadra <- "Devolveu"
			fmt.Print(" PLAYER 1 devolveu a bola para o adversário!")
		} else {
			quadra <- "Errou"
			fmt.Print(" PLAYER 1 errou a bola")
		}
	} else {
		fmt.Println("EITA")
		marcou(1)
	}
}

func player2(quadra chan string) {
	action := <-quadra
	if action == "Devolveu" {
		fmt.Print(">>> PLAYER 1 está esperando a bola...")
		time.Sleep(time.Second)
		fails := rand.Intn(2)
		if fails == 0 {
			quadra <- "Devolveu"
			fmt.Print(" PLAYER 1 devolveu a bola para o adversário!")
		} else {
			quadra <- "Errou"
			fmt.Print(" PLAYER 1 errou a bola")
		}
	} else {
		fmt.Println("EITA")
		marcou(2)
	}

}

// Função main
func main() {
	zeraPlacar()
	tela()

	quadra := make(chan string)

	go player1(quadra)
	go player2(quadra)

	time.Sleep(time.Second * 20)
}
