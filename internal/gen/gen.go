package gen

import (
	"github.com/gocolly/colly"
	"html/template"
	"io"
	"strings"
)

type Items []Item

type Item struct {
	Name string
	ID   string
}

func NewItem(name string, ID string) *Item {
	return &Item{Name: name, ID: ID}
}

func GetItems() Items {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("minecraft-ids.grahamedgecombe.com"),
	)

	var items Items
	// On every a element which has href attribute call callback
	c.OnHTML("td.row-desc", func(e *colly.HTMLElement) {
		if e.Attr("class") == "row-desc" {
			txt := strings.ReplaceAll(e.Text, ")", "")
			txt = strings.ReplaceAll(txt, " ", "")
			txt = strings.ReplaceAll(txt, "-", "")
			txt = strings.ReplaceAll(txt, "'", "")
			s := strings.Split(txt, "(")
			if !strings.ContainsAny(s[0], "0123456789") && strings.Contains(s[1], "minecraft:") {
				item := NewItem(s[0], s[1])
				items = append(items, *item)
			}
		}
	})
	// Start scraping on https://minecraft-ids.grahamedgecombe.com
	_ = c.Visit("https://minecraft-ids.grahamedgecombe.com/")
	c.Wait()

	return items
}

func (i *Items) GenerateCode(w io.Writer) error {
	tpl, err := template.New("items.tmpl").ParseFiles("templates/items.tmpl")
	if err != nil {
		return err
	}

	return tpl.Execute(w, i)
}
