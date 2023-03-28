package entity

import "github.com/hotel/app/entity/idgenerator"

func getNextID() int64 {
	return idgenerator.GetNextID()
}
