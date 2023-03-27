package main

import (
	"fmt"
	"os"
	"os/signal"
)

/*
1. Создаем канал "signals" для обрабатывания сигналов.
2. Создаем канал "jobs" для передачи заданий воркерам.
3. Запускаем указанное количество воркеров.
4. Начинаем постоянно записывать данные в канал.
5. Если поступит сигнал остановки работы программы, то мы отправляем значение "-1" в каждый запущенный воркер.
6. Воркеры останавливают свою работу, если получают значение "-1".

В данном реализации мы используем `os.Interrupt`, чтобы выходить из программы при получении сигнала `Ctrl+C`.
При получении этого сигнала, мы выводим значение "-1" в каждый запущенный воркер, чтобы остановить их работу.
Воркер проверяет значение, переданное через канал "jobs", если он получил -1, то он заканчивает свою работу.

	В противном случае он принимает задание, которое нужно выполнить.

Код конструкции `default` используется, чтобы не блокировать выполнение программы в ожидании данных от канала.
Если нет доступных задач, то мы переходим к основному циклу выполнения программы.
*/

func main() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	workCount := 200
	jobs := make(chan int)

	// запускаем воркеров
	for i := 0; i < workCount; i++ {
		go worker(i, jobs)
	}

	// записываем данные в канал
	for i := 0; ; i++ {
		select {
		case <-signals:
			// останавливаем работу воркеров
			for j := 0; j < workCount; j++ {
				jobs <- -1
			}
			return
		default:
			jobs <- i
		}
	}
}

func worker(id int, jobs chan int) {
	for {
		job := <-jobs
		if job == -1 {

			return
		}
		fmt.Println("worker", id, "processing job", job)
	}
}
