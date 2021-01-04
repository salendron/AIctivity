package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StorageInterface interface {
	SaveData(data *Data) error
}

type SQLiteStorage struct {
	SQLiteDBFile string
}

func (s *SQLiteStorage) Initialize(sqliteDBFile string) {
	s.SQLiteDBFile = sqliteDBFile
}

func (s *SQLiteStorage) connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(s.SQLiteDBFile), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&DataModel{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (s *SQLiteStorage) SaveData(data *Data) error {
	db, err := s.connect()
	if err != nil {
		return err
	}
	dataModel := data.ToDataModel()
	tx := db.Save(&dataModel)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

type DataModel struct {
	gorm.Model
	AX   int
	AY   int
	AZ   int
	GX   int
	GY   int
	GZ   int
	Temp float32
}

func (d Data) ToDataModel() DataModel {
	dm := DataModel{}
	dm.AX = *d.AX
	dm.AY = *d.AY
	dm.AZ = *d.AZ
	dm.GX = *d.GX
	dm.GY = *d.GY
	dm.GZ = *d.GZ
	dm.Temp = *d.Temp

	return dm
}
