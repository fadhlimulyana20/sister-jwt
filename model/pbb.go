package model

import "go.mongodb.org/mongo-driver/bson/primitive"

var TablePBB string = "pbb"

type DataPBB struct {
	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Tahun int                `bson:"tahun" json:"tahun"`
	Pajak int                `bson:"pajak" json:"pajak"`
	Denda int                `bson:"denda" json:"denda"`
}

type Pbb struct {
	Id        primitive.ObjectID `bson:"_id" json:"id"`
	UserID    primitive.ObjectID `bson:"user_id" json:"user_id"`
	NOPPBB    string             `bson:"noppbb" json:"noppbb"`
	Nama      string             `bson:"nama" json:"nama"`
	Alamat    string             `bson:"alamat" json:"alamat"`
	Kabupaten string             `bson:"kabupaten" json:"kabupaten"`
	Kecamatan string             `bson:"kecamatan" json:"kecamatan"`
	Desa      string             `bson:"desa" json:"desa"`
	Rt        string             `bson:"rt" json:"rt"`
	Rw        string             `bson:"rw" json:"rw"`
	DataPbb   []DataPBB          `bson:"data_pbb" json:"data_pbb"`
}
