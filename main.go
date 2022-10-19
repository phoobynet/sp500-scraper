package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
)

type SP500Element struct {
	Symbol   string
	Security string
	CIK      string
}

func (s *SP500Element) String() string {
	return fmt.Sprintf("%s %s %s\n", s.Symbol, s.Security, s.CIK)
}

func Get() ([]SP500Element, error) {
	c := colly.NewCollector()

	elements := make([]SP500Element, 0)

	// Find and visit all links
	c.OnHTML("table.wikitable tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			symbol := el.ChildText("td:nth-child(1)")

			if symbol != "" {
				ticker := SP500Element{
					Symbol:   symbol,
					Security: el.ChildText("td:nth-child(2)"),
					CIK:      el.ChildText("td:nth-child(9)"),
				}
				elements = append(elements, ticker)
			}
		})
	})

	err := c.Visit("https://en.wikipedia.org/wiki/List_of_S%26P_500_companies")
	if err != nil {
		return nil, err
	}

	return elements, nil
}
