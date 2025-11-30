package main

import (
    "fmt"
    "os"
    "strings"
)

type todo struct{
    WhoNeed string
    WhatDo string
}

func create_or_open_file(FileName string) *os.File {
    f, err := os.Open(FileName)
    if err != nil {
        if os.IsNotExist(err) {
            f, err = os.Create(FileName)
            if err != nil {
                panic(err)
            }
        } else {
            panic(err)
        }
    }
    return f
}

func write_file(file *os.File, data string) {
    byte_data := []byte(data + "\n")
    _, err := file.Write(byte_data)
    if err != nil {
        panic(err)
    }
}

func read_file(FileName string) string {
    content, err := os.ReadFile(FileName)
    if err != nil {
        panic(err)
    }
    return string(content)
}

func string_to_todo(str_todo string) todo {
    slice_todo := strings.Fields(str_todo)
    return todo{
        WhoNeed: slice_todo[0],
        WhatDo:  slice_todo[1],
    }
}

func main() {
    file_name := "todo_list.todo"
    todo_file := create_or_open_file(file_name)
    defer todo_file.Close()

    var todo_mode string
    fmt.Println("Please type the todo mode(READ, READ_ALL, WRITE):")
    fmt.Scanln(&todo_mode)

    switch todo_mode {
    case "READ_ALL":
        file_content := read_file(file_name)
        if file_content == "" {
            fmt.Println("There's nothing in todo_list.todo")
        } else {
            fmt.Println(strings.Split(file_content, "\n"))
        }
    case "READ":
        file_content := read_file(file_name)
        if file_content == "" {
            fmt.Println("There's nothing in todo_list.todo")
        } else {
            slice_todo_list := strings.Split(file_content, "\n")
            fmt.Println("Sure, please type the id:")
            var reading_id int
            fmt.Scanln(&reading_id)
            if reading_id >= 0 && reading_id < len(slice_todo_list) {
                fmt.Printf("Here's the result: %v\n", string_to_todo(slice_todo_list[reading_id]))
            } else {
                fmt.Println("There isn't this id")
            }
        }
    case "WRITE":
        var who_need string
        var what_do string
        fmt.Println("Please type your name and what do you need:")
        fmt.Scanln(&who_need, &what_do)
        write_file(todo_file, who_need+" "+what_do)
    default:
        fmt.Println("Invalid mode")
    }
}
 
