
# Go-Do CLI

A simple task management (To-Do List) application based on a Command Line Interface (CLI) built using the Go programming language. This project is part of learning Go fundamentals, covering slice manipulation, control flow, and interaction with the file system.

## Feature
- Add Task: adding tasks to the list dynamically
- Show Task: Displaying tasks with neat numbering
- Delete Task: Deleting tasks using serial numbers with index validation
- Save to File: Exporting the task list to the file "outputs.txt" for permanent storage

## Things I have learned
In the development of this application, I implemented several core Go concepts:
- **Slices**: Using `append` for adding data and *slicing* technique for delete element
- **For loops**: implemented *infinite loop* for menu and `for range` for data iteration 
- **Error Handlings**: Performing user input validation to prevent the program from *panicking*
- **File I/O**: Changing data from memory (RAM) to be physical file (`.txt`)

## Installation
1. Make sure you're already install [Go](https://go.dev/dl/)
2. Clone this repo:
```bash
