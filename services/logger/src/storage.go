/*
storage.go
Implements a simple SQLite Storage Interface

###################################################################################

MIT License

Copyright (c) 2020 Bruno Hautzenberger

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StorageInterface interface {
	SaveData(aX float32, aY float32, aZ float32, gX float32, gY float32, gZ float32) error
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

func (s *SQLiteStorage) SaveData(aX float32, aY float32, aZ float32, gX float32, gY float32, gZ float32) error {
	db, err := s.connect()
	if err != nil {
		return err
	}

	dm := DataModel{}
	dm.AX = aX
	dm.AY = aY
	dm.AZ = aZ
	dm.GX = gX
	dm.GY = gY
	dm.GZ = gZ

	tx := db.Save(&dm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

type DataModel struct {
	gorm.Model
	AX float32
	AY float32
	AZ float32
	GX float32
	GY float32
	GZ float32
}
