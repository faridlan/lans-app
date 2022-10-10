package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Rekap struct {
	Id          primitive.ObjectID `bson:"_id, omitempty" json:"id,omitempty"`
	CsName      string             `bson:"cs_name,omitempty" json:"cs_name,omitempty"`
	CusName     string             `bson:"cus_name,omitempty" json:"cus_name,omitempty"`
	RekapStatus bool               `bson:"rekap_status,omitempty" json:"rekap_status,omitempty"`
	PrintStatus bool               `bson:"print_status,omitempty" json:"print_status,omitempty"`
	RekapDate   int64              `bson:"rekap_date,omitempty" json:"rekap_date,omitempty"`
}
