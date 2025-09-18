package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "github.com/whatcrm/go-planfix"
	"github.com/whatcrm/go-planfix/models"
)

func testDataTags() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Data Tags API ===")

	// Тест 1: Получение data tag по ID
	fmt.Println("\n1. Получение data tag по ID")
	dataTag, err := client.GetDataTag(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения data tag: %v", err)
	} else {
		fmt.Printf("   Data tag получен: %+v\n", dataTag)
	}

	// Тест 2: Получение data tag без полей
	fmt.Println("\n2. Получение data tag без полей")
	dataTag2, err := client.GetDataTag(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения data tag: %v", err)
	} else {
		fmt.Printf("   Data tag получен: %+v\n", dataTag2)
	}

	// Тест 3: Получение списка data tags
	fmt.Println("\n3. Получение списка data tags")
	dataTagListRequest := &models.DataTagListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,description",
	}
	dataTags, err := client.GetDataTags(ctx, dataTagListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка data tags: %v", err)
	} else {
		fmt.Printf("   Список data tags получен: %+v\n", dataTags)
	}

	// Тест 4: Получение записи data tag по ключу
	fmt.Println("\n4. Получение записи data tag по ключу")
	dataTagEntry, err := client.GetDataTagEntry(ctx, 1, "id,key,value")
	if err != nil {
		log.Printf("   Ошибка получения записи data tag: %v", err)
	} else {
		fmt.Printf("   Запись data tag получена: %+v\n", dataTagEntry)
	}

	// Тест 5: Получение записи data tag без полей
	fmt.Println("\n5. Получение записи data tag без полей")
	dataTagEntry2, err := client.GetDataTagEntry(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения записи data tag: %v", err)
	} else {
		fmt.Printf("   Запись data tag получена: %+v\n", dataTagEntry2)
	}

	// Тест 6: Обновление записи data tag
	fmt.Println("\n6. Обновление записи data tag")
	updateRequest := &models.DataTagEntryUpdateRequest{
		CustomFieldData: []models.CustomFieldValueRequest{},
	}
	err = client.UpdateDataTagEntry(ctx, 1, updateRequest, false)
	if err != nil {
		log.Printf("   Ошибка обновления записи data tag: %v", err)
	} else {
		fmt.Println("   Запись data tag обновлена")
	}

	// Тест 7: Обновление записи data tag с silent режимом
	fmt.Println("\n7. Обновление записи data tag с silent режимом")
	updateRequest2 := &models.DataTagEntryUpdateRequest{
		CustomFieldData: []models.CustomFieldValueRequest{},
	}
	err = client.UpdateDataTagEntry(ctx, 2, updateRequest2, true)
	if err != nil {
		log.Printf("   Ошибка обновления записи data tag: %v", err)
	} else {
		fmt.Println("   Запись data tag обновлена (silent)")
	}

	// Тест 8: Удаление записи data tag
	fmt.Println("\n8. Удаление записи data tag")
	err = client.DeleteDataTagEntry(ctx, 3)
	if err != nil {
		log.Printf("   Ошибка удаления записи data tag: %v", err)
	} else {
		fmt.Println("   Запись data tag удалена")
	}

	// Тест 9: Удаление несуществующей записи data tag
	fmt.Println("\n9. Удаление несуществующей записи data tag")
	err = client.DeleteDataTagEntry(ctx, 999999)
	if err != nil {
		log.Printf("   Ошибка удаления (ожидаемо): %v", err)
	} else {
		fmt.Println("   Запись data tag удалена (неожиданно)")
	}

	// Тест 10: Получение списка записей data tag
	fmt.Println("\n10. Получение списка записей data tag")
	entryListRequest := &models.DataTagEntryListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,key,value",
	}
	dataTagEntries, err := client.GetListDataTagEntries(ctx, 1, entryListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка записей data tag: %v", err)
	} else {
		fmt.Printf("   Список записей data tag получен: %+v\n", dataTagEntries)
	}

	// Тест 11: Получение списка записей data tag без полей
	fmt.Println("\n11. Получение списка записей data tag без полей")
	entryListRequest2 := &models.DataTagEntryListRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	dataTagEntries2, err := client.GetListDataTagEntries(ctx, 2, entryListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка записей data tag: %v", err)
	} else {
		fmt.Printf("   Список записей data tag получен: %+v\n", dataTagEntries2)
	}

	// Тест 12: Получение data tag с несуществующим ID
	fmt.Println("\n12. Получение data tag с несуществующим ID")
	nonExistentDataTag, err := client.GetDataTag(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения data tag (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Data tag получен (неожиданно): %+v\n", nonExistentDataTag)
	}

	// Тест 13: Получение записи data tag с несуществующим ключом
	fmt.Println("\n13. Получение записи data tag с несуществующим ключом")
	nonExistentEntry, err := client.GetDataTagEntry(ctx, 999999, "id,key")
	if err != nil {
		log.Printf("   Ошибка получения записи data tag (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Запись data tag получена (неожиданно): %+v\n", nonExistentEntry)
	}

	// Тест 14: Обновление записи data tag с несуществующим ключом
	fmt.Println("\n14. Обновление записи data tag с несуществующим ключом")
	updateRequest3 := &models.DataTagEntryUpdateRequest{
		CustomFieldData: []models.CustomFieldValueRequest{},
	}
	err = client.UpdateDataTagEntry(ctx, 999999, updateRequest3, false)
	if err != nil {
		log.Printf("   Ошибка обновления записи data tag (ожидаемо): %v", err)
	} else {
		fmt.Println("   Запись data tag обновлена (неожиданно)")
	}

	// Тест 15: Получение списка записей data tag с несуществующим ID
	fmt.Println("\n15. Получение списка записей data tag с несуществующим ID")
	entryListRequest3 := &models.DataTagEntryListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,key,value",
	}
	nonExistentEntries, err := client.GetListDataTagEntries(ctx, 999999, entryListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка записей data tag (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Список записей data tag получен (неожиданно): %+v\n", nonExistentEntries)
	}

	fmt.Println("\n=== Тестирование Data Tags API завершено ===")
}

func main() {
	testDataTags()
}
