/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	// "golang-learning/todo-go/cmd"

	"fmt"
	task "golang-learning/todo-go/todo"
)

func main() {
	// cmd.Execute()

	t := task.Task{Name: "test", Description: "ok"}
	t = *task.SetDone(&t)

	fmt.Println(t)

}
