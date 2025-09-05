package main

import (
	"context"
	"fmt"
	"log"
	"os"

	goplanfix "go-planfix"
	"go-planfix/models"
)

func testFiles() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Files API ===")

	// Тест 1: Получение файла по ID
	fmt.Println("\n1. Получение файла по ID")
	file, err := client.GetFileByID(ctx, 1, "id,name,size,downloadUrl")
	if err != nil {
		log.Printf("   Ошибка получения файла: %v", err)
	} else {
		fmt.Printf("   Файл получен: %+v\n", file)
	}

	// Тест 2: Получение файла без полей
	fmt.Println("\n2. Получение файла без полей")
	file2, err := client.GetFileByID(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения файла: %v", err)
	} else {
		fmt.Printf("   Файл получен: %+v\n", file2)
	}

	// Тест 3: Обновление файла
	fmt.Println("\n3. Обновление файла")
	updateRequest := &models.FileUpdateRequest{
		Name: "Обновленное имя файла",
	}
	updateResult, err := client.UpdateFile(ctx, 1, updateRequest)
	if err != nil {
		log.Printf("   Ошибка обновления файла: %v", err)
	} else {
		fmt.Printf("   Файл обновлен: %+v\n", updateResult)
	}

	// Тест 4: Удаление файла
	fmt.Println("\n4. Удаление файла")
	err = client.DeleteFile(ctx, 3)
	if err != nil {
		log.Printf("   Ошибка удаления файла: %v", err)
	} else {
		fmt.Println("   Файл удален")
	}

	// Тест 5: Удаление несуществующего файла
	fmt.Println("\n5. Удаление несуществующего файла")
	err = client.DeleteFile(ctx, 999999)
	if err != nil {
		log.Printf("   Ошибка удаления файла (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл удален (неожиданно)")
	}

	// Тест 6: Скачивание файла
	fmt.Println("\n6. Скачивание файла")
	fileReader, err := client.DownloadFile(ctx, 1)
	if err != nil {
		log.Printf("   Ошибка скачивания файла: %v", err)
	} else {
		fmt.Println("   Файл скачан успешно")
		if fileReader != nil {
			fileReader.Close()
		}
	}

	// Тест 7: Скачивание несуществующего файла
	fmt.Println("\n7. Скачивание несуществующего файла")
	fileReader2, err := client.DownloadFile(ctx, 999999)
	if err != nil {
		log.Printf("   Ошибка скачивания файла (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл скачан (неожиданно)")
		if fileReader2 != nil {
			fileReader2.Close()
		}
	}

	// Тест 8: Прикрепление файла к задаче
	fmt.Println("\n8. Прикрепление файла к задаче")
	err = client.AttachFileToTask(ctx, 1, 1)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к задаче: %v", err)
	} else {
		fmt.Println("   Файл прикреплен к задаче")
	}

	// Тест 9: Прикрепление файла к контакту
	fmt.Println("\n9. Прикрепление файла к контакту")
	err = client.AttachFileToContact(ctx, 1, 1)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к контакту: %v", err)
	} else {
		fmt.Println("   Файл прикреплен к контакту")
	}

	// Тест 10: Прикрепление файла к проекту
	fmt.Println("\n10. Прикрепление файла к проекту")
	err = client.AttachFileToProject(ctx, 1, 1)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к проекту: %v", err)
	} else {
		fmt.Println("   Файл прикреплен к проекту")
	}

	// Тест 11: Загрузка файла по пути
	fmt.Println("\n11. Загрузка файла по пути")
	// Создаем временный файл для тестирования
	tempFile, err := os.CreateTemp("", "test_file_*.txt")
	if err != nil {
		log.Printf("   Ошибка создания временного файла: %v", err)
	} else {
		tempFile.WriteString("Тестовое содержимое файла")
		tempFile.Close()

		uploadResult, err := client.UploadFile(ctx, tempFile.Name())
		if err != nil {
			log.Printf("   Ошибка загрузки файла: %v", err)
		} else {
			fmt.Printf("   Файл загружен: %+v\n", uploadResult)
		}

		// Удаляем временный файл
		os.Remove(tempFile.Name())
	}

	// Тест 12: Загрузка файла по несуществующему пути
	fmt.Println("\n12. Загрузка файла по несуществующему пути")
	uploadResult2, err := client.UploadFile(ctx, "/несуществующий/путь/файл.txt")
	if err != nil {
		log.Printf("   Ошибка загрузки файла (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файл загружен (неожиданно): %+v\n", uploadResult2)
	}

	// Тест 13: Загрузка файла по URL
	fmt.Println("\n13. Загрузка файла по URL")
	uploadRequest := &models.FileUploadRequest{
		Name: "Файл по URL",
		URL:  "https://example.com/sample.pdf",
	}
	uploadResult3, err := client.UploadFileByLink(ctx, uploadRequest)
	if err != nil {
		log.Printf("   Ошибка загрузки файла по URL: %v", err)
	} else {
		fmt.Printf("   Файл загружен по URL: %+v\n", uploadResult3)
	}

	// Тест 14: Загрузка файла по невалидному URL
	fmt.Println("\n14. Загрузка файла по невалидному URL")
	uploadRequest2 := &models.FileUploadRequest{
		Name: "Файл по невалидному URL",
		URL:  "невалидный-url",
	}
	uploadResult4, err := client.UploadFileByLink(ctx, uploadRequest2)
	if err != nil {
		log.Printf("   Ошибка загрузки файла по URL (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файл загружен по URL (неожиданно): %+v\n", uploadResult4)
	}

	// Тест 15: Загрузка файла с пустым URL
	fmt.Println("\n15. Загрузка файла с пустым URL")
	uploadRequest3 := &models.FileUploadRequest{
		Name: "Файл с пустым URL",
		URL:  "",
	}
	uploadResult5, err := client.UploadFileByLink(ctx, uploadRequest3)
	if err != nil {
		log.Printf("   Ошибка загрузки файла по URL (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файл загружен по URL (неожиданно): %+v\n", uploadResult5)
	}

	// Тест 16: Прикрепление несуществующего файла к задаче
	fmt.Println("\n16. Прикрепление несуществующего файла к задаче")
	err = client.AttachFileToTask(ctx, 999999, 1)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к задаче (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл прикреплен к задаче (неожиданно)")
	}

	// Тест 17: Прикрепление файла к несуществующей задаче
	fmt.Println("\n17. Прикрепление файла к несуществующей задаче")
	err = client.AttachFileToTask(ctx, 1, 999999)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к задаче (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл прикреплен к задаче (неожиданно)")
	}

	// Тест 18: Прикрепление несуществующего файла к контакту
	fmt.Println("\n18. Прикрепление несуществующего файла к контакту")
	err = client.AttachFileToContact(ctx, 999999, 1)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к контакту (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл прикреплен к контакту (неожиданно)")
	}

	// Тест 19: Прикрепление файла к несуществующему контакту
	fmt.Println("\n19. Прикрепление файла к несуществующему контакту")
	err = client.AttachFileToContact(ctx, 1, 999999)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к контакту (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл прикреплен к контакту (неожиданно)")
	}

	// Тест 20: Прикрепление несуществующего файла к проекту
	fmt.Println("\n20. Прикрепление несуществующего файла к проекту")
	err = client.AttachFileToProject(ctx, 999999, 1)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к проекту (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл прикреплен к проекту (неожиданно)")
	}

	// Тест 21: Прикрепление файла к несуществующему проекту
	fmt.Println("\n21. Прикрепление файла к несуществующему проекту")
	err = client.AttachFileToProject(ctx, 1, 999999)
	if err != nil {
		log.Printf("   Ошибка прикрепления файла к проекту (ожидаемо): %v", err)
	} else {
		fmt.Println("   Файл прикреплен к проекту (неожиданно)")
	}

	// Тест 22: Обновление несуществующего файла
	fmt.Println("\n22. Обновление несуществующего файла")
	updateRequest2 := &models.FileUpdateRequest{
		Name: "Попытка обновления несуществующего файла",
	}
	updateResult2, err := client.UpdateFile(ctx, 999999, updateRequest2)
	if err != nil {
		log.Printf("   Ошибка обновления файла (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файл обновлен (неожиданно): %+v\n", updateResult2)
	}

	// Тест 23: Получение файла с несуществующим ID
	fmt.Println("\n23. Получение файла с несуществующим ID")
	nonExistentFile, err := client.GetFileByID(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения файла (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файл получен (неожиданно): %+v\n", nonExistentFile)
	}

	// Тест 24: Загрузка файла с пустым именем
	fmt.Println("\n24. Загрузка файла с пустым именем")
	uploadRequest4 := &models.FileUploadRequest{
		Name: "",
		URL:  "https://example.com/sample.pdf",
	}
	uploadResult6, err := client.UploadFileByLink(ctx, uploadRequest4)
	if err != nil {
		log.Printf("   Ошибка загрузки файла по URL: %v", err)
	} else {
		fmt.Printf("   Файл загружен по URL: %+v\n", uploadResult6)
	}

	// Тест 25: Загрузка файла с очень длинным именем
	fmt.Println("\n25. Загрузка файла с очень длинным именем")
	longName := "Очень длинное имя файла которое может превышать лимиты системы и должно быть обработано корректно без ошибок"
	uploadRequest5 := &models.FileUploadRequest{
		Name: longName,
		URL:  "https://example.com/sample.pdf",
	}
	uploadResult7, err := client.UploadFileByLink(ctx, uploadRequest5)
	if err != nil {
		log.Printf("   Ошибка загрузки файла по URL: %v", err)
	} else {
		fmt.Printf("   Файл загружен по URL: %+v\n", uploadResult7)
	}

	fmt.Println("\n=== Тестирование Files API завершено ===")
}

func main() {
	testFiles()
}
