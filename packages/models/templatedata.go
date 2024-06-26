package models

// TemplateData holds data sent to the templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSFRToken string
	Falsh     string
	Warning   string
	Error     string
}