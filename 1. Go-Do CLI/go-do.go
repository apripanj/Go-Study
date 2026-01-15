package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var num int
	var tugas string
	var count int
	var tasks []string
	for {
		fmt.Println("====MENU CATATAN TUGAS====")
		fmt.Println("[1] Add Task")
		fmt.Println("[2] Show Task")
		fmt.Println("[3] Delete Task")
		fmt.Println("[4] Save to file")
		fmt.Print("Masukkan angka menu: ")
		fmt.Scanln(&num)

		switch num {
		case 1:
			fmt.Println("Mau input berapa tugas? ")
			fmt.Scanln(&count)
			for i := 0; i < count; i++ {
				fmt.Println("Masukkan tugas ke-", i+1)
				fmt.Scanln(&tugas)
				tasks = append(tasks, tugas)
			}
		case 2:
			fmt.Println("Tugas yang tersedia")
			for i, task := range tasks {
				fmt.Printf("%d %s\n", i+1, task)
			}

		case 3:
			var index int
			fmt.Printf("Masukkan index tugas yang mau dihapus: ")
			fmt.Scanln(&index)
			indeks := index - 1
			if indeks < 0 || indeks >= len(tasks) {
				fmt.Println("Index tidak valid")
			} else {
				tasks = append(tasks[:indeks], tasks[indeks+1:]...)
				fmt.Println("List tugas setelah dihapus: ", tasks)
			}
		case 4:
			content := strings.Join(tasks, "\n")
			data := []byte(content)
			os.WriteFile("outputs.txt", data, 0644)
		}
	}
}
