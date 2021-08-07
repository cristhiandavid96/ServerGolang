package models

type Promocion struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Descripcion    string `json:"descripcion"`
	Porcentaje     string `json:"porcentaje"`
	FechaInicio    string `json:"fecha_inicio"`
	FechaFin       string `json:"fecha_fin"`
}


// TableName overrides the table name used by User to `profiles`
func (Promocion) TableName() string {
  return "promocion"
}