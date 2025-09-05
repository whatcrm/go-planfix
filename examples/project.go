package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "go-planfix"
	"go-planfix/models"
)

func testProjects() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Projects API ===")

	// Тест 1: Получение проекта по ID
	fmt.Println("\n1. Получение проекта по ID")
	project, err := client.GetProjectByID(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения проекта: %v", err)
	} else {
		fmt.Printf("   Проект получен: %+v\n", project)
	}

	// Тест 2: Получение проекта без полей
	fmt.Println("\n2. Получение проекта без полей")
	project2, err := client.GetProjectByID(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения проекта: %v", err)
	} else {
		fmt.Printf("   Проект получен: %+v\n", project2)
	}

	// Тест 3: Обновление проекта
	fmt.Println("\n3. Обновление проекта")
	updateProject := &models.Project{
		Name:        "Обновленный проект",
		Description: "Обновленное описание проекта",
	}
	updateResult, err := client.UpdateProject(ctx, 1, updateProject)
	if err != nil {
		log.Printf("   Ошибка обновления проекта: %v", err)
	} else {
		fmt.Printf("   Проект обновлен: %+v\n", updateResult)
	}

	// Тест 4: Получение списка проектов
	fmt.Println("\n4. Получение списка проектов")
	projectListRequest := &models.ProjectListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,description",
	}
	projects, err := client.GetListsProjects(ctx, projectListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка проектов: %v", err)
	} else {
		fmt.Printf("   Список проектов получен: %+v\n", projects)
	}

	// Тест 5: Получение списка проектов без полей
	fmt.Println("\n5. Получение списка проектов без полей")
	projectListRequest2 := &models.ProjectListRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	projects2, err := client.GetListsProjects(ctx, projectListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка проектов: %v", err)
	} else {
		fmt.Printf("   Список проектов получен: %+v\n", projects2)
	}

	// Тест 6: Создание проекта
	fmt.Println("\n6. Создание проекта")
	newProject := &models.Project{
		Name:        "Новый проект",
		Description: "Описание нового проекта",
	}
	createdProject, err := client.CreateProject(ctx, newProject)
	if err != nil {
		log.Printf("   Ошибка создания проекта: %v", err)
	} else {
		fmt.Printf("   Проект создан: %+v\n", createdProject)
	}

	// Тест 7: Получение файлов проекта
	fmt.Println("\n7. Получение файлов проекта")
	projectFiles, err := client.GetProjectFiles(ctx, 1, 0, 10)
	if err != nil {
		log.Printf("   Ошибка получения файлов проекта: %v", err)
	} else {
		fmt.Printf("   Файлы проекта получены: %+v\n", projectFiles)
	}

	// Тест 8: Получение файлов проекта без параметров
	fmt.Println("\n8. Получение файлов проекта без параметров")
	projectFiles2, err := client.GetProjectFiles(ctx, 1, 0, 0)
	if err != nil {
		log.Printf("   Ошибка получения файлов проекта: %v", err)
	} else {
		fmt.Printf("   Файлы проекта получены: %+v\n", projectFiles2)
	}

	// Тест 9: Получение шаблонов проектов
	fmt.Println("\n9. Получение шаблонов проектов")
	projectTemplates, err := client.GetProjectTemplates(ctx, 0, 10, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения шаблонов проектов: %v", err)
	} else {
		fmt.Printf("   Шаблоны проектов получены: %+v\n", projectTemplates)
	}

	// Тест 10: Получение шаблонов проектов без полей
	fmt.Println("\n10. Получение шаблонов проектов без полей")
	projectTemplates2, err := client.GetProjectTemplates(ctx, 0, 5)
	if err != nil {
		log.Printf("   Ошибка получения шаблонов проектов: %v", err)
	} else {
		fmt.Printf("   Шаблоны проектов получены: %+v\n", projectTemplates2)
	}

	// Тест 11: Получение групп проектов
	fmt.Println("\n11. Получение групп проектов")
	projectGroups, err := client.GetProjectGroups(ctx, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения групп проектов: %v", err)
	} else {
		fmt.Printf("   Группы проектов получены: %+v\n", projectGroups)
	}

	// Тест 12: Получение групп проектов без полей
	fmt.Println("\n12. Получение групп проектов без полей")
	projectGroups2, err := client.GetProjectGroups(ctx)
	if err != nil {
		log.Printf("   Ошибка получения групп проектов: %v", err)
	} else {
		fmt.Printf("   Группы проектов получены: %+v\n", projectGroups2)
	}

	// Тест 13: Получение проекта с несуществующим ID
	fmt.Println("\n13. Получение проекта с несуществующим ID")
	nonExistentProject, err := client.GetProjectByID(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения проекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Проект получен (неожиданно): %+v\n", nonExistentProject)
	}

	// Тест 14: Обновление проекта с несуществующим ID
	fmt.Println("\n14. Обновление проекта с несуществующим ID")
	updateProject2 := &models.Project{
		Name: "Попытка обновления несуществующего проекта",
	}
	updateResult2, err := client.UpdateProject(ctx, 999999, updateProject2)
	if err != nil {
		log.Printf("   Ошибка обновления проекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Проект обновлен (неожиданно): %+v\n", updateResult2)
	}

	// Тест 15: Создание проекта с пустыми обязательными полями
	fmt.Println("\n15. Создание проекта с пустыми обязательными полями")
	emptyProject := &models.Project{
		Name:        "",
		Description: "",
	}
	emptyResult, err := client.CreateProject(ctx, emptyProject)
	if err != nil {
		log.Printf("   Ошибка создания проекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Проект создан (неожиданно): %+v\n", emptyResult)
	}

	// Тест 16: Получение файлов несуществующего проекта
	fmt.Println("\n16. Получение файлов несуществующего проекта")
	nonExistentFiles, err := client.GetProjectFiles(ctx, 999999, 0, 10)
	if err != nil {
		log.Printf("   Ошибка получения файлов проекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файлы проекта получены (неожиданно): %+v\n", nonExistentFiles)
	}

	// Тест 17: Получение шаблонов проектов с невалидными полями
	fmt.Println("\n17. Получение шаблонов проектов с невалидными полями")
	projectTemplates3, err := client.GetProjectTemplates(ctx, 0, 10, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения шаблонов проектов: %v", err)
	} else {
		fmt.Printf("   Шаблоны проектов получены: %+v\n", projectTemplates3)
	}

	// Тест 18: Получение групп проектов с невалидными полями
	fmt.Println("\n18. Получение групп проектов с невалидными полями")
	projectGroups3, err := client.GetProjectGroups(ctx, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения групп проектов: %v", err)
	} else {
		fmt.Printf("   Группы проектов получены: %+v\n", projectGroups3)
	}

	// Тест 19: Получение списка проектов с большим offset
	fmt.Println("\n19. Получение списка проектов с большим offset")
	projectListRequest3 := &models.ProjectListRequest{
		Offset:   1000,
		PageSize: 10,
		Fields:   "id,name",
	}
	projects3, err := client.GetListsProjects(ctx, projectListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка проектов: %v", err)
	} else {
		fmt.Printf("   Список проектов получен: %+v\n", projects3)
	}

	// Тест 20: Получение списка проектов с нулевым pageSize
	fmt.Println("\n20. Получение списка проектов с нулевым pageSize")
	projectListRequest4 := &models.ProjectListRequest{
		Offset:   0,
		PageSize: 0,
		Fields:   "id,name",
	}
	projects4, err := client.GetListsProjects(ctx, projectListRequest4)
	if err != nil {
		log.Printf("   Ошибка получения списка проектов: %v", err)
	} else {
		fmt.Printf("   Список проектов получен: %+v\n", projects4)
	}

	// Тест 21: Получение файлов проекта с отрицательными параметрами
	fmt.Println("\n21. Получение файлов проекта с отрицательными параметрами")
	projectFiles3, err := client.GetProjectFiles(ctx, 1, -1, -1)
	if err != nil {
		log.Printf("   Ошибка получения файлов проекта: %v", err)
	} else {
		fmt.Printf("   Файлы проекта получены: %+v\n", projectFiles3)
	}

	// Тест 22: Получение шаблонов проектов с отрицательными параметрами
	fmt.Println("\n22. Получение шаблонов проектов с отрицательными параметрами")
	projectTemplates4, err := client.GetProjectTemplates(ctx, -1, -1, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения шаблонов проектов: %v", err)
	} else {
		fmt.Printf("   Шаблоны проектов получены: %+v\n", projectTemplates4)
	}

	// Тест 23: Получение проекта с нулевым ID
	fmt.Println("\n23. Получение проекта с нулевым ID")
	zeroProject, err := client.GetProjectByID(ctx, 0, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения проекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Проект получен (неожиданно): %+v\n", zeroProject)
	}

	// Тест 24: Получение проекта с отрицательным ID
	fmt.Println("\n24. Получение проекта с отрицательным ID")
	negativeProject, err := client.GetProjectByID(ctx, -1, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения проекта (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Проект получен (неожиданно): %+v\n", negativeProject)
	}

	// Тест 25: Создание проекта с очень длинным именем
	fmt.Println("\n25. Создание проекта с очень длинным именем")
	longNameProject := &models.Project{
		Name:        "Очень длинное имя проекта которое может превышать лимиты системы и должно быть обработано корректно без ошибок",
		Description: "Описание проекта с длинным именем",
	}
	longNameResult, err := client.CreateProject(ctx, longNameProject)
	if err != nil {
		log.Printf("   Ошибка создания проекта: %v", err)
	} else {
		fmt.Printf("   Проект создан: %+v\n", longNameResult)
	}

	fmt.Println("\n=== Тестирование Projects API завершено ===")
}

func main() {
	testProjects()
}
