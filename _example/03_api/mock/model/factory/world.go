// Code generated by eevee. DO NOT EDIT!

package factory

import (
	"api/entity"
	"api/model"
)

func DefaultWorld() *model.World {
	value := &model.World{World: &entity.World{
		ID:   0,
		Name: "",
	}}
	return value
}

func DefaultWorlds() *model.Worlds {
	values := &model.Worlds{}
	{
		value := &model.World{World: &entity.World{
			ID:   0,
			Name: "",
		}}
		values.Add(value)
	}
	return values
}