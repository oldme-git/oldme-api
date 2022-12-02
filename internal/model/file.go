package model

type FileInfo struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"path" dc:"文件URL"`
}
