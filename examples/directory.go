package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "github.com/whatcrm/go-planfix"
	"github.com/whatcrm/go-planfix/models"
)

func testDirectories() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Directories API ===")

	// Тест 1: Получение директории по ID
	fmt.Println("\n1. Получение директории по ID")
	directory, err := client.GetDirectoryByID(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения директории: %v", err)
	} else {
		fmt.Printf("   Директория получена: %+v\n", directory)
	}

	// Тест 2: Получение директории без полей
	fmt.Println("\n2. Получение директории без полей")
	directory2, err := client.GetDirectoryByID(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения директории: %v", err)
	} else {
		fmt.Printf("   Директория получена: %+v\n", directory2)
	}

	// Тест 3: Получение списка директорий
	fmt.Println("\n3. Получение списка директорий")
	directoryRequest := &models.DirectoryRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,description",
	}
	directories, err := client.GetListDirectories(ctx, directoryRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка директорий: %v", err)
	} else {
		fmt.Printf("   Список директорий получен: %+v\n", directories)
	}

	// Тест 4: Добавление записи в директорию
	fmt.Println("\n4. Добавление записи в директорию")
	entryRequest := &models.DirectoryEntryRequest{
		Value: "Тестовая запись",
	}
	createdEntry, err := client.AddDirectoryEntry(ctx, 1, entryRequest)
	if err != nil {
		log.Printf("   Ошибка добавления записи в директорию: %v", err)
	} else {
		fmt.Printf("   Запись в директорию добавлена: %+v\n", createdEntry)
	}

	// Тест 5: Получение записи директории по ключу
	fmt.Println("\n5. Получение записи директории по ключу")
	directoryEntry, err := client.GetDirectoryEntry(ctx, 1, 1, "id,key,name,description")
	if err != nil {
		log.Printf("   Ошибка получения записи директории: %v", err)
	} else {
		fmt.Printf("   Запись директории получена: %+v\n", directoryEntry)
	}

	// Тест 6: Получение записи директории без полей
	fmt.Println("\n6. Получение записи директории без полей")
	directoryEntry2, err := client.GetDirectoryEntry(ctx, 1, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения записи директории: %v", err)
	} else {
		fmt.Printf("   Запись директории получена: %+v\n", directoryEntry2)
	}

	// Тест 7: Обновление записи директории
	fmt.Println("\n7. Обновление записи директории")
	updateRequest := &models.DirectoryEntryRequest{
		Value: "Обновленная запись",
	}
	err = client.UpdateDirectoryEntry(ctx, 1, 1, updateRequest)
	if err != nil {
		log.Printf("   Ошибка обновления записи директории: %v", err)
	} else {
		fmt.Println("   Запись директории обновлена")
	}

	// Тест 8: Удаление записи директории
	fmt.Println("\n8. Удаление записи директории")
	err = client.DeleteDirectoryEntry(ctx, 1, 3)
	if err != nil {
		log.Printf("   Ошибка удаления записи директории: %v", err)
	} else {
		fmt.Println("   Запись директории удалена")
	}

	// Тест 9: Удаление несуществующей записи директории
	fmt.Println("\n9. Удаление несуществующей записи директории")
	err = client.DeleteDirectoryEntry(ctx, 1, 999999)
	if err != nil {
		log.Printf("   Ошибка удаления записи директории (ожидаемо): %v", err)
	} else {
		fmt.Println("   Запись директории удалена (неожиданно)")
	}

	// Тест 10: Получение списка записей директории
	fmt.Println("\n10. Получение списка записей директории")
	entryListRequest := &models.DirectoryRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,key,name,description",
	}
	directoryEntries, err := client.GetListDirectoryEntries(ctx, 1, entryListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка записей директории: %v", err)
	} else {
		fmt.Printf("   Список записей директории получен: %+v\n", directoryEntries)
	}

	// Тест 11: Получение списка записей директории без полей
	fmt.Println("\n11. Получение списка записей директории без полей")
	entryListRequest2 := &models.DirectoryRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	directoryEntries2, err := client.GetListDirectoryEntries(ctx, 2, entryListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка записей директории: %v", err)
	} else {
		fmt.Printf("   Список записей директории получен: %+v\n", directoryEntries2)
	}

	// Тест 12: Получение групп директорий
	fmt.Println("\n12. Получение групп директорий")
	directoryGroups, err := client.GetListDirectoryGroups(ctx, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения групп директорий: %v", err)
	} else {
		fmt.Printf("   Группы директорий получены: %+v\n", directoryGroups)
	}

	// Тест 13: Получение групп директорий без полей
	fmt.Println("\n13. Получение групп директорий без полей")
	directoryGroups2, err := client.GetListDirectoryGroups(ctx)
	if err != nil {
		log.Printf("   Ошибка получения групп директорий: %v", err)
	} else {
		fmt.Printf("   Группы директорий получены: %+v\n", directoryGroups2)
	}

	// Тест 14: Получение фильтров записей директории
	fmt.Println("\n14. Получение фильтров записей директории")
	directoryFilters, err := client.GetListDirectoryEntriesFilters(ctx, "1")
	if err != nil {
		log.Printf("   Ошибка получения фильтров директории: %v", err)
	} else {
		fmt.Printf("   Фильтры директории получены: %+v\n", directoryFilters)
	}

	// Тест 15: Получение директории с несуществующим ID
	fmt.Println("\n15. Получение директории с несуществующим ID")
	nonExistentDirectory, err := client.GetDirectoryByID(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения директории (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Директория получена (неожиданно): %+v\n", nonExistentDirectory)
	}

	// Тест 16: Получение записи директории с несуществующим ключом
	fmt.Println("\n16. Получение записи директории с несуществующим ключом")
	nonExistentEntry, err := client.GetDirectoryEntry(ctx, 1, 999999, "id,key")
	if err != nil {
		log.Printf("   Ошибка получения записи директории (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Запись директории получена (неожиданно): %+v\n", nonExistentEntry)
	}

	// Тест 17: Обновление записи директории с несуществующим ключом
	fmt.Println("\n17. Обновление записи директории с несуществующим ключом")
	updateRequest2 := &models.DirectoryEntryRequest{
		Value: "Попытка обновления несуществующей записи",
	}
	err = client.UpdateDirectoryEntry(ctx, 1, 999999, updateRequest2)
	if err != nil {
		log.Printf("   Ошибка обновления записи директории (ожидаемо): %v", err)
	} else {
		fmt.Println("   Запись директории обновлена (неожиданно)")
	}

	// Тест 18: Получение списка записей директории с несуществующим ID
	fmt.Println("\n18. Получение списка записей директории с несуществующим ID")
	entryListRequest3 := &models.DirectoryRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,key,name",
	}
	nonExistentEntries, err := client.GetListDirectoryEntries(ctx, 999999, entryListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка записей директории (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Список записей директории получен (неожиданно): %+v\n", nonExistentEntries)
	}

	// Тест 19: Получение фильтров записей директории с несуществующим ID
	fmt.Println("\n19. Получение фильтров записей директории с несуществующим ID")
	nonExistentFilters, err := client.GetListDirectoryEntriesFilters(ctx, "999999")
	if err != nil {
		log.Printf("   Ошибка получения фильтров директории (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Фильтры директории получены (неожиданно): %+v\n", nonExistentFilters)
	}

	// Тест 20: Добавление записи в несуществующую директорию
	fmt.Println("\n20. Добавление записи в несуществующую директорию")
	entryRequest2 := &models.DirectoryEntryRequest{
		Value: "Запись в несуществующую директорию",
	}
	_, err = client.AddDirectoryEntry(ctx, 999999, entryRequest2)
	if err != nil {
		log.Printf("   Ошибка добавления записи в директорию (ожидаемо): %v", err)
	} else {
		fmt.Println("   Запись в директорию добавлена (неожиданно)")
	}

	fmt.Println("\n=== Тестирование Directories API завершено ===")
}

func main() {
	testDirectories()
}
