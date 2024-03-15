package template

type template struct {
	Name      string      `json:"name"`
	Version   float64     `json:"version"`
	Input     []Input     `json:"input"`
	Tickets   []Ticket    `json:"tickets"`
	OffDelpsm interface{} `json:"delpsm"`
}

type Input struct {
	Name string `json:"name"`
	On   string `json:"on"`  // 打开
	Off  string `json:"off"` //关闭状态的字符串，没有是空
}

type Ticket struct {
	Code     int64    `json:"code"`
	Type     string   `json:"type"` //tcc,eeconf_v2,feature_gating
	Name     string   `json:"name"`
	Template string   `json:"template"`
	Positive bool     `json:"template"` //控制fg的开关，如果是true，就是和功能开关一致，否则就是反的
	Psm      []string `json: psm`       // eeconf工单影响的服务eeconf修改了需要重新发布
}
