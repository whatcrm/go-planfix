package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "go-planfix"
)

func testProcesses() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Processes API ===")

	// Тест 1: Получение процессов задач
	fmt.Println("\n1. Получение процессов задач")
	taskProcesses, err := client.GetTaskProcesses(ctx, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения процессов задач: %v", err)
	} else {
		fmt.Printf("   Процессы задач получены: %+v\n", taskProcesses)
	}

	// Тест 2: Получение процессов задач без полей
	fmt.Println("\n2. Получение процессов задач без полей")
	taskProcesses2, err := client.GetTaskProcesses(ctx)
	if err != nil {
		log.Printf("   Ошибка получения процессов задач: %v", err)
	} else {
		fmt.Printf("   Процессы задач получены: %+v\n", taskProcesses2)
	}

	// Тест 3: Получение статусов задач для процесса
	fmt.Println("\n3. Получение статусов задач для процесса")
	taskStatuses, err := client.GetTaskStatusesForProcess(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса: %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены: %+v\n", taskStatuses)
	}

	// Тест 4: Получение статусов задач для процесса без полей
	fmt.Println("\n4. Получение статусов задач для процесса без полей")
	taskStatuses2, err := client.GetTaskStatusesForProcess(ctx, 1)
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса: %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены: %+v\n", taskStatuses2)
	}

	// Тест 5: Получение процессов контактов
	fmt.Println("\n5. Получение процессов контактов")
	contactProcesses, err := client.GetContactProcesses(ctx, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения процессов контактов: %v", err)
	} else {
		fmt.Printf("   Процессы контактов получены: %+v\n", contactProcesses)
	}

	// Тест 6: Получение процессов контактов без полей
	fmt.Println("\n6. Получение процессов контактов без полей")
	contactProcesses2, err := client.GetContactProcesses(ctx)
	if err != nil {
		log.Printf("   Ошибка получения процессов контактов: %v", err)
	} else {
		fmt.Printf("   Процессы контактов получены: %+v\n", contactProcesses2)
	}

	// Тест 7: Получение статусов задач для несуществующего процесса
	fmt.Println("\n7. Получение статусов задач для несуществующего процесса")
	nonExistentStatuses, err := client.GetTaskStatusesForProcess(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены (неожиданно): %+v\n", nonExistentStatuses)
	}

	// Тест 8: Получение процессов задач с невалидными полями
	fmt.Println("\n8. Получение процессов задач с невалидными полями")
	taskProcesses3, err := client.GetTaskProcesses(ctx, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения процессов задач: %v", err)
	} else {
		fmt.Printf("   Процессы задач получены: %+v\n", taskProcesses3)
	}

	// Тест 9: Получение процессов контактов с невалидными полями
	fmt.Println("\n9. Получение процессов контактов с невалидными полями")
	contactProcesses3, err := client.GetContactProcesses(ctx, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения процессов контактов: %v", err)
	} else {
		fmt.Printf("   Процессы контактов получены: %+v\n", contactProcesses3)
	}

	// Тест 10: Получение статусов задач для процесса с невалидными полями
	fmt.Println("\n10. Получение статусов задач для процесса с невалидными полями")
	taskStatuses3, err := client.GetTaskStatusesForProcess(ctx, 1, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса: %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены: %+v\n", taskStatuses3)
	}

	// Тест 11: Получение процессов задач с пустыми полями
	fmt.Println("\n11. Получение процессов задач с пустыми полями")
	taskProcesses4, err := client.GetTaskProcesses(ctx, "")
	if err != nil {
		log.Printf("   Ошибка получения процессов задач: %v", err)
	} else {
		fmt.Printf("   Процессы задач получены: %+v\n", taskProcesses4)
	}

	// Тест 12: Получение процессов контактов с пустыми полями
	fmt.Println("\n12. Получение процессов контактов с пустыми полями")
	contactProcesses4, err := client.GetContactProcesses(ctx, "")
	if err != nil {
		log.Printf("   Ошибка получения процессов контактов: %v", err)
	} else {
		fmt.Printf("   Процессы контактов получены: %+v\n", contactProcesses4)
	}

	// Тест 13: Получение статусов задач для процесса с пустыми полями
	fmt.Println("\n13. Получение статусов задач для процесса с пустыми полями")
	taskStatuses4, err := client.GetTaskStatusesForProcess(ctx, 1, "")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса: %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены: %+v\n", taskStatuses4)
	}

	// Тест 14: Получение статусов задач для процесса с нулевым ID
	fmt.Println("\n14. Получение статусов задач для процесса с нулевым ID")
	zeroStatuses, err := client.GetTaskStatusesForProcess(ctx, 0, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены (неожиданно): %+v\n", zeroStatuses)
	}

	// Тест 15: Получение статусов задач для процесса с отрицательным ID
	fmt.Println("\n15. Получение статусов задач для процесса с отрицательным ID")
	negativeStatuses, err := client.GetTaskStatusesForProcess(ctx, -1, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены (неожиданно): %+v\n", negativeStatuses)
	}

	// Тест 16: Получение процессов задач с очень длинными полями
	fmt.Println("\n16. Получение процессов задач с очень длинными полями")
	longFields := "id,name,description,very_long_field_name_that_might_cause_issues,another_very_long_field_name"
	taskProcesses5, err := client.GetTaskProcesses(ctx, longFields)
	if err != nil {
		log.Printf("   Ошибка получения процессов задач: %v", err)
	} else {
		fmt.Printf("   Процессы задач получены: %+v\n", taskProcesses5)
	}

	// Тест 17: Получение процессов контактов с очень длинными полями
	fmt.Println("\n17. Получение процессов контактов с очень длинными полями")
	contactProcesses5, err := client.GetContactProcesses(ctx, longFields)
	if err != nil {
		log.Printf("   Ошибка получения процессов контактов: %v", err)
	} else {
		fmt.Printf("   Процессы контактов получены: %+v\n", contactProcesses5)
	}

	// Тест 18: Получение статусов задач для процесса с очень длинными полями
	fmt.Println("\n18. Получение статусов задач для процесса с очень длинными полями")
	taskStatuses5, err := client.GetTaskStatusesForProcess(ctx, 1, longFields)
	if err != nil {
		log.Printf("   Ошибка получения статусов задач для процесса: %v", err)
	} else {
		fmt.Printf("   Статусы задач для процесса получены: %+v\n", taskStatuses5)
	}

	// Тест 19: Получение процессов задач с полями, содержащими специальные символы
	fmt.Println("\n19. Получение процессов задач с полями, содержащими специальные символы")
	specialFields := "id,name,description,field-with-dashes,field_with_underscores,field.with.dots"
	taskProcesses6, err := client.GetTaskProcesses(ctx, specialFields)
	if err != nil {
		log.Printf("   Ошибка получения процессов задач: %v", err)
	} else {
		fmt.Printf("   Процессы задач получены: %+v\n", taskProcesses6)
	}

	// Тест 20: Получение процессов контактов с полями, содержащими специальные символы
	fmt.Println("\n20. Получение процессов контактов с полями, содержащими специальные символы")
	contactProcesses6, err := client.GetContactProcesses(ctx, specialFields)
	if err != nil {
		log.Printf("   Ошибка получения процессов контактов: %v", err)
	} else {
		fmt.Printf("   Процессы контактов получены: %+v\n", contactProcesses6)
	}

	fmt.Println("\n=== Тестирование Processes API завершено ===")
}

func main() {
	testProcesses()
}
