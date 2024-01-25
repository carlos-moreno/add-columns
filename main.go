package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		log.Println("Chamada do comando incorreta!")
		log.Println("Padrão da chamada deve ser => ./xpto `nome_do_arquivo` `'lista_tamanh_coluna'`")
		log.Println("Exemplo de chamada: ./xpto ARQUIVO_XPTO.txt '3, 3, 4, 1, 3'")
		os.Exit(1)
	}

	content, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	strContent := string(content)
	listStr := strings.Split(os.Args[2], ",")
	var list []int

	for _, str := range listStr {
		length, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, length)
	}

	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(strContent))

	for scanner.Scan() {
		line := scanner.Text()
		start := 0
		var result []string
		for _, length := range list {
			result = append(result, line[start:start+length])
			start += length
		}

		lines = append(lines, strings.Join(result, ";"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	output := strings.Join(lines, "\n")
	err = os.WriteFile(os.Args[1], []byte(output), 0644)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Arquivo alterado com sucesso!")
}
