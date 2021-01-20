package editfiles

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ReadText(pathFile string) ([]string, error) {
	arquivo, err := os.Open(pathFile)
	if err != nil {
		return nil, err
	}
	defer arquivo.Close()
	var rows []string
	scanner := bufio.NewScanner(arquivo)
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}
	return rows, scanner.Err()
}

func writeText(conteudo []string, pathFile string) error {
	arquivo, err := os.Create(pathFile)
	if err != nil {
		return err
	}
	defer arquivo.Close()
	escritor := bufio.NewWriter(arquivo)
	for _, linha := range conteudo {
		fmt.Fprintln(escritor, linha)
	}
	return escritor.Flush()
}


// SetValueFile send path and data
func SetValueFile(path string, data []string) string{
	var file = path
	err := writeText(data, file)
	if err != nil {
		log.Fatalf("Erro:", err)
	}	
	var dadosArquivo, erro = ReadText(file)
	if erro != nil {
		log.Fatalf("Erro:", erro)
	}	
	fmt.Println(dadosArquivo)
	return("Updated")	
}