package main

import "fmt"

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	lastNotifications := make(map[int]int) // Хранение последнего оповещения для каждого пользователя
	globalNotifications := 0               // Глобальное количество оповещений

	for i := 0; i < q; i++ {
		var t, id int
		fmt.Scan(&t, &id)

		switch t {
		case 1:
			if id == 0 { // Глобальное оповещение
				globalNotifications++
			} else { // Персональное оповещение
				lastNotifications[id] = globalNotifications // Обновляем последнее оповещение для пользователя id
			}
		case 2:
			if id == 0 { // Запрос глобального оповещения
				fmt.Println(globalNotifications)
			} else { // Запрос номера последнего оповещения для пользователя id
				if last, ok := lastNotifications[id]; ok {
					fmt.Println(globalNotifications - last) // Выводим количество оповещений, полученных пользователем id
				} else {
					fmt.Println(0) // Если пользователь еще не получал оповещений
				}
			}
		}
	}
}
