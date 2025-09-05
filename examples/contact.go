package main

import (
	"context"
	"fmt"
	goplanfix "go-planfix"
	"log"
)

func main() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	//fmt.Println("1. Создание простого контакта")
	//simpleContact := models.Contact{
	//	Template: &models.ContactTemplate{Id: 1},
	//	Name:     "Simple User",
	//	Email:    "testEmailwwww11Hi@example.com",
	//}
	//
	//result1, err := client.CreateContact(ctx, &simpleContact)
	//if err != nil {
	//	log.Printf("   Ошибка: %v", err)
	//} else {
	//	fmt.Printf("   Создан контакт ID=%d\n", result1.ID)
	//}
	//
	//fmt.Println("1. Создание контакта c tg")
	//telegramContact := models.Contact{
	//	Template:    &models.ContactTemplate{Id: 1},
	//	Name:        "Telegram User",
	//	Midname:     "Test",
	//	Lastname:    "Contact",
	//	Email:       "telegram212121@example.com",
	//	TelegramId:  "123456789",
	//	Telegram:    "https://t.me/test_telegram_user",
	//	Facebook:    "https://facebook.com/testtelegram",
	//	Vk:          "https://vk.com/testtelegram",
	//	Instagram:   "https://instagram.com/testtelegram",
	//	Description: "Contact with social media integration",
	//	Position:    "Developer",
	//}
	//
	//result2, err := client.CreateContact(ctx, &telegramContact)
	//if err != nil {
	//	log.Printf("   Ошибка: %v", err)
	//} else {
	//	fmt.Printf("   Создан Telegram контакт ID=%d\n", result2.ID)
	//}
	//
	//fmt.Println("\n3. Создание контакта с кастомными полями")
	//customFieldContact := models.Contact{
	//	Template: &models.ContactTemplate{Id: 1},
	//	Name:     "Custom Field User",
	//	Midname:  "Test",
	//	Lastname: "Fields",
	//	Email:    "custo33232m@example.com",
	//	CustomFieldData: []interface{}{
	//		map[string]interface{}{
	//			"field": map[string]interface{}{"id": 250},
	//			"value": map[string]interface{}{
	//				"address":   "New York",
	//				"longitude": 40.71041692830323,
	//				"latitude":  -74.00097407111207,
	//			},
	//		},
	//	},
	//}
	//
	//result3, err := client.CreateContact(ctx, &customFieldContact)
	//if err != nil {
	//	log.Printf("   Ошибка: %v", err)
	//} else {
	//	fmt.Printf("   Создан контакт с кастомными полями ID=%d\n", result3.ID)
	//}
	//
	//fmt.Println("\n4. Создание контакта с телефонами")
	//phoneContact := models.Contact{
	//	Template: &models.ContactTemplate{Id: 1},
	//	Name:     "Phone User",
	//	Email:    "phone3232@example.com",
	//	Phones: []models.ContactPhone{
	//		{Number: "89643752416", Type: 1}, // Mobile
	//		{Number: "84953752416", Type: 2}, // Work
	//	},
	//}
	//
	//result4, err := client.CreateContact(ctx, &phoneContact)
	//if err != nil {
	//	log.Printf("   Ошибка: %v", err)
	//} else {
	//	fmt.Printf("   Создан контакт с телефонами ID=%d\n", result4.ID)
	//}

	// Только базовые поля
	contact1, err := client.GetContactByID(ctx, 12, "id,name,email")
	if err != nil {
		log.Printf("   Ошибка получения базовых полей: %v", err)
	} else {
		fmt.Printf("   Базовые поля: %+v\n", contact1)
	}

	// Telegram поля
	contact2, err := client.GetContactByID(ctx, 12, "id,telegramId,telegram,facebook,vk,instagram")
	if err != nil {
		log.Printf("   Ошибка получения Telegram полей: %v", err)
	} else {
		fmt.Printf("   Telegram поля: %+v\n", contact2)
	}

	// Телефоны

	contact3, err := client.GetContactByID(ctx, 12, "id,phones")
	if err != nil {
		log.Printf("   Ошибка получения телефонов: %v", err)
	} else {
		fmt.Printf("   Телефоны: %+v\n", contact3)
	}

	//groups, err := client.GetContactGroups(ctx)
	//if err != nil {
	//	log.Printf("   Ошибка получения групп: %v", err)
	//} else {
	//	fmt.Printf("   Группы контактов: %+v\n", groups)
	//}
	//
	//templates, err := client.GetContactTemplates(ctx)
	//if err != nil {
	//	log.Printf("   Ошибка получения шаблонов: %v", err)
	//} else {
	//	fmt.Printf("   Шаблоны контактов: %+v\n", templates)
	//}
	//
	//filters, err := client.GetContactFilters(ctx)
	//if err != nil {
	//	log.Printf("   Ошибка получения фильтров: %v", err)
	//} else {
	//	fmt.Printf("   Фильтры контактов: %+v\n", filters)
	//}
	//
	//contactFiles, err := client.GetContactFiles(ctx, 11)
	//if err != nil {
	//	log.Printf("   Ошибка получения файлов: %v", err)
	//} else {
	//	fmt.Printf("   Файлы контакта: %+v\n", contactFiles)
	//}
}
