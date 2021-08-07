package models

type Medicamento struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	Nombre    string `json:"nombre"`
	Precio    string `json:"precio"`
	Ubicacion string `json:"ubicacion"`
}
