package main

import (
	"context"
	"fmt"
	"log"

	goplanfix "go-planfix"
	"go-planfix/models"
)

func testReports() {
	token := ""
	client := goplanfix.NewClient(token)
	ctx := context.Background()

	fmt.Println("=== Тестирование Reports API ===")

	// Тест 1: Получение отчета по ID
	fmt.Println("\n1. Получение отчета по ID")
	report, err := client.GetReportByID(ctx, 1, "id,name,description")
	if err != nil {
		log.Printf("   Ошибка получения отчета: %v", err)
	} else {
		fmt.Printf("   Отчет получен: %+v\n", report)
	}

	// Тест 2: Получение отчета без полей
	fmt.Println("\n2. Получение отчета без полей")
	report2, err := client.GetReportByID(ctx, 2, "")
	if err != nil {
		log.Printf("   Ошибка получения отчета: %v", err)
	} else {
		fmt.Printf("   Отчет получен: %+v\n", report2)
	}

	// Тест 3: Получение списка отчетов
	fmt.Println("\n3. Получение списка отчетов")
	reportListRequest := &models.ReportListRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,description",
	}
	reports, err := client.GetReports(ctx, reportListRequest)
	if err != nil {
		log.Printf("   Ошибка получения списка отчетов: %v", err)
	} else {
		fmt.Printf("   Список отчетов получен: %+v\n", reports)
	}

	// Тест 4: Получение списка отчетов без полей
	fmt.Println("\n4. Получение списка отчетов без полей")
	reportListRequest2 := &models.ReportListRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	reports2, err := client.GetReports(ctx, reportListRequest2)
	if err != nil {
		log.Printf("   Ошибка получения списка отчетов: %v", err)
	} else {
		fmt.Printf("   Список отчетов получен: %+v\n", reports2)
	}

	// Тест 5: Получение сохранений отчета
	fmt.Println("\n5. Получение сохранений отчета")
	reportSaves, err := client.GetReportSaves(ctx, 1)
	if err != nil {
		log.Printf("   Ошибка получения сохранений отчета: %v", err)
	} else {
		fmt.Printf("   Сохранения отчета получены: %+v\n", reportSaves)
	}

	// Тест 6: Получение данных сохранения отчета
	fmt.Println("\n6. Получение данных сохранения отчета")
	reportSaveDataRequest := &models.ReportSaveDataRequest{
		Offset:   0,
		PageSize: 10,
		Fields:   "id,name,value",
	}
	reportSaveData, err := client.GetReportSaveData(ctx, 1, 1, reportSaveDataRequest)
	if err != nil {
		log.Printf("   Ошибка получения данных сохранения отчета: %v", err)
	} else {
		fmt.Printf("   Данные сохранения отчета получены: %+v\n", reportSaveData)
	}

	// Тест 7: Получение данных сохранения отчета без полей
	fmt.Println("\n7. Получение данных сохранения отчета без полей")
	reportSaveDataRequest2 := &models.ReportSaveDataRequest{
		Offset:   0,
		PageSize: 5,
		Fields:   "",
	}
	reportSaveData2, err := client.GetReportSaveData(ctx, 1, 1, reportSaveDataRequest2)
	if err != nil {
		log.Printf("   Ошибка получения данных сохранения отчета: %v", err)
	} else {
		fmt.Printf("   Данные сохранения отчета получены: %+v\n", reportSaveData2)
	}

	// Тест 8: Генерация отчета
	fmt.Println("\n8. Генерация отчета")
	reportGenerateRequest := &models.ReportGenerateRequest{}
	reportGenerate, err := client.GenerateReport(ctx, 1, reportGenerateRequest)
	if err != nil {
		log.Printf("   Ошибка генерации отчета: %v", err)
	} else {
		fmt.Printf("   Отчет сгенерирован: %+v\n", reportGenerate)
	}

	// Тест 9: Генерация отчета с пустыми параметрами
	fmt.Println("\n9. Генерация отчета с пустыми параметрами")
	reportGenerateRequest2 := &models.ReportGenerateRequest{}
	reportGenerate2, err := client.GenerateReport(ctx, 1, reportGenerateRequest2)
	if err != nil {
		log.Printf("   Ошибка генерации отчета: %v", err)
	} else {
		fmt.Printf("   Отчет сгенерирован: %+v\n", reportGenerate2)
	}

	// Тест 10: Получение статуса отчета
	fmt.Println("\n10. Получение статуса отчета")
	reportStatus, err := client.GetReportStatus(ctx, "test-request-id")
	if err != nil {
		log.Printf("   Ошибка получения статуса отчета: %v", err)
	} else {
		fmt.Printf("   Статус отчета получен: %+v\n", reportStatus)
	}

	// Тест 11: Получение статуса отчета с пустым ID
	fmt.Println("\n11. Получение статуса отчета с пустым ID")
	reportStatus2, err := client.GetReportStatus(ctx, "")
	if err != nil {
		log.Printf("   Ошибка получения статуса отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статус отчета получен (неожиданно): %+v\n", reportStatus2)
	}

	// Тест 12: Получение отчета с несуществующим ID
	fmt.Println("\n12. Получение отчета с несуществующим ID")
	nonExistentReport, err := client.GetReportByID(ctx, 999999, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Отчет получен (неожиданно): %+v\n", nonExistentReport)
	}

	// Тест 13: Получение сохранений несуществующего отчета
	fmt.Println("\n13. Получение сохранений несуществующего отчета")
	nonExistentSaves, err := client.GetReportSaves(ctx, 999999)
	if err != nil {
		log.Printf("   Ошибка получения сохранений отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Сохранения отчета получены (неожиданно): %+v\n", nonExistentSaves)
	}

	// Тест 14: Получение данных сохранения несуществующего отчета
	fmt.Println("\n14. Получение данных сохранения несуществующего отчета")
	nonExistentSaveData, err := client.GetReportSaveData(ctx, 999999, 1, reportSaveDataRequest)
	if err != nil {
		log.Printf("   Ошибка получения данных сохранения отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Данные сохранения отчета получены (неожиданно): %+v\n", nonExistentSaveData)
	}

	// Тест 15: Генерация несуществующего отчета
	fmt.Println("\n15. Генерация несуществующего отчета")
	nonExistentGenerate, err := client.GenerateReport(ctx, 999999, reportGenerateRequest)
	if err != nil {
		log.Printf("   Ошибка генерации отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Отчет сгенерирован (неожиданно): %+v\n", nonExistentGenerate)
	}

	// Тест 16: Получение статуса несуществующего отчета
	fmt.Println("\n16. Получение статуса несуществующего отчета")
	nonExistentStatus, err := client.GetReportStatus(ctx, "non-existent-request-id")
	if err != nil {
		log.Printf("   Ошибка получения статуса отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Статус отчета получен (неожиданно): %+v\n", nonExistentStatus)
	}

	// Тест 17: Получение отчета с нулевым ID
	fmt.Println("\n17. Получение отчета с нулевым ID")
	zeroReport, err := client.GetReportByID(ctx, 0, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Отчет получен (неожиданно): %+v\n", zeroReport)
	}

	// Тест 18: Получение отчета с отрицательным ID
	fmt.Println("\n18. Получение отчета с отрицательным ID")
	negativeReport, err := client.GetReportByID(ctx, -1, "id,name")
	if err != nil {
		log.Printf("   Ошибка получения отчета (ожидаемо): %v", err)
	} else {
		fmt.Printf("   Отчет получен (неожиданно): %+v\n", negativeReport)
	}

	// Тест 19: Получение списка отчетов с большим offset
	fmt.Println("\n19. Получение списка отчетов с большим offset")
	reportListRequest3 := &models.ReportListRequest{
		Offset:   1000,
		PageSize: 10,
		Fields:   "id,name",
	}
	reports3, err := client.GetReports(ctx, reportListRequest3)
	if err != nil {
		log.Printf("   Ошибка получения списка отчетов: %v", err)
	} else {
		fmt.Printf("   Список отчетов получен: %+v\n", reports3)
	}

	// Тест 20: Получение списка отчетов с нулевым pageSize
	fmt.Println("\n20. Получение списка отчетов с нулевым pageSize")
	reportListRequest4 := &models.ReportListRequest{
		Offset:   0,
		PageSize: 0,
		Fields:   "id,name",
	}
	reports4, err := client.GetReports(ctx, reportListRequest4)
	if err != nil {
		log.Printf("   Ошибка получения списка отчетов: %v", err)
	} else {
		fmt.Printf("   Список отчетов получен: %+v\n", reports4)
	}

	// Тест 21: Получение данных сохранения отчета с отрицательными параметрами
	fmt.Println("\n21. Получение данных сохранения отчета с отрицательными параметрами")
	reportSaveDataRequest3 := &models.ReportSaveDataRequest{
		Offset:   -1,
		PageSize: -1,
		Fields:   "id,name",
	}
	reportSaveData3, err := client.GetReportSaveData(ctx, 1, 1, reportSaveDataRequest3)
	if err != nil {
		log.Printf("   Ошибка получения данных сохранения отчета: %v", err)
	} else {
		fmt.Printf("   Данные сохранения отчета получены: %+v\n", reportSaveData3)
	}

	// Тест 22: Генерация отчета с невалидными параметрами
	fmt.Println("\n22. Генерация отчета с невалидными параметрами")
	reportGenerateRequest3 := &models.ReportGenerateRequest{}
	reportGenerate3, err := client.GenerateReport(ctx, 1, reportGenerateRequest3)
	if err != nil {
		log.Printf("   Ошибка генерации отчета: %v", err)
	} else {
		fmt.Printf("   Отчет сгенерирован: %+v\n", reportGenerate3)
	}

	// Тест 23: Получение статуса отчета с очень длинным ID
	fmt.Println("\n23. Получение статуса отчета с очень длинным ID")
	longRequestID := "very-long-request-id-that-might-cause-issues-with-the-api-and-should-be-handled-correctly"
	reportStatus3, err := client.GetReportStatus(ctx, longRequestID)
	if err != nil {
		log.Printf("   Ошибка получения статуса отчета: %v", err)
	} else {
		fmt.Printf("   Статус отчета получен: %+v\n", reportStatus3)
	}

	// Тест 24: Получение статуса отчета с ID, содержащим специальные символы
	fmt.Println("\n24. Получение статуса отчета с ID, содержащим специальные символы")
	specialRequestID := "request-id-with-special-chars-!@#$%^&*()_+-=[]{}|;':\",./<>?"
	reportStatus4, err := client.GetReportStatus(ctx, specialRequestID)
	if err != nil {
		log.Printf("   Ошибка получения статуса отчета: %v", err)
	} else {
		fmt.Printf("   Статус отчета получен: %+v\n", reportStatus4)
	}

	// Тест 25: Генерация отчета с очень большими параметрами
	fmt.Println("\n25. Генерация отчета с очень большими параметрами")
	reportGenerateRequest4 := &models.ReportGenerateRequest{}
	reportGenerate4, err := client.GenerateReport(ctx, 1, reportGenerateRequest4)
	if err != nil {
		log.Printf("   Ошибка генерации отчета: %v", err)
	} else {
		fmt.Printf("   Отчет сгенерирован: %+v\n", reportGenerate4)
	}

	fmt.Println("\n=== Тестирование Reports API завершено ===")
}

func main() {
	testReports()
}
