package main

import (
	"context"
	"fmt"
	goplanfix "github.com/whatcrm/go-planfix"
	"github.com/whatcrm/go-planfix/models"
	"log"
)

func RunCustomFieldExamples() {
	token := "14346af5498357e860e1a585c2117860"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Custom Field API Examples ===\n")

	// 1. Получение типов кастомных полей
	fmt.Println("1. Типы кастомных полей")
	fieldTypes, err := client.GetCustomFieldTypes(ctx)
	if err != nil {
		log.Printf("   Ошибка получения типов полей: %v", err)
	} else {
		fmt.Printf("   Типы полей: %+v\n", fieldTypes)
	}

	// 2. Получение основных кастомных полей
	fmt.Println("\n2. Основные кастомные поля")
	mainFields, err := client.GetMainCustomFields(ctx, "id,name,type,objectType")
	if err != nil {
		log.Printf("   Ошибка получения основных полей: %v", err)
	} else {
		fmt.Printf("   Основные поля: %+v\n", mainFields)
	}

	// 3. Создание основного кастомного поля
	fmt.Println("\n3. Создание основного поля")
	mainField := &models.CustomField{
		Name: "API Test Field",
		Type: 0,
	}

	createMainResult, err := client.CreateMainCustomField(ctx, mainField)
	if err != nil {
		log.Printf("   Ошибка создания основного поля: %v", err)
	} else {
		fmt.Printf("   Основное поле создано: %+v\n", createMainResult)
	}

	// 4. Кастомные поля задач
	fmt.Println("\n4. Кастомные поля задач")
	taskFields, err := client.GetTaskCustomFields(ctx, "id,name,type,objectType")
	if err != nil {
		log.Printf("   Ошибка получения полей задач: %v", err)
	} else {
		fmt.Printf("   Поля задач: %+v\n", taskFields)
	}

	// Получение полей для конкретной задачи
	taskFieldsForTask, err := client.GetCustomFieldsForTask(ctx, 1, "id,name,type")
	if err != nil {
		log.Printf("   Ошибка получения полей для задачи: %v", err)
	} else {
		fmt.Printf("   Поля для задачи ID=1: %+v\n", taskFieldsForTask)
	}

	// Создание поля для задачи
	taskField := &models.CustomField{
		Name: "Task API Field",
		Type: 0, // Number type
	}

	createTaskFieldResult, err := client.CreateTaskCustomField(ctx, taskField)
	if err != nil {
		log.Printf("   Ошибка создания поля задачи: %v", err)
	} else {
		fmt.Printf("   Поле задачи создано: %+v\n", createTaskFieldResult)
	}

	// Наборы полей задач
	taskFieldSets, err := client.GetTaskCustomFieldSets(ctx)
	if err != nil {
		log.Printf("   Ошибка получения наборов полей задач: %v", err)
	} else {
		fmt.Printf("   Наборы полей задач: %+v\n", taskFieldSets)
	}

	// 5. Кастомные поля контактов
	fmt.Println("\n5. Кастомные поля контактов")
	contactFields, err := client.GetContactCustomFields(ctx, "id,name,type,objectType")
	if err != nil {
		log.Printf("   Ошибка получения полей контактов: %v", err)
	} else {
		fmt.Printf("   Поля контактов: %+v\n", contactFields)
	}

	// Получение полей для конкретного контакта
	contactFieldsForContact, err := client.GetCustomFieldsForContact(ctx, "11", "id,name,type")
	if err != nil {
		log.Printf("   Ошибка получения полей для контакта: %v", err)
	} else {
		fmt.Printf("   Поля для контакта ID=11: %+v\n", contactFieldsForContact)
	}

	// Создание поля для контакта
	contactField := &models.CustomField{
		Name: "Contact API Field",
		Type: 7, // Checkbox type
	}

	createContactFieldResult, err := client.CreateContactCustomField(ctx, contactField)
	if err != nil {
		log.Printf("   Ошибка создания поля контакта: %v", err)
	} else {
		fmt.Printf("   Поле контакта создано: %+v\n", createContactFieldResult)
	}

	// 6. Кастомные поля проектов
	fmt.Println("\n6. Кастомные поля проектов")
	projectFields, err := client.GetProjectCustomFields(ctx, "id,name,type,objectType")
	if err != nil {
		log.Printf("   Ошибка получения полей проектов: %v", err)
	} else {
		fmt.Printf("   Поля проектов: %+v\n", projectFields)
	}

	// 7. Кастомные поля пользователей
	fmt.Println("\n7. Кастомные поля пользователей")
	userFields, err := client.GetUserCustomFields(ctx, "id,name,type,objectType")
	if err != nil {
		log.Printf("   Ошибка получения полей пользователей: %v", err)
	} else {
		fmt.Printf("   Поля пользователей: %+v\n", userFields)
	}

	// 8. Кастомные поля для data tag
	fmt.Println("\n8. Кастомные поля для data tag")
	dataTagFields, err := client.GetCustomFieldsForDataTag(ctx, 1, "id,name,type")
	if err != nil {
		log.Printf("   Ошибка получения полей data tag: %v", err)
	} else {
		fmt.Printf("   Поля data tag: %+v\n", dataTagFields)
	}

	// 9. Кастомные поля для директории
	fmt.Println("\n9. Кастомные поля для директории")
	directoryFields, err := client.GetCustomFieldsForDirectory(ctx, 1, "id,name,type")
	if err != nil {
		log.Printf("   Ошибка получения полей директории: %v", err)
	} else {
		fmt.Printf("   Поля директории: %+v\n", directoryFields)
	}

	fmt.Println("\n=== Custom Field Examples завершены ===")
}
