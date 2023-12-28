package types

type TestData struct {
	test string `json:"test"`
}

type TestRes struct {
	Code int       `json:"code"`
	Msg  string    `json:"msg"`
	Data *TestData `json:"data,omitempty"` // 使用指针类型表示可选字段，并使用omitempty忽略空值
}
