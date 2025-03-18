package parsingPrice

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// AllParsing парсит цену с повторными попытками
func AllParsing(url string) (float64, error) {
	var maxRetries = 3
	var price float64
	var err error
	for i := 1; i <= maxRetries; i++ {
		price, err = tryParsing(url)
		if err == nil {
			return price, nil // Успех — возвращаем цену
		}
		log.Printf("Попытка %d: ошибка - %v", i, err)

		if i < maxRetries {
			time.Sleep(time.Duration(2+rand.Intn(2)) * time.Second) // Задержка перед следующей попыткой
		}
	}
	return 0, fmt.Errorf("не удалось выполнить парсинг после %d попыток: %v", maxRetries, err)
}

func tryParsing(url string) (float64, error) {
	// Общий таймаут для всех операций (30 секунд)
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Случайная задержка
	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(2+rand.Intn(3)) * time.Second)

	// Настройка браузера
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true), // false для визуального слежения, дебага
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("disable-web-security", "1"), //отключение ключевых механизмов веб-безопасности, не использовать в продакшене
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/120.0.0.0 Safari/537.36"),
		chromedp.DisableGPU, // Добавляем для стабильности
	)

	allocCtx, allocCancel := chromedp.NewExecAllocator(ctx, opts...)
	defer allocCancel()

	browserCtx, browserCancel := chromedp.NewContext(allocCtx)
	defer browserCancel()

	// Извлекаем текст цены
	var priceText string
	err := chromedp.Run(browserCtx,
		chromedp.Navigate(url),
		chromedp.WaitReady(`div[data-widget="webPrice"] span`, chromedp.ByQuery),
		chromedp.Text(`div[data-widget="webPrice"] span:first-child`, &priceText),
	)
	if err != nil {
		return 0, fmt.Errorf("ошибка браузера: %v", err)
	}

	return convertPriceString(priceText)
}

func convertPriceString(priceText string) (float64, error) {

	// Заменяем все нестандартные пробелы и символы перед извлечением чисел
	priceText = strings.NewReplacer(
		"\u202f", "", // Удаляем неразрывные пробелы (U+202F)
		"\u2009", "", // Удаляем тонкие пробелы (U+2009)
		"\u00A0", "", // Удаляем неразрывный пробел (U+00A0)
		" ", "", // Удаляем обычные пробелы
		"₽", "", // Удаляем символ валюты
		",", ".", // Заменяем запятые на точки (если нужно)
	).Replace(priceText)

	// Удаляем все символы, кроме цифр, точек и запятых
	re := regexp.MustCompile(`\d+\.?\d*`)
	match := re.FindString(priceText)
	if match == "" {
		return 0, fmt.Errorf("не найдена цена: %s", priceText)
	}

	//Преобразуем во float64
	price, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, fmt.Errorf("ошибка конвертации: %v", err)
	}

	return price, nil
}