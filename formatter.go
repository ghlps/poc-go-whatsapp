package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func fmtEmoji(meal Meal) string {
	emojis := ""
	for _, icon := range meal.Icons {
		if emoji, exists := iconsMap[icon]; exists {
			emojis += " " + emoji
		}
	}
	return meal.Name + emojis
}

func fmtMeal(meals []Meal) string {
	var formatted []string
	for _, meal := range meals {
		formatted = append(formatted, fmtEmoji(meal))
	}
	return strings.Join(formatted, "\n")
}

func fmtMenu(evt EventLambda) string {
	var sections []string
	layout := "02/01/2006"
	menu := evt.ResponsePayload

	log.Printf("Formatting menu: %+v", menu)

	t, err := time.Parse(layout, menu.Date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	weekday := t.Weekday()
	titleStr := fmt.Sprintf(title, strings.ToUpper(menu.Restaurant.Name), weekDay[int(weekday)], menu.Date)

	header := titleStr
	if evt.RunType == "CHECKUP" {
		header = modWarning + "\n\n" + titleStr
	}

	for i, mealServed := range menu.Served {
		if meals, exists := menu.Meals[mealServed]; exists {
			mealSection := mealsHeaders[i] + "\n" + fmtMeal(meals)
			if i == 0 {
				sections = append(sections, header+"\n\n"+mealSection)
			} else {
				sections = append(sections, mealSection)
			}
		}
	}

	sections = append(sections, legend, fmt.Sprintf(taken, menu.Restaurant.Url), fmt.Sprintf(channel, os.Getenv("TARGET_NUMBER")), copyright)
	return strings.Join(sections, "\n\n")

}

var iconsMap = map[string]string{
	"Simbolo-vegano-300x300":               "🌱",
	"Origem-animal-site":                   "🥩",
	"Gluten-site":                          "🌾",
	"Leite-e-derivados-site":               "🥛",
	"Ovo-site":                             "🍳",
	"Alergenicos-site":                     "⚠️",
	"Simbolo-pimenta-300x300":              "🌶️",
	"vegan":                                "🌱",
	"carne":                                "🥩",
	"gluten":                               "🌾",
	"lactose":                              "🥛",
	"ovo":                                  "🍳",
	"476bf979-2cbb-476b-8739-02ed26485235": "🐷",
	"daf04cd5-bacd-4ea4-91ce-48ea45cb0ac4": "⚠️",
	"Indicado para veganos":                "🌱",
	"Contém pimenta":                       "🌶️",
	"Contêm produtos de origem suína":      "🐷",
	"Contém ingrediente(s) potencialmente alergênico(s)": "⚠️",
	"Contêm produtos de origem animal":                   "🥩",
	"Contém glúten":                                      "🌾",
	"Contém leite e/ou derivados":                        "🥛",
	"Contêm ovos":                                        "🍳",
}

var weekDay = []string{
	"DOMINGO",
	"SEGUNDA",
	"TERÇA",
	"QUARTA",
	"QUINTA",
	"SEXTA",
	"SÁBADO",
}

var legend = `🌱 - Indicado para veganos
🥩 - Contém produtos de origem animal
🐷 - Contém produtos de origem suína
🌾 - Não indicado para celíacos por conter glúten
🥛 - Não indicado para intolerantes à lactose por conter lactose
🍳 - Contêm ovos
⚠️ - Contém produto(s) alergênico(s)
🍯 - Contém mel
🌶️ - Contém pimenta`

var mealsHeaders = []string{"*CAFÉ DA MANHÃ*", "*ALMOÇO*", "*JANTAR*"}

var title = `*CARDÁPIO RU %s - %s - %s*`

var modWarning = `Cardápio atualizado`

var taken = `Cardápio retirado de forma automatizada do site oficial do restaurante universitário disponível no link %s`

var channel = `Essa mensagem foi enviada ao canal de WhatsApp disponível no link %s`

var copyright = "Esta mensagem e este canal não possuem relação com a universidade nem com o restaurante universitário"
