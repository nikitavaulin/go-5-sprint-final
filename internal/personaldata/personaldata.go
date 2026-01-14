// personaldata пакет для работы с данными пользователя
package personaldata

import "fmt"

// Personal структура представляющая с данными о пользователе
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

func (p Personal) String() string {
	return fmt.Sprintf("Имя: %s\nВес: %.2f кг.\nРост: %.2f м.", p.Name, p.Weight, p.Height)
}

// Print выводит на печать персональные данные пользователя
func (p Personal) Print() {
	fmt.Println(p)
}
