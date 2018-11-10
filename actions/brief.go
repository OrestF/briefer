package actions

import (
	"strings"

	"github.com/gobuffalo/buffalo"
	"github.com/orest/briefer/analyzers"
)

// BriefHandler is a default handler to serve up
// a home page.
func BriefHandler(c buffalo.Context) error {
	text := c.Param("text")
	anal := analyzers.TextAnalyzer{}
	brief := anal.Brief(text)
	return c.Render(200, r.JSON(map[string]string{"original": text, "brief": strings.Join(brief, ", ")}))
}
