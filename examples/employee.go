package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "github.com/whatcrm/go-planfix"
	"github.com/whatcrm/go-planfix/models"
)

func testEmployees() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Employees API ===")

	// Тест 1: Получение пользователя по ID
	fmt.Println("\n1. Получение пользователя по ID")
	user, err := client.GetUserByID(ctx, "1", "id,name,email")
	if err != nil {
		log.Printf("   Ошибка получения пользователя: %v", err)
	} else {
		fmt.Printf("   Пользователь получен: %+v\n", user)
	}

	// Тест 2: Получение пользователя без полей
	fmt.Println("\n2. Получение пользователя без полей")
	user2, err := client.GetUserByID(ctx, "2", "")
	if err != nil {
		log.Printf("   Ошибка получения пользователя: %v", err)
	} else {
		fmt.Printf("   Пользователь получен: %+v\n", user2)
	}

	// Тест 3: Обновление пользователя
	fmt.Println("\n3. Обновление пользователя")
	updateUser := &models.User{
		Name:     "Обновленное имя",
		Lastname: "Обновленная фамилия",
		Email:    "updated@example.com",
	}
	updateResult, err := client.UpdateUser(ctx, "1", updateUser)
	if err != nil {
		log.Printf("   Ошибка обновления пользователя: %v", err)
	} else {
		fmt.Printf("   Пользователь обновлен: %+v\n", updateResult)
	}

	// Тест 4: Получение списка пользователей
	fmt.Println("\n4. Получение списка пользователей")
	userListRequest := &models.UserListRequest{
		Offset:     0,
		PageSize:   10,
		OnlyActive: true,
		Fields:     "id,name,email,isActive",
	}
	users, err := client.GetListUsers(ctx, userListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка пользователей: %v", err)
	} else {
		fmt.Printf("   Список пользователей получен: %+v\n", users)
	}

	// Тест 5: Получение списка пользователей без фильтров
	fmt.Println("\n5. Получение списка пользователей без фильтров")
	userListRequest2 := &models.UserListRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	users2, err := client.GetListUsers(ctx, userListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка пользователей: %v", err)
	} else {
		fmt.Printf("   Список пользователей получен: %+v\n", users2)
	}

	// Тест 6: Создание пользователя
	fmt.Println("\n6. Создание пользователя")
	newUser := &models.User{
		Name:     "Новый пользователь",
		Lastname: "Тестовый",
		Email:    "newuser@example.com",
		Login:    "newuser",
		Password: "password123",
		Role:     "user",
		Status:   "active",
	}
	createdUser, err := client.CreateUser(ctx, newUser)
	if err != nil {
		log.Printf("   Ошибка создания пользователя: %v", err)
	} else {
		fmt.Printf("   Пользователь создан: %+v\n", createdUser)
	}

	// Тест 7: Получение должностей пользователей
	fmt.Println("\n7. Получение должностей пользователей")
	positions, err := client.GetUserPositions(ctx, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения должностей: %v", err)
	} else {
		fmt.Printf("   Должности получены: %+v\n", positions)
	}

	// Тест 8: Получение должностей без полей
	fmt.Println("\n8. Получение должностей без полей")
	positions2, err := client.GetUserPositions(ctx)
	if err != nil {
		log.Printf("   Ошибка получения должностей: %v", err)
	} else {
		fmt.Printf("   Должности получены: %+v\n", positions2)
	}

	// Тест 9: Получение групп пользователей
	fmt.Println("\n9. Получение групп пользователей")
	userGroups, err := client.GetListUserGroups(ctx, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения групп пользователей: %v", err)
	} else {
		fmt.Printf("   Группы пользователей получены: %+v\n", userGroups)
	}

	// Тест 10: Получение групп пользователей без полей
	fmt.Println("\n10. Получение групп пользователей без полей")
	userGroups2, err := client.GetListUserGroups(ctx)
	if err != nil {
		log.Printf("   Ошибка получения групп пользователей: %v", err)
	} else {
		fmt.Printf("   Группы пользователей получены: %+v\n", userGroups2)
	}

	// Тест 11: Получение пользователя с несуществующим ID
	fmt.Println("\n11. Получение пользователя с несуществующим ID")
	nonExistentUser, err := client.GetUserByID(ctx, "999999", "id,name")
	if err != nil {
		log.Printf("   Ошибка получения пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь получен (неожиданно): %+v\n", nonExistentUser)
	}

	// Тест 12: Обновление пользователя с несуществующим ID
	fmt.Println("\n12. Обновление пользователя с несуществующим ID")
	updateUser2 := &models.User{
		Name:  "Попытка обновления несуществующего пользователя",
		Email: "nonexistent@example.com",
	}
	updateResult2, err := client.UpdateUser(ctx, "999999", updateUser2)
	if err != nil {
		log.Printf("   Ошибка обновления пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь обновлен (неожиданно): %+v\n", updateResult2)
	}

	// Тест 13: Создание пользователя с дублирующимся email
	fmt.Println("\n13. Создание пользователя с дублирующимся email")
	duplicateUser := &models.User{
		Name:     "Дублирующийся пользователь",
		Lastname: "Тестовый",
		Email:    "duplicate@example.com",
		Login:    "duplicate",
		Password: "password123",
		Role:     "user",
		Status:   "active",
	}
	duplicateResult, err := client.CreateUser(ctx, duplicateUser)
	if err != nil {
		log.Printf("   Ошибка создания пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь создан (неожиданно): %+v\n", duplicateResult)
	}

	// Тест 14: Создание пользователя с невалидными данными
	fmt.Println("\n14. Создание пользователя с невалидными данными")
	invalidUser := &models.User{
		Name:     "",              // Пустое имя
		Email:    "invalid-email", // Невалидный email
		Password: "",              // Пустой пароль
	}
	invalidResult, err := client.CreateUser(ctx, invalidUser)
	if err != nil {
		log.Printf("   Ошибка создания пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь создан (неожиданно): %+v\n", invalidResult)
	}

	// Тест 15: Получение списка пользователей с большим offset
	fmt.Println("\n15. Получение списка пользователей с большим offset")
	userListRequest3 := &models.UserListRequest{
		Offset:     1000,
		PageSize:   10,
		OnlyActive: false,
		Fields:     "id,name",
	}
	users3, err := client.GetListUsers(ctx, userListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка пользователей: %v", err)
	} else {
		fmt.Printf("   Список пользователей получен: %+v\n", users3)
	}

	// Тест 16: Получение списка пользователей с нулевым pageSize
	fmt.Println("\n16. Получение списка пользователей с нулевым pageSize")
	userListRequest4 := &models.UserListRequest{
		Offset:   0,
		PageSize: 0,
		Fields:   "id,name",
	}
	users4, err := client.GetListUsers(ctx, userListRequest4)
	if err != nil {
		log.Printf("   Ошибка получения списка пользователей: %v", err)
	} else {
		fmt.Printf("   Список пользователей получен: %+v\n", users4)
	}

	// Тест 17: Получение пользователя с пустым ID
	fmt.Println("\n17. Получение пользователя с пустым ID")
	emptyUser, err := client.GetUserByID(ctx, "", "id,name")
	if err != nil {
		log.Printf("   Ошибка получения пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь получен (неожиданно): %+v\n", emptyUser)
	}

	// Тест 18: Обновление пользователя с пустым ID
	fmt.Println("\n18. Обновление пользователя с пустым ID")
	updateUser3 := &models.User{
		Name:  "Попытка обновления с пустым ID",
		Email: "empty@example.com",
	}
	updateResult3, err := client.UpdateUser(ctx, "", updateUser3)
	if err != nil {
		log.Printf("   Ошибка обновления пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь обновлен (неожиданно): %+v\n", updateResult3)
	}

	// Тест 19: Создание пользователя с пустыми обязательными полями
	fmt.Println("\n19. Создание пользователя с пустыми обязательными полями")
	emptyUser2 := &models.User{
		Name:     "",
		Lastname: "",
		Email:    "",
		Login:    "",
		Password: "",
	}
	emptyResult, err := client.CreateUser(ctx, emptyUser2)
	if err != nil {
		log.Printf("   Ошибка создания пользователя (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Пользователь создан (неожиданно): %+v\n", emptyResult)
	}

	// Тест 20: Получение должностей с невалидными полями
	fmt.Println("\n20. Получение должностей с невалидными полями")
	positions3, err := client.GetUserPositions(ctx, "invalid,fields")
	if err != nil {
		log.Printf("   Ошибка получения должностей: %v", err)
	} else {
		fmt.Printf("   Должности получены: %+v\n", positions3)
	}

	fmt.Println("\n=== Тестирование Employees API завершено ===")
}

func main() {
	testEmployees()
}
