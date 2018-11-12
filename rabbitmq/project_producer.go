package rabbitmq

import (
	"encoding/json"

	"github.com/orest/briefer/analyzers"
)

// func InitProjectProducer(id string, title string, text string) ProjectProducer {
// return ProjectProducer{id: id, title: title, description: text}
// }

type ProjectProducer struct {
	Id          string
	Title       string
	Description string
}

func (p *ProjectProducer) Push() {
	client := InitClient()
	analyzer := analyzers.TextAnalyzer{}
	brief := analyzer.Brief(p.Title, p.Description)
	message := map[string]string{
		"id":    p.Id,
		"brief": brief,
	}
	json, _ := json.Marshal(message)
	client.PublishAsync("briefer.projects", string(json))
}
