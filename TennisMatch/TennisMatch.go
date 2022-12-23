package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var placar [3][2]int   //Placar do jogo
var games, sets = 1, 1 //Regras de pontuação do jogo. Podem ser alteradas pelo usuário (ver funções tela e regras)
var waitGroup sync.WaitGroup
var terminou bool = false

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
	fmt.Println("Número mínimo de games para vencer o set: ")
	fmt.Scanln(&games)
	fmt.Println("Número mínimo de sets para vencer a partida: ")
	fmt.Scanln(&sets)
}

func marcou(player int) {
	//Essa função simplesmente informa que o PLAYER marcou um ponto e atualiza e imprime o placar

	fmt.Println("-----------------------------------------------------------------------------------------")
	fmt.Printf(">>> PLAYER %d marcou ponto.\n", player)
	if player == 1 { //Se o vencedor do ponto for o PLAYER 1...
		atualizaPlacar(0, 1) //...o placar é atualizado com a informação de quem ganhou e de quem perdeu o ponto (winner, loser)
	} else { //Se o vencedor do ponto for o PLAYER 2...
		atualizaPlacar(1, 0) //...o placar é atualizado com a informação de quem ganhou e de quem perdeu o ponto (winner, loser)
	}
	fmt.Printf("  | PLAYER 1 | %d | %d | %d |\n", placar[0][0], placar[1][0], placar[2][0])
	fmt.Printf("  | PLAYER 2 | %d | %d | %d |\n", placar[0][1], placar[1][1], placar[2][1])
	fmt.Println("-----------------------------------------------------------------------------------------")

	if terminou == true { //Se a partida tiver terminado com este ponto conquistado...
		fmt.Printf(">>> PLAYER %d venceu o jogo.", player)
		fmt.Println("                                      FIM DE JOGO")
		fmt.Println("-----------------------------------------------------------------------------------------")
	} else {
		fmt.Println("                                      NOVA JOGADA")
		fmt.Println("-----------------------------------------------------------------------------------------")
	}
}

func atualizaPlacar(winner, loser int) {
	//Essa função atualiza a pontuação de uma player. Pode alterar uma variável de valor boolean que é positivo quando a partida é encerrada.
	//Ela descreve todos os cenários de pontuação possíveis numa partida de tênis para dois
	//jogadores.

	/* placar[0][x] : Equivale ao número de pontos do PLAYER x
	   placar[1][x] : Equivale ao número de games do PLAYER x
	   placar[2][x] : Equivale ao número de sets do PLAYER x
	*/

	if placar[0][winner] == 45 { //O PLAYER que marcou o ponto está prestes a ganhar o game...
		if placar[0][loser] == 45 { //...se o adversário também tem 45 pontos, configura um DEUCE (empate, na regra do tênis)...
			placar[0][winner] += 5 //...portanto, PLAYER recebe mais 5 pontos, simbolizando o conceito da regra ADVANTAGE
		} else if placar[0][loser] > 45 { //...se o adversário também tinha 45 pontos, mas já ganhou o ADVANTAGE, o ADVANTAGE é retirado dele
			placar[0][loser] -= 5 //...portanto, o adversário perde 5 pontos
		} else if placar[0][loser] < 45 { //...se o adversário tinha menos de 45 pontos...
			placar[0][winner] += 15 //...15 pontos são somados a pontuação do PLAYER e ele fecha este game
		}
	} else if placar[0][winner] == 50 { //...se o PLAYER possuía o ADVANTAGE e marcou mais um ponto, vencendo o game...
		placar[0][winner] += 10 //...apenas 10 pontos são adicionados a sua pontuação, para alcançar o limite máximo de 60 pontos
	} else { //Descreve todos os outros cenários de pontuação, que inclui as possibilidades de quando PLAYER não está prestes a ganhar o game
		placar[0][winner] += 15 //15 pontos são adicionados a sua pontuação
	}

	if placar[0][winner] == 60 { //Se o PLAYER possui 60 pontos, significa que ele ganhou o game, portanto...
		placar[1][winner]++ //...a contagem de games vencidos pelo PLAYER é incrementada em 1 unidade
		//Placares de pontuação são zerados para o próximo game
		placar[0][winner] = 0
		placar[0][loser] = 0
	}

	if placar[1][winner] >= games && placar[1][loser] < placar[1][winner] { //Se o PLAYER possui a quantidade de games necessária para vencer o set com uma diferença mínima de dois games em relação ao adversário (regra do tênis)...
		placar[2][winner]++ //...a contagem de sets vencidos pelo PLAYER é incrementada em 1 unidade
		//Placares de game são zerados para o próximo set
		placar[1][winner] = 0
		placar[1][loser] = 0
	}

	if placar[2][winner] == sets { //Se o PLAYER possui a quantidade de sets necessários para vencer a partida...
		terminou = true //...a variável que controla o estado da partida recebe valor true, significando que a partida terminou
	}
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
func player1(quadra chan string) { //PLAYER 1
	defer waitGroup.Done() //Goroutine sinalizará ao WaitGroup que terminou ao final da execução

	event := <-quadra //Lê o que está sendo transmitido pelo canal. No contexto, identifica qual ação foi executada anteriormente.

	if event == "Começa o jogo" { //Caso a ação anterior seja o início do jogo, ou seja, nenhum jogador devolveu a bola ainda...
		fmt.Println(">>> PLAYER 1 sacou a bola.") //...PLAYER 1 saca a bola...
		quadra <- "Devolveu"                      //...e a ação é informada no canal
	} else if event == "Devolveu" { //Caso a ação anterior seja uma devolução de PLAYER 2
		fmt.Print(">>> PLAYER 1 está esperando a bola... ") //PLAYER 1 se encontra esperando a bola que foi rebatida pelo PLAYER 2...
		time.Sleep(time.Millisecond * 500)                  //...aguarda a bola chegar...

		fails := rand.Intn(2) //...o programa decide randomicamente se o PLAYER 1 vai acertar a rebatida e devolver a bola ou se vai errar a tentativa.
		/*  fails == 0 : Significa que o PLAYER 1 não falhou, ou seja, devolveu a bola com sucesso.
		    fails == 1 : Significa que o PLAYER 1 falhou, ou seja, não devolveu a bola com sucesso.
		*/
		if fails == 0 { //Se PLAYER 1 conseguiu devolver a bola
			fmt.Println("PLAYER 1 devolveu a bola para PLAYER 2") //PLAYER 1 está devolvendo a bola...
			quadra <- "Devolveu"                                  //...e a ação é informada no canal
		} else { //Se PLAYER 1 não conseguiu devolver a bola
			fmt.Println("PLAYER 1 errou.") //PLAYER 1 errou...
			quadra <- "Errou"              //...e a ação é informada no canal
		}
	} else if event == "Errou" { //Caso a ação anterior seja um erro de PLAYER 2
		marcou(1)                 //O ponto é contabilizado para PLAYER 1...
		quadra <- "Começa o jogo" //...e o reinício do jogo com um novo saque é informado no canal
	}
}

func player2(quadra chan string) {
	defer waitGroup.Done() //Goroutine sinalizará ao WaitGroup que terminou ao final da execução

	event := <-quadra //Lê o que está sendo transmitido pelo canal. No contexto, identifica qual ação foi executada anteriormente.

	if event == "Começa o jogo" { //Caso a ação anterior seja o início do jogo, ou seja, nenhum jogador devolveu a bola ainda...
		fmt.Println(">>> PLAYER 2 sacou a bola...") //...PLAYER 2 saca a bola...
		quadra <- "Devolveu"                        //...e a ação é informada no canal
	} else if event == "Devolveu" { //Caso a ação anterior seja uma devolução de PLAYER 1
		fmt.Print(">>> PLAYER 2 está esperando a bola... ") //PLAYER 2 se encontra esperando a bola que foi rebatida pelo PLAYER 2...
		time.Sleep(time.Millisecond * 500)                  //...aguarda a bola chegar...

		fails := rand.Intn(2) //...o programa decide randomicamente se o PLAYER 2 vai acertar a rebatida e devolver a bola ou se vai errar a tentativa.
		/*  fails == 0 : Significa que o PLAYER 2 não falhou, ou seja, devolveu a bola com sucesso.
		    fails == 1 : Significa que o PLAYER 2 falhou, ou seja, não devolveu a bola com sucesso.
		*/
		if fails == 0 { //Se PLAYER 2 conseguiu devolver a bola
			fmt.Println("PLAYER 2 devolveu a bola para PLAYER 1") //PLAYER 2 está devolvendo a bola...
			quadra <- "Devolveu"                                  //...e a ação é informada no canal
		} else { //Se PLAYER 2 não conseguiu devolver a bola
			fmt.Println("PLAYER 2 errou.") //PLAYER 2 errou...
			quadra <- "Errou"              //...e a ação é informada no canal
		}
	} else if event == "Errou" { //Caso a ação anterior seja um erro de PLAYER 1
		marcou(2)                 //O ponto é contabilizado para PLAYER 2...
		quadra <- "Começa o jogo" //...e o reinício do jogo com um novo saque é informado no canal
	}
}

// Função main
func main() {
	zeraPlacar() //Inicializa o placar
	tela()       //Executa a "tela inicial"

	quadra := make(chan string, 1) //Cria o canal com buffer de tamanho 1
	quadra <- "Começa o jogo"      //Sinaliza à primeira goroutine que será executada que o jogo começou

	//Enquanto o jogo estiver acontecendo...
	for true {
		waitGroup.Add(2)      //Duas novas goroutines são adicionadas ao WaitGroup
		go player1(quadra)    //Goroutine do PLAYER 1
		go player2(quadra)    //Goroutine do PLAYER 2
		if terminou == true { //Se a partida tiver terminado...
			break //...o laço é quebrado
		}
	}

	waitGroup.Wait() //O WaitGroup garante que a goroutine main não terminará sua execução enquanto as outras goroutines não tiverem terminado as suas respectivas execuções e o WaitGroup estiver vazio
}
