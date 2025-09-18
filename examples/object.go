package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "github.com/whatcrm/go-planfix"
	"github.com/whatcrm/go-planfix/models"
)

func testObjects() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Objects API ===")

	// Тест 1: Получение объекта по ID
	fmt.Println("\n1. Получение объекта по ID")
	object, err := client.GetObject(ctx, 1)
	if err != nil {
		log.Printf("   Ошибка получения объекта: %v", err)
	} else {
		fmt.Printf("   Объект получен: %+v\n", object)
	}

	// Тест 2: Получение объекта с несуществующим ID
	fmt.Println("\n2. Получение объекта с несуществующим ID")
	nonExistentObject, err := client.GetObject(ctx, 999999)
	if err != nil {
		log.Printf("   Ошибка получения объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Объект получен (неожиданно): %+v\n", nonExistentObject)
	}

	// Тест 3: Получение списка объектов
	fmt.Println("\n3. Получение списка объектов")
	objectListRequest := &models.ObjectListRequest{
		Offset:   0,
		PageSize: 10,
	}
	objects, err := client.GetObjects(ctx, objectListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects)
	}

	// Тест 4: Получение списка объектов без полей
	fmt.Println("\n4. Получение списка объектов без полей")
	objectListRequest2 := &models.ObjectListRequest{
		Offset:   0,
		PageSize: 5,
	}
	objects2, err := client.GetObjects(ctx, objectListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects2)
	}

	// Тест 5: Получение статусов задач объекта
	fmt.Println("\n5. Получение статусов задач объекта")
	objectStatuses, err := client.GetObjectTaskStatuses(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта: %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены: %+v\n", objectStatuses)
	}

	// Тест 6: Получение статусов задач объекта без полей
	fmt.Println("\n6. Получение статусов задач объекта без полей")
	objectStatuses2, err := client.GetObjectTaskStatuses(ctx, 1)
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта: %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены: %+v\n", objectStatuses2)
	}

	// Тест 7: Получение статусов задач несуществующего объекта
	fmt.Println("\n7. Получение статусов задач несуществующего объекта")
	nonExistentStatuses, err := client.GetObjectTaskStatuses(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены (неожиданно): %+v\n", nonExistentStatuses)
	}

	// Тест 8: Получение объекта с нулевым ID
	fmt.Println("\n8. Получение объекта с нулевым ID")
	zeroObject, err := client.GetObject(ctx, 0)
	if err != nil {
		log.Printf("   Ошибка получения объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Объект получен (неожиданно): %+v\n", zeroObject)
	}

	// Тест 9: Получение списка объектов с большим offset
	fmt.Println("\n9. Получение списка объектов с большим offset")
	objectListRequest3 := &models.ObjectListRequest{
		Offset:   1000,
		PageSize: 10,
	}
	objects3, err := client.GetObjects(ctx, objectListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects3)
	}

	// Тест 10: Получение списка объектов с нулевым pageSize
	fmt.Println("\n10. Получение списка объектов с нулевым pageSize")
	objectListRequest4 := &models.ObjectListRequest{
		Offset:   0,
		PageSize: 0,
	}
	objects4, err := client.GetObjects(ctx, objectListRequest4)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects4)
	}

	// Тест 11: Получение статусов задач объекта с невалидными полями
	fmt.Println("\n11. Получение статусов задач объекта с невалидными полями")
	objectStatuses3, err := client.GetObjectTaskStatuses(ctx, 1, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта: %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены: %+v\n", objectStatuses3)
	}

	// Тест 12: Получение объекта с отрицательным ID
	fmt.Println("\n12. Получение объекта с отрицательным ID")
	negativeObject, err := client.GetObject(ctx, -1)
	if err != nil {
		log.Printf("   Ошибка получения объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Объект получен (неожиданно): %+v\n", negativeObject)
	}

	// Тест 13: Получение статусов задач объекта с отрицательным ID
	fmt.Println("\n13. Получение статусов задач объекта с отрицательным ID")
	negativeStatuses, err := client.GetObjectTaskStatuses(ctx, -1, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены (неожиданно): %+v\n", negativeStatuses)
	}

	// Тест 14: Получение списка объектов с отрицательным offset
	fmt.Println("\n14. Получение списка объектов с отрицательным offset")
	objectListRequest5 := &models.ObjectListRequest{
		Offset:   -1,
		PageSize: 10,
	}
	objects5, err := client.GetObjects(ctx, objectListRequest5)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects5)
	}

	// Тест 15: Получение списка объектов с отрицательным pageSize
	fmt.Println("\n15. Получение списка объектов с отрицательным pageSize")
	objectListRequest6 := &models.ObjectListRequest{
		Offset:   0,
		PageSize: -1,
	}
	objects6, err := client.GetObjects(ctx, objectListRequest6)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects6)
	}

	// Тест 16: Получение объекта с очень большим ID
	fmt.Println("\n16. Получение объекта с очень большим ID")
	bigObject, err := client.GetObject(ctx, 999999999)
	if err != nil {
		log.Printf("   Ошибка получения объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Объект получен (неожиданно): %+v\n", bigObject)
	}

	// Тест 17: Получение статусов задач объекта с очень большим ID
	fmt.Println("\n17. Получение статусов задач объекта с очень большим ID")
	bigStatuses, err := client.GetObjectTaskStatuses(ctx, 999999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены (неожиданно): %+v\n", bigStatuses)
	}

	// Тест 18: Получение списка объектов с очень большим offset
	fmt.Println("\n18. Получение списка объектов с очень большим offset")
	objectListRequest7 := &models.ObjectListRequest{
		Offset:   999999999,
		PageSize: 10,
	}
	objects7, err := client.GetObjects(ctx, objectListRequest7)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects7)
	}

	// Тест 19: Получение списка объектов с очень большим pageSize
	fmt.Println("\n19. Получение списка объектов с очень большим pageSize")
	objectListRequest8 := &models.ObjectListRequest{
		Offset:   0,
		PageSize: 999999999,
	}
	objects8, err := client.GetObjects(ctx, objectListRequest8)
	if err != nil {
		log.Printf("   Ошибка получения списка объектов: %v", err)
	} else {
		fmt.Printf("   Список объектов получен: %+v\n", objects8)
	}

	// Тест 20: Получение статусов задач объекта с пустыми полями
	fmt.Println("\n20. Получение статусов задач объекта с пустыми полями")
	emptyStatuses, err := client.GetObjectTaskStatuses(ctx, 1, "")
	if err != nil {
		log.Printf("   Ошибка получения статусов задач объекта: %v", err)
	} else {
		fmt.Printf("   Статусы задач объекта получены: %+v\n", emptyStatuses)
	}

	fmt.Println("\n=== Тестирование Objects API завершено ===")
}

func main() {
	testObjects()
}
