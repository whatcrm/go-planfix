package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "go-planfix"
	"go-planfix/models"
)

func testTasks() {
	token := ""
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Tasks API ===")

	// Тест 1: Получение задачи по ID
	fmt.Println("\n1. Получение задачи по ID")
	task, err := client.GetTaskByID(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения задачи: %v", err)
	} else {
		fmt.Printf("   Задача получена: %+v\n", task)
	}

	// Тест 2: Получение задачи без полей
	fmt.Println("\n2. Получение задачи без полей")
	task2, err := client.GetTaskByID(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения задачи: %v", err)
	} else {
		fmt.Printf("   Задача получена: %+v\n", task2)
	}

	// Тест 3: Обновление задачи
	fmt.Println("\n3. Обновление задачи")
	updateTask := &models.Task{
		Name:        "Обновленная задача",
		Description: "Обновленное описание задачи",
	}
	updateResult, err := client.UpdateTask(ctx, 1, updateTask)
	if err != nil {
		log.Printf("   Ошибка обновления задачи: %v", err)
	} else {
		fmt.Printf("   Задача обновлена: %+v\n", updateResult)
	}

	// Тест 4: Получение списка задач
	fmt.Println("\n4. Получение списка задач")
	taskListRequest := &models.TaskListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,description",
	}
	tasks, err := client.GetListTasks(ctx, taskListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка задач: %v", err)
	} else {
		fmt.Printf("   Список задач получен: %+v\n", tasks)
	}

	// Тест 5: Получение списка задач без полей
	fmt.Println("\n5. Получение списка задач без полей")
	taskListRequest2 := &models.TaskListRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	tasks2, err := client.GetListTasks(ctx, taskListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка задач: %v", err)
	} else {
		fmt.Printf("   Список задач получен: %+v\n", tasks2)
	}

	// Тест 6: Создание задачи
	fmt.Println("\n6. Создание задачи")
	newTask := &models.Task{
		Name:        "Новая задача",
		Description: "Описание новой задачи",
	}
	createdTask, err := client.CreateTask(ctx, newTask)
	if err != nil {
		log.Printf("   Ошибка создания задачи: %v", err)
	} else {
		fmt.Printf("   Задача создана: %+v\n", createdTask)
	}

	// Тест 7: Получение комментариев задачи
	fmt.Println("\n7. Получение комментариев задачи")
	commentListRequest := &models.CommentListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,description,owner",
	}
	taskComments, err := client.GetTaskComments(ctx, 1, commentListRequest)
	if err != nil {
		log.Printf("   Ошибка получения комментариев задачи: %v", err)
	} else {
		fmt.Printf("   Комментарии задачи получены: %+v\n", taskComments)
	}

	// Тест 8: Получение комментариев задачи без полей
	fmt.Println("\n8. Получение комментариев задачи без полей")
	commentListRequest2 := &models.CommentListRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	taskComments2, err := client.GetTaskComments(ctx, 1, commentListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения комментариев задачи: %v", err)
	} else {
		fmt.Printf("   Комментарии задачи получены: %+v\n", taskComments2)
	}

	// Тест 9: Добавление комментария к задаче
	fmt.Println("\n9. Добавление комментария к задаче")
	commentRequest := &models.CommentRequest{
		Description: "Новый комментарий к задаче",
	}
	commentResult, err := client.AddCommentTask(ctx, 1, commentRequest)
	if err != nil {
		log.Printf("   Ошибка добавления комментария к задаче: %v", err)
	} else {
		fmt.Printf("   Комментарий добавлен: %+v\n", commentResult)
	}

	// Тест 10: Обновление комментария задачи
	fmt.Println("\n10. Обновление комментария задачи")
	updateCommentRequest := &models.CommentRequest{
		Description: "Обновленный комментарий к задаче",
	}
	err = client.UpdateCommentTask(ctx, 1, 1, updateCommentRequest)
	if err != nil {
		log.Printf("   Ошибка обновления комментария задачи: %v", err)
	} else {
		fmt.Println("   Комментарий задачи обновлен")
	}

	// Тест 11: Добавление data tag с новым комментарием
	fmt.Println("\n11. Добавление data tag с новым комментарием")
	dataTagRequest := &models.DataTagEntryRequest{
		DataTag: &models.BaseEntity{ID: 1},
		Items: []struct {
			CustomFieldData []models.CustomFieldValueRequest `json:"customFieldData"`
		}{},
	}
	dataTagResult, err := client.AddDataTagNewComment(ctx, 1, dataTagRequest)
	if err != nil {
		log.Printf("   Ошибка добавления data tag: %v", err)
	} else {
		fmt.Printf("   Data tag добавлен: %+v\n", dataTagResult)
	}

	// Тест 12: Добавление data tag к существующему комментарию
	fmt.Println("\n12. Добавление data tag к существующему комментарию")
	dataTagRequest2 := &models.DataTagEntryRequest{
		DataTag: &models.BaseEntity{ID: 1},
		Items: []struct {
			CustomFieldData []models.CustomFieldValueRequest `json:"customFieldData"`
		}{},
	}
	dataTagResult2, err := client.AddDataTagExistingComment(ctx, 1, 1, dataTagRequest2)
	if err != nil {
		log.Printf("   Ошибка добавления data tag к комментарию: %v", err)
	} else {
		fmt.Printf("   Data tag добавлен к комментарию: %+v\n", dataTagResult2)
	}

	// Тест 13: Получение файлов задачи
	fmt.Println("\n13. Получение файлов задачи")
	taskFiles, err := client.GetTaskFiles(ctx, 1, false)
	if err != nil {
		log.Printf("   Ошибка получения файлов задачи: %v", err)
	} else {
		fmt.Printf("   Файлы задачи получены: %+v\n", taskFiles)
	}

	// Тест 14: Получение файлов задачи только из описания
	fmt.Println("\n14. Получение файлов задачи только из описания")
	taskFiles2, err := client.GetTaskFiles(ctx, 1, true)
	if err != nil {
		log.Printf("   Ошибка получения файлов задачи: %v", err)
	} else {
		fmt.Printf("   Файлы задачи получены: %+v\n", taskFiles2)
	}

	// Тест 15: Получение шаблонов задач
	fmt.Println("\n15. Получение шаблонов задач")
	taskTemplates, err := client.GetTaskTemplates(ctx, 0, 10, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения шаблонов задач: %v", err)
	} else {
		fmt.Printf("   Шаблоны задач получены: %+v\n", taskTemplates)
	}

	// Тест 16: Получение шаблонов задач без полей
	fmt.Println("\n16. Получение шаблонов задач без полей")
	taskTemplates2, err := client.GetTaskTemplates(ctx, 0, 5)
	if err != nil {
		log.Printf("   Ошибка получения шаблонов задач: %v", err)
	} else {
		fmt.Printf("   Шаблоны задач получены: %+v\n", taskTemplates2)
	}

	// Тест 17: Получение повторяющихся задач
	fmt.Println("\n17. Получение повторяющихся задач")
	recurringTasks, err := client.GetRecurringTasks(ctx, 0, 10, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения повторяющихся задач: %v", err)
	} else {
		fmt.Printf("   Повторяющиеся задачи получены: %+v\n", recurringTasks)
	}

	// Тест 18: Получение повторяющихся задач без полей
	fmt.Println("\n18. Получение повторяющихся задач без полей")
	recurringTasks2, err := client.GetRecurringTasks(ctx, 0, 5)
	if err != nil {
		log.Printf("   Ошибка получения повторяющихся задач: %v", err)
	} else {
		fmt.Printf("   Повторяющиеся задачи получены: %+v\n", recurringTasks2)
	}

	// Тест 19: Получение фильтров задач
	fmt.Println("\n19. Получение фильтров задач")
	taskFilters, err := client.GetTaskFilters(ctx)
	if err != nil {
		log.Printf("   Ошибка получения фильтров задач: %v", err)
	} else {
		fmt.Printf("   Фильтры задач получены: %+v\n", taskFilters)
	}

	// Тест 20: Получение чек-листа задачи
	fmt.Println("\n20. Получение чек-листа задачи")
	checklistRequest := &models.ChecklistRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,completed",
	}
	taskChecklist, err := client.GetTaskChecklist(ctx, 1, checklistRequest)
	if err != nil {
		log.Printf("   Ошибка получения чек-листа задачи: %v", err)
	} else {
		fmt.Printf("   Чек-лист задачи получен: %+v\n", taskChecklist)
	}

	// Тест 21: Получение задачи с несуществующим ID
	fmt.Println("\n21. Получение задачи с несуществующим ID")
	nonExistentTask, err := client.GetTaskByID(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения задачи (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Задача получена (неожиданно): %+v\n", nonExistentTask)
	}

	// Тест 22: Обновление задачи с несуществующим ID
	fmt.Println("\n22. Обновление задачи с несуществующим ID")
	updateTask2 := &models.Task{
		Name: "Попытка обновления несуществующей задачи",
	}
	updateResult2, err := client.UpdateTask(ctx, 999999, updateTask2)
	if err != nil {
		log.Printf("   Ошибка обновления задачи (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Задача обновлена (неожиданно): %+v\n", updateResult2)
	}

	// Тест 23: Создание задачи с пустыми обязательными полями
	fmt.Println("\n23. Создание задачи с пустыми обязательными полями")
	emptyTask := &models.Task{
		Name:        "",
		Description: "",
	}
	emptyResult, err := client.CreateTask(ctx, emptyTask)
	if err != nil {
		log.Printf("   Ошибка создания задачи (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Задача создана (неожиданно): %+v\n", emptyResult)
	}

	// Тест 24: Получение комментариев несуществующей задачи
	fmt.Println("\n24. Получение комментариев несуществующей задачи")
	nonExistentComments, err := client.GetTaskComments(ctx, 999999, commentListRequest)
	if err != nil {
		log.Printf("   Ошибка получения комментариев задачи (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Комментарии задачи получены (неожиданно): %+v\n", nonExistentComments)
	}

	// Тест 25: Добавление комментария к несуществующей задаче
	fmt.Println("\n25. Добавление комментария к несуществующей задаче")
	commentRequest2 := &models.CommentRequest{
		Description: "Комментарий к несуществующей задаче",
	}
	commentResult2, err := client.AddCommentTask(ctx, 999999, commentRequest2)
	if err != nil {
		log.Printf("   Ошибка добавления комментария к задаче (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Комментарий добавлен (неожиданно): %+v\n", commentResult2)
	}

	// Тест 26: Обновление комментария несуществующей задачи
	fmt.Println("\n26. Обновление комментария несуществующей задачи")
	updateCommentRequest2 := &models.CommentRequest{
		Description: "Обновление комментария несуществующей задачи",
	}
	err = client.UpdateCommentTask(ctx, 999999, 1, updateCommentRequest2)
	if err != nil {
		log.Printf("   Ошибка обновления комментария задачи (ожидаемо): %v", err)
	} else {
		fmt.Println("   Комментарий задачи обновлен (неожиданно)")
	}

	// Тест 27: Получение файлов несуществующей задачи
	fmt.Println("\n27. Получение файлов несуществующей задачи")
	nonExistentFiles, err := client.GetTaskFiles(ctx, 999999, false)
	if err != nil {
		log.Printf("   Ошибка получения файлов задачи (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Файлы задачи получены (неожиданно): %+v\n", nonExistentFiles)
	}

	// Тест 28: Получение чек-листа несуществующей задачи
	fmt.Println("\n28. Получение чек-листа несуществующей задачи")
	nonExistentChecklist, err := client.GetTaskChecklist(ctx, 999999, checklistRequest)
	if err != nil {
		log.Printf("   Ошибка получения чек-листа задачи (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Чек-лист задачи получен (неожиданно): %+v\n", nonExistentChecklist)
	}

	// Тест 29: Получение списка задач с большим offset
	fmt.Println("\n29. Получение списка задач с большим offset")
	taskListRequest3 := &models.TaskListRequest{
		Offset:   1000,
		PageSize: 10,
		Fields:   "id,name",
	}
	tasks3, err := client.GetListTasks(ctx, taskListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка задач: %v", err)
	} else {
		fmt.Printf("   Список задач получен: %+v\n", tasks3)
	}

	// Тест 30: Получение списка задач с нулевым pageSize
	fmt.Println("\n30. Получение списка задач с нулевым pageSize")
	taskListRequest4 := &models.TaskListRequest{
		Offset:   0,
		PageSize: 0,
		Fields:   "id,name",
	}
	tasks4, err := client.GetListTasks(ctx, taskListRequest4)
	if err != nil {
		log.Printf("   Ошибка получения списка задач: %v", err)
	} else {
		fmt.Printf("   Список задач получен: %+v\n", tasks4)
	}

	fmt.Println("\n=== Тестирование Tasks API завершено ===")
}

func main() {
	testTasks()
}
