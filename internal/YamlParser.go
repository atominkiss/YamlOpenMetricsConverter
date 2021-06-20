package internal

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

// Data - слайс из currency. Такая область видимости, чтобы данные были видны в шаблоне на Frontend'e.
type Data struct {
	Currencies []Currency `yaml:"currencies"`
}

// Currency - наименование, значение (используем float64 так как этот формат используется в OpenMetrics, как я понял из документации).
//Такая область видимости, чтобы данные были видны в шаблоне на Frontend'e.
type Currency struct {
	Name  string  `yaml:"name"`
	Value float64 `yaml:"value"`
}

/*
 YamlParser parse yaml file to structure and return this parsed structure.
 Парсим файл с данными. Из условия мы не знаем будут ли файлы появляться в директории, так что смысла в многопоточке нет.
 Передаем распаресенную структуру хендлеру на фронт.
 Согласно https://prometheus.io/docs/instrumenting/exposition_formats/ необходимо использовать текстовый формат,
 так что вариант с gRPC и protobuff откладываем.
*/
func YamlParser() Data {
	t := Data{}
	//m := make(map[string]interface{})

	//TODO: файл захардкожен для простоты, можно изменить на сканирование директории на наличие файлов.
	// Было бы полезно при работе в многопоточке. Но в условии нет данных о том, могут ли появляться новые файлы в директории.
	file, _ := os.Open("./config/Data.yaml")

	// TODO: Обработать возможную ошибку при закрытии файла??? Но с другой стороны - многопоточки нет,
	//  так что некому этот файл еще читать.
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	/*
		Из условия не понятно, могут значения повторяться или нет. Если не могут, то тогда можно было бы анмаршалить в мапу,
		но в мапе порядок следования данных на выходе не гарантируется. На сколько это критично?
		Если предположить что идет речь о системе мониторинга, то порядок следования значений все таки важен.

	*/
	err = yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return t
	/*
		Анмаршалинг в мапу
		err = yaml.Unmarshal([]byte(data), &m)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		fmt.Println(m)
		return m
	*/
}
