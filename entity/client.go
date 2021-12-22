package entity

type Client struct {
	Id       string `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string `bson:"name,omitempty" json:"name,omitempty"`
	LastName string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Ruc      string `bson:"ruc,omitempty" json:"ruc,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
}
