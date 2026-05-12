package helper

import "inventory-indra/model"

func GetStatusStock(stock int) model.StatusStock {
	if stock == 0 {
		return model.StatusStockSoldOut
	} 

	if stock < 10 {
		return model.StatusStockLow
	}

	return model.StatusStockSafe
}