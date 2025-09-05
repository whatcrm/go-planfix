package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "go-planfix"
)

func testComments() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Comments API ===")

	// Тест 1: Получение комментария по ID
	fmt.Println("\n1. Получение комментария по ID")
	comment, err := client.GetComment(ctx, "123", "id,description,owner", "test-source")
	if err != nil {
		log.Printf("   Ошибка получения комментария: %v", err)
	} else {
		fmt.Printf("   Комментарий получен: %+v\n", comment)
	}

	// Тест 2: Получение комментария без дополнительных параметров
	fmt.Println("\n2. Получение комментария без параметров")
	comment2, err := client.GetComment(ctx, "456", "", "")
	if err != nil {
		log.Printf("   Ошибка получения комментария: %v", err)
	} else {
		fmt.Printf("   Комментарий получен: %+v\n", comment2)
	}

	// Тест 3: Удаление комментария
	fmt.Println("\n3. Удаление комментария")
	err = client.DeleteComment(ctx, "789")
	if err != nil {
		log.Printf("   Ошибка удаления комментария: %v", err)
	} else {
		fmt.Println("   Комментарий успешно удален")
	}

	// Тест 4: Удаление несуществующего комментария
	fmt.Println("\n4. Удаление несуществующего комментария")
	err = client.DeleteComment(ctx, "999999")
	if err != nil {
		log.Printf("   Ошибка удаления (ожидаемо): %v", err)
	} else {
		fmt.Println("   Комментарий удален (неожиданно)")
	}

	fmt.Println("\n=== Тестирование Comments API завершено ===")
}

func main() {
	testComments()
}
