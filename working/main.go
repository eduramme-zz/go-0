package main

import (
	"fmt"
	"sort"
)

func main() {
	nomes, _ := os10maioresEstadosDoBrasil()

	for _, nome := range nomes {
		fmt.Println(nome)
	}
}

func os10maioresEstadosDoBrasil() ([]string, error) {

	var data []string

	type estado struct {
		nome     string
		extensao int
	}

	// Estados pegos da lista da wikipédia
	estados := []estado{
		{"Acre", 164123},
		{"Alagoas", 27778},
		{"Amapá", 142828},
		{"Amazonas", 1559159},
		{"Bahia", 564733},
		{"Ceará", 148920},
		{"Distrito Federal", 5779},
		{"Espirito Santo", 46095},
		{"Goiás", 340111},
		{"Maranhão", 331937},
		{"Mato Grosso do Sul", 357145},
		{"Minas Gerais", 586522},
		{"Pará", 1247954},
		{"Paraíba", 56585},
		{"Paraná", 199307},
		{"Pernambuco", 98311},
		{"Piauí", 251577},
		{"Rio de Janeiro", 43780},
		{"Rio Grande do Norte", 52811},
		{"Rio Grande do Sul", 281730},
		{"Rondônia", 237590},
		{"Roraima", 224300},
		{"Santa Catarina", 95736},
		{"São Paulo", 248222},
		{"Sergipe", 21915},
		{"Tocantins", 277720},
		{"Mato Grosso", 903366},
	}

	sort.Slice(estados[:], func(i, j int) bool {
		return estados[i].extensao > estados[j].extensao
	})

	for _, estado := range estados[:10] {
		data = append(data, estado.nome)
	}

	// fmt.Println(data)

	return data, nil
}
