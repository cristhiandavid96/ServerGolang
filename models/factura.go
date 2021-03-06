package models

type Factura struct {
	ID            int64  `json:"id" gorm:"primary_key;auto_increment"`
	PagoTotal    string  `json:"pago_total"`
	FechaCrear   string    `json:"fecha_crear"`
}


// TableName overrides the table name used by User to `profiles`
func (Factura) TableName() string {
  return "factura"
}