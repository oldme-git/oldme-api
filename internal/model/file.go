package model

type FileInfo struct {
	Name string `json:"name" dc:"文件名称"`
	Url  string `json:"path" dc:"文件URL"`
	Dir  string `json:"-" dc:"文件DIR"` // dir不能暴露到外部，他包含着服务器的目录结构
}
