package entity

type RUC struct {
	Name           string `json:"nombre,omitempty" bson:"nombre,omitempty"`
	TypeDocument   string `json:"tipoDocumento,omitempty" bson:"tipoDocumento,omitempty"`
	DocumentNumber string `json:"numeroDocumento,omitempty" bson:"numeroDocumento,omitempty"`
	State          string `json:"estado,omitempty" bson:"estado,omitempty"`
	Condition      string `json:"condicion,omitempty" bson:"condicion,omitempty"`
	Direction      string `json:"direccion,omitempty" bson:"direccion,omitempty"`
	Ubigeo         string `json:"ubigeo,omitempty" bson:"ubigeo,omitempty"`
	ViaType        string `json:"viaTipo,omitempty" bson:"viaTipo,omitempty"`
	ViaName        string `json:"viaNombre,omitempty" bson:"viaNombre,omitempty"`
	District       string `json:"distrito,omitempty" bson:"distrito,omitempty"`
	Province       string `json:"provincia,omitempty" bson:"provincia,omitempty"`
	Departament    string `json:"departamento,omitempty" bson:"departamento,omitempty"`
}
