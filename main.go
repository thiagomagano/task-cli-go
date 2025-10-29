package main

import (
	"fmt"
	"os"
)

type Task struct {
	id     int
	title  string
	status string
}

//tasksList := make[]

func add(task string) string {
	// Ler arquivo
	// Json to Struct []hash
	// Escrever no arquivo -> task
	// retornar mensagem de sucesso
	newTask := new(Task)
	newTask.id = 1
	newTask.title = task
	newTask.status = "todo"

	return fmt.Sprintf("A tarefa \"%s\" foi adicionada com sucesso: (%d)", newTask.title, newTask.id)
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
			fmt.Printf(msg)
		}
	}
}
