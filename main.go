package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Доставка еды Киров")
	w.Resize(fyne.NewSize(200, 400))

	title := widget.NewLabel("Оформление заказа")

	name_label := widget.NewLabel("Ваше имя:")
	name := widget.NewEntry()

	food_label := widget.NewLabel("Выберите еду для заказа:")
	food := widget.NewCheckGroup([]string{"Пицца", "Роллы", "Шаурма", "Бургеры", "Напитки"}, func(s []string) {})

	male_label := widget.NewLabel("Ваш пол:")
	male := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(s string) {})

	adress_label := widget.NewLabel("Введите адрес доставки:")
	adress := widget.NewEntry()

	result := widget.NewLabel("")

	button := widget.NewButton("ОФОРМИТЬ", func() {
		username := name.Text
		food_arr := food.Selected
		male_val := male.Selected
		adress_dos := adress.Text

		result.SetText(username + "\n" + male_val + "\n" + adress_dos + "\n")

		for _, i := range food_arr {
			result.SetText(result.Text + i + "\n")
		}
	})

	w.SetContent(container.NewVBox(
		title,
		name_label,
		name,
		food_label,
		food,
		male_label,
		male,
		adress_label,
		adress,
		button,
		result,
	))

	w.ShowAndRun()
}
