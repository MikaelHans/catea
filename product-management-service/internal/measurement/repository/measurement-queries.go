package repository

import (
	"database/sql"
	"github.com/MikaelHans/catea/product-management-service/pkg/structs"
	"github.com/MikaelHans/catea/product-management-service/pkg/util"
)

func Create_Measurements(mData structs.Measurements_Data)(*sql.Rows, error){
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `INSERT INTO public.measurement(measurement_name) VALUES ($1);`
	rows, err := db.Query(query, mData.Name)
	db.Close()
	return rows, err
}

func Get_Measurements()(*sql.Rows, error){
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `SELECT * FROM measurements;`
	rows, err := db.Query(query)
	db.Close()
	return rows, err
}

func Get_Measurements_By_Name(data structs.Measurements_Data)(*sql.Rows, error){
	db, err := util.GetDBConnection()

	if err != nil {
		return nil, err
	}

	var query string = `SELECT * FROM measurements WHERE measurement_name = $1;`
	rows, err := db.Query(query, data.Name)
	db.Close()
	return rows, err
}