package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

const DB_PATH = "db.json" //Caminho para o arquivo

// Ler -> Modificar -> Sobrescrever

func leEescreveNoJson(path string, task string) (int, error) {
	var tasks []Task

	//Caso arquivo não exista, cria vazio
	if _, err := os.Stat(path); os.IsNotExist(err) {
		tasks = []Task{}
	} else {
		// Le os arquivos e coloca num slice array
		data, _ := os.ReadFile(path)
		json.Unmarshal(data, &tasks)
	}

	// adiciona nova tarefa
	newTask := Task{ID: len(tasks) + 1, Title: task, Status: "todo"}
	tasks = append(tasks, newTask)

	// Salva tarefa no 'banco'
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(path, data, 0644)

	return newTask.ID, nil
}

func add(title string) string {
	id, err := leEescreveNoJson(DB_PATH, title)

	if err != nil {
		fmt.Printf("Erro encontrado: %s", err)
	}

	return fmt.Sprintf("Tarefa adicionada com sucesso ID: %d", id)

}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Não tem argumento")
		return
	}
	// all args (excluding program name)
	args := os.Args[1:]

	if args[0] == "add" {
		if args[1] != "" {
			msg := add(args[1])
			fmt.Println(msg)
		}
	}
}
