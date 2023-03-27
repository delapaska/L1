package main

import "fmt"

// Интерфейс базы данных
type Database interface {
	Connect()
	Select(query string)
	Disconnect()
}

// Реализация интерфейса базы данных
type MySQL struct{}

func (m *MySQL) Connect() {
	fmt.Println("Connect to MySQL")
}

func (m *MySQL) Select(query string) {
	fmt.Printf("Performing query %s in MySQL\n", query)
}

func (m *MySQL) Disconnect() {
	fmt.Println("Disconnect from MySQL")
}

// Новый интерфейс базы данных
type NewDatabase interface {
	Search(query string)
}

// Адаптер для преобразования интерфейса базы данных в новый интерфейс
type DatabaseAdapter struct {
	Database
}

func (d *DatabaseAdapter) Search(query string) {
	d.Connect()
	d.Select(query)
	d.Disconnect()
}

// Функция, которая использует новый интерфейс базы данных
func runSearch(db NewDatabase, query string) {
	db.Search(query)
}

func main() {
	// Создаем объект базы данных
	mysql := &MySQL{}

	// Преобразуем интерфейс базы данных в новый интерфейс с помощью адаптера
	newDB := &DatabaseAdapter{mysql}

	// Выполняем поиск
	runSearch(newDB, "SELECT * FROM users")
}
