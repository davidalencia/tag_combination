package notion

import (
	"fmt"

	"github.com/buger/jsonparser"
	"github.com/gofiber/fiber/v2"
	"github.com/lucasb-eyer/go-colorful"
)


func notionGet(endpoint string) []byte {
	request := fiber.Get("https://api.notion.com/v1/" + endpoint)
	// request.Debug()
	// request := fiber.Get("https://api.notion.com/v1/databases/62d01a2d65284e2f847809abfe3b88da")

	// to set headers
	request.Set("Authorization", `Bearer secret_g0dFRGqvrY1KbPYGXPmMstUF27Vpc752Z7xst8eV8Fq`)
	request.Set("Notion-Version", `2022-06-28`)

	_, data, err := request.Bytes()
	if err != nil {
		panic(err)
	}
	return data
}

func notionPost(endpoint string) []byte {
	request := fiber.Post("https://api.notion.com/v1/" + endpoint)
	// request.Debug()
	// request := fiber.Get("https://api.notion.com/v1/databases/62d01a2d65284e2f847809abfe3b88da")

	// to set headers
	request.Set("Authorization", `Bearer secret_g0dFRGqvrY1KbPYGXPmMstUF27Vpc752Z7xst8eV8Fq`)
	request.Set("Notion-Version", `2022-06-28`)

	_, data, err := request.Bytes()
	if err != nil {
		panic(err)
	}
	return data
}

func  RGB() (r, g, b uint8) {
	r = uint8(1)
	g = uint8(2)
	b = uint8(3)
	return
}

func UpdateDatabase(page_id string){
	colors_mapping := map[string]string{
		"default": "#9B9A97",
		"gray": "#9B9A97",
		"brown": "#64473A",
		"orange": "#D9730D",
		"yellow": "#DFAB01",
		"green": "#0F7B6C",
		"blue": "#0B6E99",
		"purple": "#6940A5",
		"pink": "#AD1A72",
		"red": "#E03E3E",
	}

	db := notionPost("databases/62d01a2d65284e2f847809abfe3b88da/query")
	jsonparser.ArrayEach(db, func(result []byte, dataType jsonparser.ValueType, offset int, err error) {
		colors := []string{}
		// id, _ := jsonparser.GetString(result, "id")
		jsonparser.ArrayEach(result, func(tag []byte, dataType jsonparser.ValueType, offset int, err error) {
			color, _ := jsonparser.GetString(tag, "color")
			colors = append(colors, color)
		}, "properties", "Tags", "multi_select")

		colorfulColors := make([]colorful.Color, len(colors))
		for k, v := range colors {
			colorfulColors[k], _ = colorful.Hex(colors_mapping[v])
		}
		fmt.Println(mixColors(colorfulColors))
	},"results")
}

func mixColors (colorArray []colorful.Color) colorful.Color{
	if len(colorArray) == 0 {
		mixedColor, _ := colorful.Hex("#ffffff")
		return mixedColor
	}
	mixedColor := colorArray[0]
	for i := 1; i < len(colorArray); i++ {
		mixedColor = mixedColor.BlendHsv(colorArray[i], 0.5)
	}
	return mixedColor
}