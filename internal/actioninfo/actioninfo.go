package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	// TODO: добавить методы
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("parse error: %v (data=%q)", err, data)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("action info error: %v (row=%q)", err, data)
			continue
		}
		fmt.Println(info)
	}
}
