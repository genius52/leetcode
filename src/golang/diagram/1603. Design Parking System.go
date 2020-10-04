package diagram

//carType is 1, 2, or 3

type CarType int32
const (
	BIG_SLOT 	CarType = 1
	MEDIUM_SLOT CarType = 2
	SMALL_SLOT	CarType = 3
)

type ParkingSystem struct {
	slots [3]int
}

func Constructor1603(big int, medium int, small int) ParkingSystem {
	var obj ParkingSystem
	obj.slots[BIG_SLOT - 1] = big
	obj.slots[MEDIUM_SLOT - 1] = medium
	obj.slots[SMALL_SLOT - 1] = small
	return obj
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.slots[carType - 1] <= 0{
		return false
	}
	this.slots[carType - 1]--
	return true
}

