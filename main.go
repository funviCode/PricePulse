package main

import (
	"PricePulse/connect"
	"PricePulse/parsingPrice"
	"PricePulse/products"
	"PricePulse/telegramBot"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type ParseResult struct {
	Product products.Product
	Price   float64
	Error   error
}

func main() {
	// Настройка логирования в файл app.log
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Не удалось открыть файл для логирования: %v", err)
	}
	defer func() { //обрабатываем ошибку defer logFile.Close()
		if err := logFile.Close(); err != nil {
			log.Printf("Ошибка при закрытии файла логирования: %v", err)
		}
	}()
	log.SetOutput(logFile)

	// Загружаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка загрузки .env файла: ", err)
	}

	// Инициализация бота
	chatID, bot, err := telegramBot.InitBot()
	if err != nil {
		log.Fatal("Ошибка при инициализации бота: ", err)
	}

	// Открываем подключение к базе данных (один раз для всех продуктов)
	db, err := connect.ConnectDB()
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	//Закрываем подключение к базе данных
	defer func() { //обрабатываем ошибку defer db.Close()
		if err := db.Close(); err != nil {
			log.Printf("Ошибка при закрытии базы данных: %v", err)
		}
	}()

	// Канал для получения результатов
	results := make(chan ParseResult, len(products.Products))
	var wg sync.WaitGroup

	// Запуск горутин для каждого продукта
	for _, product := range products.Products {
		wg.Add(1)
		go func(p products.Product) {
			defer wg.Done()
			price, err := parsingPrice.AllParsing(p.Url)
			results <- ParseResult{Product: p, Price: price, Error: err}
		}(product)
	}

	// Закрываем канал, когда все горутины завершатся
	go func() {
		wg.Wait()
		close(results)
	}()

	// Обработка результатов
	for result := range results {
		if result.Error != nil {
			log.Printf("Ошибка для %s: %v\n", result.Product.Name, result.Error)
			continue
		}

		// Обновляем данные в базе данных
		if err := updateProductAndPrice(db, result.Product, result.Price); err != nil {
			log.Printf("Ошибка обновления данных для %s: %v", result.Product.Name, err)
			continue
		}

		log.Printf("Текущая цена %s: %.2f", result.Product.Name, result.Price)

		if result.Price < result.Product.PriceThreshold {
			msg := fmt.Sprintf("🚨 Цена упала до %.2f в продукте %s!\n%s", result.Price, result.Product.Name, result.Product.Url)
			if _, err := bot.Send(chatID, msg); err != nil {
				log.Printf("Ошибка отправки уведомления для %s: %v", result.Product.Name, err)
			} else {
				log.Printf("Уведомление отправлено для %s: %s", result.Product.Name, msg)
			}
		}
	}
	log.Println()
}

func updateProductAndPrice(db *sql.DB, product products.Product, price float64) error {
	// 1. Вставляем или обновляем запись в таблице products
	var productID int
	err := db.QueryRow(`
		INSERT INTO products (product_name, product_url) 
		VALUES ($1, $2)
		ON CONFLICT (product_url) DO UPDATE 
		SET product_name = EXCLUDED.product_name 
		RETURNING product_id
	`, product.Name, product.Url).Scan(&productID)
	if err != nil {
		return fmt.Errorf("ошибка при работе с products для %s: %w", product.Name, err)
	}

	// 2. Вставляем или обновляем запись в таблице prices
	_, err = db.Exec(`
		INSERT INTO prices (product_id, price) 
		VALUES ($1, $2)
		ON CONFLICT (product_id, date) DO UPDATE 
		SET price = EXCLUDED.price
	`, productID, price)
	if err != nil {
		return fmt.Errorf("ошибка при работе с prices для %s: %w", product.Name, err)
	}

	return nil
}
