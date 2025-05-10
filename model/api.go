package model

type Api struct {
	LunarModel
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
	Group  string `json:"group"`
}
