package entity

type Client struct {
	TypeDocument string `bson:"type_document,omitempty" json:"type_document,omitempty"`
	Name         string `bson:"name,omitempty" json:"name,omitempty"`
	Dir          string `bson:"dir,omitempty" json:"dir,omitempty"`
	Ubigeo       string `bson:"ubigeo,omitempty" json:"ubigeo,omitempty"`
	Ruc          string `bson:"ruc,omitempty" json:"ruc,omitempty"`
	Email        string `bson:"email,omitempty" json:"email,omitempty"`
}
