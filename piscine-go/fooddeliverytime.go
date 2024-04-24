package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var order_ food

	if order == "burger" {
		order_.preptime = 15
		return order_.preptime
	} else if order == "chips" {
		order_.preptime = 10
		return order_.preptime
	} else if order == "nuggets" {
		order_.preptime = 12
		return order_.preptime
	}
	return 404
}
