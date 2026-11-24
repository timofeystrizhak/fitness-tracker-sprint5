package actioninfo

import "log"

// DataParser — интерфейс для структур Training и DaySteps.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info обрабатывает слайс строк и выводит информацию.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Println("ошибка парсинга:", err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Println("ошибка формирования информации:", err)
			continue
		}

		log.Println(info)
	}
}
