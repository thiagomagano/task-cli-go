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

const DB_PATH = "db.json" //Caminho para o arquivo // Connection Pool

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

func getTasks() ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(DB_PATH)

	if err != nil {
		msg := fmt.Errorf("Ocorreu um erro ao ler tasks do db: %w", err)
		fmt.Println(msg)
		return []Task{}, err
	} else {
		json.Unmarshal(data, &tasks)
	}

	return tasks, nil
}

func list() {
	tasks, err := getTasks()

	if err != nil {
		return
	}

	msg := "======== Listando Tasks ========= \n\n"

	for _, task := range tasks {
		msg += fmt.Sprintf("- [ ] %v\n", task.Title)
	}
	msg += fmt.Sprintf("\n\nTotal de tarefas: %d\n", len(tasks))

	fmt.Println(msg)
}

func add(title string) {
	id, err := leEescreveNoJson(DB_PATH, title)

	if err != nil {
		fmt.Printf("Erro encontrado: %s", err)
		return
	}

	msg := fmt.Sprintf("Tarefa adicionada com sucesso ID: %d", id)
	fmt.Println(msg)
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Não tem argumento")
		return
	}
	// all args (excluding program name)
	args := os.Args[1:]
	//Primeiro argumento 'comando'
	command := args[0]

	switch command {
	case "add":
		if args[1] != "" {
			//Segundo argumento no comando add 'titulo' da task
			title := args[1]
			add(title)
		}
	case "list":
		list()
	}
}
