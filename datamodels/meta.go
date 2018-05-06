package datamodels

type Meta struct {
	ColumnName string `json:"column_name"`
	DataType   string `json:"data_type"`
}

func (Meta) TableName() string {
	return "information_scheme.columns"
}
