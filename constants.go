package main

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
