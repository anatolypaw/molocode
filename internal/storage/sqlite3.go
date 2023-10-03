package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite3 struct {
	db *sql.DB
}

// Возвращает подключение к базе данных
// Инициализирует базу
func NewSqlite3(storagePath string) (*sqlite3, error) {
	const op = "storage.sqlite.New"

	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	//Таблица для хранения продуктов
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS goods(
		good_id INTEGER PRIMARY KEY AUTOINCREMENT,
		gtin TEXT NOT NULL UNIQUE,
		description TEXT);
	CREATE TABLE IF NOT EXISTS arm(
		arm_id INTEGER PRIMARY KEY AUTOINCREMENT,
		arm_name TEXT NOT NULL UNIQUE
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &sqlite3{db: db}, nil
}

// Закрыть подключение к базе данных
func (s *sqlite3) Close() error {
	const op = "storage.sqlite3.Close"
	err := s.db.Close()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

// Вывести имеющиеся продукты
func (s *sqlite3) GetGoods() ([]good, error) {
	const op = "storage.sqlite3.GetGoods"

	rows, err := s.db.Query(`select good_id, gtin, description from goods;`)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var goods []good

	for rows.Next() {
		var g good
		err := rows.Scan(&g.good_id, &g.gtin, &g.description)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		goods = append(goods, g)
	}

	return goods, nil
}

// Добавить продукт в хранилище
func (s *sqlite3) AddGood(gtin string, description string) error {
	const op = "storage.sqlite3.AddGood"

	stmt, err := s.db.Prepare(`
	INSERT INTO goods(gtin, description) VALUES (?, ?);
	`)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = stmt.Exec(gtin, description)
	if err != nil {
		return fmt.Errorf("%s: %w, gtin: %s, description: %s", op, err, gtin, description)
	}

	stmt.Close()
	return nil
}
