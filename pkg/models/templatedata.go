package models

// TemplateData is a  type to handle any type of template data rendered to templates
type TemplateData struct {
	StringMap map[string]string
	NumberMap map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} // store any other type of data
	CSRFToken string                 // security token for forms

	// messages for the users
	Flash   string
	Warning string
	Error   string
}
