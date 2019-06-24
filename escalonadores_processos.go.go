package main

import "fmt"

type process struct {
	name        int
	burst       int
	waitingTime int
	burstTotal  int
	chegada     int
}

var processo [20]process
var aguardando [20]int

var choice int
var j int
var k int
var i int

func menu() {
	fmt.Print("Escolha algum algoritmo de escalonamento a baixo: \n")
	fmt.Print("1 - FCFS \n")
	fmt.Print("2 - SJF \n")
	fmt.Print("3 - SRTF \n")
	fmt.Print("4 - Round Robin \n")
	fmt.Print("5 - sair \n")
	fmt.Print("Digite a escolha (1 à 5): ")
	fmt.Scanln(&choice)
}

func fcfs() {
	var bt [20]int
	var n int
	var max int
	var wt [20]int
	var tat [20]int
	var avwt = 0
	var avtat = 0

	fmt.Print("Entre com o total de numero de processos (no maximo 20): ")
	fmt.Scanln(&n)
	fmt.Print("Entre com o burst time do processo: \n")

	for i := 0; i < n; i++ {
		fmt.Printf("processo[%d] : ", i+1)
		fmt.Scanln(&bt[i])
	}

	wt[0] = 0

	for i := 1; i < n; i++ {
		wt[i] = 0
		for j := 0; j < i; j++ {
			wt[i] += bt[j]
		}
	}

	fmt.Print("\nProcess\t\tBurst Time\tWaiting Time\tTurnaround Time\n")

	for i := 0; i < n; i++ {
		tat[i] = bt[i] + wt[i]
		avwt += wt[i]
		avtat += tat[i]
		fmt.Print("\n", "   ", i+1, "\t", "\t", "   ", bt[i], "\t", "\t", "   ", wt[i], "\t", "\t", "   ", tat[i], "\n")
		max = i

	}

	avwt /= max
	avtat /= max
	fmt.Printf("\n média do tempo de espera: %d \n", avwt)
	fmt.Printf("\n média turnaround %d \n", avtat)
}

func RR() {
	var n int
	var max int
	var tat [20]int
	var avwt = 0
	var avtat = 0
	var quantum int

	fmt.Print("Entre com o numero de processos (maximo 20): \n") // entrada de numero de processos
	fmt.Scanln(&n)

	fmt.Print("Digite o Quantum de tempo da CPU: \n")
	fmt.Scanln(&quantum)

	for i := 0; i < n; i++ { // preenchimento de burst dos processos
		fmt.Print("Entre com o Burst Time do processo: \n")
		fmt.Printf("processo[%d] : ", i+1)
		processo[i].name = i + 1
		fmt.Scanln(&processo[i].burst) // Burst é gravado na struct do processo
		processo[i].waitingTime = 0    // waiting time de cada processo se inicia com 0
		processo[i].burstTotal = processo[i].burst
	}

	for i := 0; i < n; i++ { // enquanto tiver processo para ser preocessado, ele executara
		if processo[i].burst > 0 {
			executaProcessoRR(quantum, n)
			i = 0
		}
	}

	fmt.Print("\nProcess\t    Burst time\t    Waiting Time   Turnaround Time") // cabeçalho

	for i := 0; i < n; i++ {
		tat[i] = processo[i].burstTotal + processo[i].waitingTime // turn around time
		avwt += processo[i].waitingTime                           // average waiting time
		avtat += tat[i]                                           // average turn around time

		fmt.Print("\n", "  ", "p", processo[i].name, "\t", "\t", " ", processo[i].burstTotal, "\t", "\t", " ", processo[i].waitingTime, "\t", "\t", " ", tat[i], "\n")
		max = i
	}

	avwt /= max
	avtat /= max
	fmt.Printf("\n média do tempo de espera: %d \n", avwt)
	fmt.Printf("\n média turnaround: %d \n", avtat)
}

func executaProcessoRR(quantum int, n int) {
	for i := 0; i < n; i++ { //percorre processos
		if processo[i].burst > 0 {
			for j := 0; j != quantum; j++ { //processo sendo executado no tempo do quantum
				if processo[i].burst != 0 {
					processo[i].burst -= 1
					contaWaitingTimeRR(i, n)
				} else {
					break
				}

			}

		}
	}
}

func contaWaitingTimeRR(atual int, n int) {

	for i := 0; i < n; i++ { // percorre lista de processos
		if i != atual {
			if i != atual {
				if processo[i].burst > 0 {
					processo[i].waitingTime += 1 // conta o tempo de espera do processo
				}
			}

		}
	}
}

func sjf() {
	var n int
	var max int
	var bt [20]int
	var wt [20]int
	var tat [20]int
	var avwt = 0
	var avtat = 0
	// var aux = 0
	// var k = 0
	var procura_processo = 0
	var lista_bt_ord [20]int

	fmt.Print("Entre com o numero de processos (maximo 20): ") // entrada de numero de processos
	fmt.Scanln(&n)
	fmt.Print("Entre com o busrt time do processo: \n")

	for i := 0; i < n; i++ { // preenche o vetor de burst
		fmt.Printf("processo[%d] : ", i+1)
		fmt.Scanln(&bt[i]) // burst é gravado no array bt
	}

	lista_bt_ord[0] = bt[0]  // Insere o primeiro burst na lista ordenada
	for i := 1; i < n; i++ { // leitura do vertor burts
		for j := 0; j < i; j++ { // leitura de inserção na lista ordenada
			if bt[i] < lista_bt_ord[j] { // verifica se o burst for menor que algum que ja estiver na lista, ele entra no lugar
				for k := i; k > j; k-- { // organiza a lista
					lista_bt_ord[k] = lista_bt_ord[k-1]
				}
				lista_bt_ord[j] = bt[i]
				break
			} else if j == i {
				lista_bt_ord[j] = bt[i] // Insere no fim do vetor se for maior
				break
			}
		}
	}

	wt[0] = 0

	for i := 1; i < n; i++ {
		wt[i] = 0
		for j := 0; j < i; j++ {
			wt[i] += lista_bt_ord[j] // somatorio do tempo de waiting time
		}
	}

	fmt.Print("\nProcess\t\tBurst Time\tWaiting Time\tTurnaround Time\n")

	for i := 0; i < n; i++ {
		tat[i] = lista_bt_ord[i] + wt[i] // turnaroundtime = bursttime + waitingtime
		avwt += wt[i]                    // averagewaitingtime
		avtat += tat[i]                  // averagearoundtime
		procura_processo = lista_bt_ord[i]

		for j := 0; j < n; j++ { // verifica o processo que contem esse tempo de burst
			if procura_processo == bt[j] {
				procura_processo = j
				break
			}
		}
		fmt.Print("\n", "   ", procura_processo+1, "\t", "\t", "   ", lista_bt_ord[i], "\t", "\t", "   ", wt[i], "\t", "\t", "   ", tat[i], "\n") // apresenta tabela com os valores
		max = i
	}

	avwt /= max
	avtat /= max
	fmt.Printf("\n média do tempo de espera: %d \n", avwt)
	fmt.Printf("\n média turnaround: %d \n", avtat)
}

func mainSRTF() {
	var n int
	var max int
	var tat [20]int
	var avwt = 0
	var avtat = 0
	var pronto = 0
	var aux_pronto = 0

	fmt.Print("Entre com o total de numero de processos (no maximo 20): ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Print("Entre com o Burst Time do processo: \n")
		fmt.Printf("processo[%d] : ", i+1)
		processo[i].name = i + 1
		fmt.Scanln(&processo[i].burst)
		processo[i].burstTotal = processo[i].burst
		fmt.Print("Entre com o tempo de chegada do processo: \n")
		fmt.Scanln(&processo[i].chegada)
	}

	for i := 0; i < 100; i++ {
		aux_pronto = verificaEntradaSRFT(i, n)
		if aux_pronto > 0 {
			if pronto != 0 && processo[aux_pronto-1].burst < processo[pronto-1].burst {
				pronto = aux_pronto
				executaProcessoSRTF(pronto, i, n)
			} else if pronto != 0 && processo[aux_pronto-1].burst >= processo[pronto-1].burst {
				if processo[pronto-1].burst != 0 {
					executaProcessoSRTF(pronto, i, n)
				} else {
					pronto = selecionaNovoProcessoSRTF(n)
					if pronto == 0 {
						continue
					} else {
						executaProcessoSRTF(pronto, i, n)
					}
				}
			} else {
				pronto = aux_pronto
				executaProcessoSRTF(pronto, i, n)
			}
		} else if aux_pronto == 0 && pronto != 0 {
			if processo[pronto-1].burst != 0 {
				executaProcessoSRTF(pronto, i, n)
			} else {
				pronto = selecionaNovoProcessoSRTF(n)
				if pronto == 0 {
					break
				} else {
					executaProcessoSRTF(pronto, i, n)
				}
			}
		}
		if verificaFimSRTF(n) == 0 {
			break
		}
	}

	fmt.Print("\nProcess\t    Tempo de chegada\t     Burst time\t      Waiting Time     Turnaround Time") // cabeçalho

	for i := 0; i < n; i++ {
		tat[i] = processo[i].burstTotal + processo[i].waitingTime // turn around time
		avwt += processo[i].waitingTime                           // average waiting time
		avtat += tat[i]                                           // average turn around time

		fmt.Print("\n", "  ", "p", processo[i].name, "\t", "\t", "   ", processo[i].chegada, "\t", "\t", "           ", processo[i].burstTotal, "\t", "\t", "   ", processo[i].waitingTime, "\t", "\t", " ", tat[i], "\n")
		max = i
	}
	avwt /= max
	avtat /= max
	fmt.Printf("\n média do tempo de espera: %d \n", avwt)
	fmt.Printf("\n média turnaround: %d \n", avtat)
}

func verificaFimSRTF(n int) int {
	var cont = 0

	for i := 0; i < n; i++ {
		if processo[i].burst != 0 {
			cont++
		}
	}

	if cont > 0 {
		return 1
	} else {
		return 0
	}
}

func selecionaNovoProcessoSRTF(n int) int {
	var burst = 0
	var name = 0

	for i := 0; i < n; i++ {
		if processo[i].burst != 0 {
			name = processo[i].name
			burst = processo[i].burst

			for j := i; j < n; j++ {
				if processo[i].burst > 0 && processo[j].burst < burst {
					name = processo[j].name
					burst = processo[j].burst
				}
			}
			return name
		}
	}
	return burst
}

func verificaWaitingTimeSRTF(tempoAtual int, n int, id int) {

	for i := 0; i < n; i++ {
		if i != id-1 && processo[i].chegada <= tempoAtual && processo[i].burst != 0 {
			processo[i].waitingTime = processo[i].waitingTime + 1
		}
	}
}

func executaProcessoSRTF(id int, tempo int, n int) {
	processo[id-1].burst = processo[id-1].burst - 1
	verificaWaitingTimeSRTF(tempo, n, id)
}

func verificaEntradaSRFT(tempo int, n int) int {
	var pronto = 0

	for i := 0; i < n; i++ {
		if processo[i].chegada == tempo {
			pronto = processo[i].name
		}
	}
	return pronto
}

func main() {

	menu()

	switch choice {
	case 1:
		fmt.Print("*************************FCFS******************************\n")
		fcfs()
		fmt.Print("*************************************************************\n")
		menu()
		fcfs()
	case 2:
		fmt.Print("************************SJF******************************\n")
		sjf()
		fmt.Print("*************************************************************\n")
		menu()
		sjf()
	case 3:
		fmt.Print("************************SRTF******************************\n")
		mainSRTF()
		fmt.Print("*************************************************************\n")
		menu()
		mainSRTF()
	case 4:
		fmt.Print("*************************Round Robin******************************\n")
		RR()
		fmt.Print("*************************************************************\n")
		menu()
		RR()
	case 5:
		fmt.Print("flws\n")
		break
	}
}
