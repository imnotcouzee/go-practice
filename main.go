package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ADD    = "Создать"
	LIST   = "Список"
	DELETE = "Удалить"
	CLEAR  = "Очистить"
	EXIT   = "Выйти"
	HELP   = "help"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var list []string

	for {

		fmt.Print("Введите команду: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Ошибка ввода")
			return
		}

		input := scanner.Text()
		sliceInput := strings.Fields(input)

		if len(sliceInput) == 0 {
			fmt.Println("Вы ничего не ввели")
			continue
		}

		cmd := sliceInput[0]

		if cmd != HELP && cmd != EXIT && cmd != ADD && cmd != LIST && cmd != DELETE && cmd != CLEAR {
			fmt.Println("Ошибка : Неизвестная команда")
			fmt.Println("Чтобы узнать спиоск команд напишите {help}")
		}

		if cmd == HELP {
			fmt.Println("Список команд:")
			fmt.Println("{", ADD, "}", " <--- команда которая позволяет добавить заметку")
			fmt.Println("{", DELETE, "}", " <--- команда которая позволяет удалить заметку по номеру")
			fmt.Println("{", CLEAR, "}", " <--- команда которая позволяет удалить все заметки")
			fmt.Println("{", LIST, "}", " <--- команда которая позволяет вывести список заметок")
			fmt.Println("{", EXIT, "}", " <--- команда которая позволяет выйти из программы")
		}

		if cmd == EXIT {
			fmt.Println("Всего доброго!")
			return
		}

		if cmd == ADD {
			if len(sliceInput) == 1 {
				fmt.Println("Введите текст заметки, которую нужно создать")
				continue
			}

			text := ""

			for i := 1; i < len(sliceInput); i++ {
				text += sliceInput[i]

				if i != len(sliceInput)-1 {
					text += " "
				}

			}

			list = append(list, text)
			fmt.Println("Создана заметка:", text)

		}

		if cmd == LIST {
			if len(list) == 0 {
				fmt.Println("Заметок для отображения нет")
				continue
			}

			for i := 1; i <= len(list); i++ {
				fmt.Println(i, ".", list[i-1])
			}
		}

		if cmd == DELETE {
			if len(list) == 0 {
				fmt.Println("Заметок для удаления нет")
				continue
			}

			getIntFromSlice, err := strconv.Atoi(sliceInput[1])

			if err != nil {
				fmt.Println("Значение не является числом")
				continue
			}

			deleteIndex := getIntFromSlice
			var findIndex bool
			var newSlice []string

			for i := 1; i <= len(list); i++ {

				if deleteIndex == i {
					findIndex = true
					newSlice = append(list[:deleteIndex-1], list[deleteIndex:]...)
					fmt.Println("Заметка успешно удалена")
				}

			}

			list = newSlice

			if !findIndex {
				fmt.Println("Заметка с таким номером не найдена")
				continue
			}
		}

		if cmd == CLEAR {
			if len(list) == 0 {
				fmt.Println("Заметок для удаления нет")
				continue
			}

			list = nil
			fmt.Println("Вы успешно удалили все заметки")
		}
	}
}
