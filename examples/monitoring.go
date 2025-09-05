package main

import (
	"context"
	"fmt"
	"log"
	"time"

	goplanfix "go-planfix"
)

func testMonitoring() {
	token := "5fd5c11f9ad5ec21ddffd43a26332241"
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Monitoring API ===")

	// Тест 1: Проверка доступности REST сервиса (Ping)
	fmt.Println("\n1. Проверка доступности REST сервиса (Ping)")
	pingResult, err := client.Ping(ctx)
	if err != nil {
		log.Printf("   Ошибка ping: %v", err)
	} else {
		fmt.Printf("   Ping успешен: %+v\n", pingResult)
	}

	// Тест 2: Проверка ping с отмененным контекстом
	fmt.Println("\n2. Проверка ping с отмененным контекстом")
	cancelCtx, cancel := context.WithCancel(ctx)
	cancel() // Отменяем контекст сразу
	pingResult2, err := client.Ping(cancelCtx)
	if err != nil {
		log.Printf("   Ошибка ping с отмененным контекстом (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping успешен (неожиданно): %+v\n", pingResult2)
	}

	// Тест 3: Проверка ping с истекшим контекстом
	fmt.Println("\n3. Проверка ping с истекшим контекстом")
	timeoutCtx, cancel := context.WithTimeout(ctx, 0) // Сразу истекает
	defer cancel()
	pingResult3, err := client.Ping(timeoutCtx)
	if err != nil {
		log.Printf("   Ошибка ping с истекшим контекстом (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping успешен (неожиданно): %+v\n", pingResult3)
	}

	// Тест 4: Проверка ping с nil контекстом
	fmt.Println("\n4. Проверка ping с nil контекстом")
	pingResult4, err := client.Ping(nil)
	if err != nil {
		log.Printf("   Ошибка ping с nil контекстом (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping успешен (неожиданно): %+v\n", pingResult4)
	}

	// Тест 5: Множественные ping запросы для проверки стабильности
	fmt.Println("\n5. Множественные ping запросы для проверки стабильности")
	successCount := 0
	totalRequests := 5

	for i := 0; i < totalRequests; i++ {
		_, err := client.Ping(ctx)
		if err != nil {
			log.Printf("   Ошибка ping #%d: %v", i+1, err)
		} else {
			successCount++
			fmt.Printf("   Ping #%d успешен\n", i+1)
		}
	}

	fmt.Printf("   Успешных ping запросов: %d из %d\n", successCount, totalRequests)

	// Тест 6: Проверка ping с разными таймаутами
	fmt.Println("\n6. Проверка ping с разными таймаутами")
	timeouts := []int{1, 5, 10, 30} // секунды

	for _, timeout := range timeouts {
		timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		_, err := client.Ping(timeoutCtx)
		cancel()

		if err != nil {
			log.Printf("   Ошибка ping с таймаутом %ds: %v", timeout, err)
		} else {
			fmt.Printf("   Ping с таймаутом %ds успешен\n", timeout)
		}
	}

	// Тест 7: Проверка ping с очень коротким таймаутом
	fmt.Println("\n7. Проверка ping с очень коротким таймаутом")
	shortTimeoutCtx, cancel := context.WithTimeout(ctx, 1*time.Millisecond)
	defer cancel()
	pingResult5, err := client.Ping(shortTimeoutCtx)
	if err != nil {
		log.Printf("   Ошибка ping с коротким таймаутом (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping с коротким таймаутом успешен (неожиданно): %+v\n", pingResult5)
	}

	// Тест 8: Проверка ping с очень длинным таймаутом
	fmt.Println("\n8. Проверка ping с очень длинным таймаутом")
	longTimeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()
	pingResult6, err := client.Ping(longTimeoutCtx)
	if err != nil {
		log.Printf("   Ошибка ping с длинным таймаутом: %v", err)
	} else {
		fmt.Printf("   Ping с длинным таймаутом успешен: %+v\n", pingResult6)
	}

	// Тест 9: Проверка ping с контекстом с дедлайном
	fmt.Println("\n9. Проверка ping с контекстом с дедлайном")
	deadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(10*time.Second))
	defer cancel()
	pingResult7, err := client.Ping(deadlineCtx)
	if err != nil {
		log.Printf("   Ошибка ping с дедлайном: %v", err)
	} else {
		fmt.Printf("   Ping с дедлайном успешен: %+v\n", pingResult7)
	}

	// Тест 10: Проверка ping с контекстом с истекшим дедлайном
	fmt.Println("\n10. Проверка ping с контекстом с истекшим дедлайном")
	pastDeadlineCtx, cancel := context.WithDeadline(ctx, time.Now().Add(-1*time.Hour))
	defer cancel()
	pingResult8, err := client.Ping(pastDeadlineCtx)
	if err != nil {
		log.Printf("   Ошибка ping с истекшим дедлайном (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping с истекшим дедлайном успешен (неожиданно): %+v\n", pingResult8)
	}

	// Тест 11: Проверка ping с контекстом с значениями
	fmt.Println("\n11. Проверка ping с контекстом с значениями")
	valueCtx := context.WithValue(ctx, "test_key", "test_value")
	pingResult9, err := client.Ping(valueCtx)
	if err != nil {
		log.Printf("   Ошибка ping с контекстом с значениями: %v", err)
	} else {
		fmt.Printf("   Ping с контекстом с значениями успешен: %+v\n", pingResult9)
	}

	// Тест 12: Проверка ping с контекстом с отменой после задержки
	fmt.Println("\n12. Проверка ping с контекстом с отменой после задержки")
	delayedCancelCtx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	pingResult10, err := client.Ping(delayedCancelCtx)
	if err != nil {
		log.Printf("   Ошибка ping с отменой после задержки (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping с отменой после задержки успешен (неожиданно): %+v\n", pingResult10)
	}

	// Тест 13: Проверка ping с контекстом с таймаутом и отменой
	fmt.Println("\n13. Проверка ping с контекстом с таймаутом и отменой")
	timeoutCancelCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	pingResult11, err := client.Ping(timeoutCancelCtx)
	if err != nil {
		log.Printf("   Ошибка ping с таймаутом и отменой (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping с таймаутом и отменой успешен (неожиданно): %+v\n", pingResult11)
	}

	// Тест 14: Проверка ping с контекстом с дедлайном и отменой
	fmt.Println("\n14. Проверка ping с контекстом с дедлайном и отменой")
	deadlineCancelCtx, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	go func() {
		time.Sleep(1 * time.Second)
		cancel()
	}()
	pingResult12, err := client.Ping(deadlineCancelCtx)
	if err != nil {
		log.Printf("   Ошибка ping с дедлайном и отменой (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Ping с дедлайном и отменой успешен (неожиданно): %+v\n", pingResult12)
	}

	// Тест 15: Проверка ping с контекстом с множественными значениями
	fmt.Println("\n15. Проверка ping с контекстом с множественными значениями")
	multiValueCtx := context.WithValue(ctx, "key1", "value1")
	multiValueCtx = context.WithValue(multiValueCtx, "key2", "value2")
	multiValueCtx = context.WithValue(multiValueCtx, "key3", "value3")
	pingResult13, err := client.Ping(multiValueCtx)
	if err != nil {
		log.Printf("   Ошибка ping с множественными значениями: %v", err)
	} else {
		fmt.Printf("   Ping с множественными значениями успешен: %+v\n", pingResult13)
	}

	// Тест 16: Проверка ping с контекстом с вложенными значениями
	fmt.Println("\n16. Проверка ping с контекстом с вложенными значениями")
	nestedValueCtx := context.WithValue(ctx, "nested", map[string]interface{}{
		"level1": map[string]interface{}{
			"level2": map[string]interface{}{
				"level3": "deep_value",
			},
		},
	})
	pingResult14, err := client.Ping(nestedValueCtx)
	if err != nil {
		log.Printf("   Ошибка ping с вложенными значениями: %v", err)
	} else {
		fmt.Printf("   Ping с вложенными значениями успешен: %+v\n", pingResult14)
	}

	// Тест 17: Проверка ping с контекстом с большими значениями
	fmt.Println("\n17. Проверка ping с контекстом с большими значениями")
	largeValueCtx := context.WithValue(ctx, "large_key", make([]byte, 1024*1024)) // 1MB
	pingResult15, err := client.Ping(largeValueCtx)
	if err != nil {
		log.Printf("   Ошибка ping с большими значениями: %v", err)
	} else {
		fmt.Printf("   Ping с большими значениями успешен: %+v\n", pingResult15)
	}

	// Тест 18: Проверка ping с контекстом с специальными символами в ключах
	fmt.Println("\n18. Проверка ping с контекстом с специальными символами в ключах")
	specialKeyCtx := context.WithValue(ctx, "key-with-dashes", "value1")
	specialKeyCtx = context.WithValue(specialKeyCtx, "key_with_underscores", "value2")
	specialKeyCtx = context.WithValue(specialKeyCtx, "key.with.dots", "value3")
	specialKeyCtx = context.WithValue(specialKeyCtx, "key with spaces", "value4")
	pingResult16, err := client.Ping(specialKeyCtx)
	if err != nil {
		log.Printf("   Ошибка ping с специальными символами: %v", err)
	} else {
		fmt.Printf("   Ping с специальными символами успешен: %+v\n", pingResult16)
	}

	// Тест 19: Проверка ping с контекстом с nil значениями
	fmt.Println("\n19. Проверка ping с контекстом с nil значениями")
	nilValueCtx := context.WithValue(ctx, "nil_key", nil)
	pingResult17, err := client.Ping(nilValueCtx)
	if err != nil {
		log.Printf("   Ошибка ping с nil значениями: %v", err)
	} else {
		fmt.Printf("   Ping с nil значениями успешен: %+v\n", pingResult17)
	}

	// Тест 20: Проверка ping с контекстом с пустыми значениями
	fmt.Println("\n20. Проверка ping с контекстом с пустыми значениями")
	emptyValueCtx := context.WithValue(ctx, "empty_string", "")
	emptyValueCtx = context.WithValue(emptyValueCtx, "empty_slice", []string{})
	emptyValueCtx = context.WithValue(emptyValueCtx, "empty_map", map[string]interface{}{})
	pingResult18, err := client.Ping(emptyValueCtx)
	if err != nil {
		log.Printf("   Ошибка ping с пустыми значениями: %v", err)
	} else {
		fmt.Printf("   Ping с пустыми значениями успешен: %+v\n", pingResult18)
	}

	fmt.Println("\n=== Тестирование Monitoring API завершено ===")
}

func main() {
	testMonitoring()
}
