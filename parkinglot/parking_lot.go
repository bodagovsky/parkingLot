package parkinglot


type ParkingLot struct {
	TotalSpace uint32
	AllocatedLots map[uint32]*Lot
}

type Lot struct {
	Plate string
	Colour string
}


func (p *ParkingLot) GetSlotByPlate(plate string) uint32 {

	for i:= 1; i <= int(p.TotalSpace); i++ {
		if car, ok := p.AllocatedLots[uint32(i)]; ok && car.Plate == plate{
			return uint32(i)
		}
	}

	return 0
}

func (p *ParkingLot) GetPlatesByColor(color string) []string {
	var plates []string

	for i:= 1; i <= int(p.TotalSpace); i++ {
		if car, ok := p.AllocatedLots[uint32(i)]; ok && car.Colour == color{
			plates = append(plates, car.Plate)
		}
	}
	if len(plates) == 0 {
		plates = append(plates, "No")
	}
	return plates
}


func (p *ParkingLot) GetSlotByColor(color string) []uint32 {
	var cars []uint32

	for i:= 1; i <= int(p.TotalSpace); i++ {
		if pColor, ok := p.AllocatedLots[uint32(i)]; ok && pColor.Colour == color{
			cars = append(cars,uint32(i))
		}
	}
	if len(cars) == 0 {
		cars = append(cars, 0)
	}
	return cars
}

func (p *ParkingLot) LeaveSlot(slot uint32 ) error  {
	if _, ok:= p.AllocatedLots[slot]; !ok {
		return SlotIsEmptyErr(SlotIsEmptyError)
	}
	delete(p.AllocatedLots, slot)
	return nil
}

func (p *ParkingLot) AllocSlot(plate string, color string) uint32{
	emptySlot := p.findEmptySlots()[0]

	if emptySlot != 0 {
		newLot := &Lot{
			Plate:plate,
			Colour:color,
		}
		p.AllocatedLots[emptySlot] = newLot
		return emptySlot
	}

	return 0
}

func (p *ParkingLot) findEmptySlots() []uint32 {
	var slots []uint32

	for i:= 1; i <= int(p.TotalSpace); i++ {
		if _, ok := p.AllocatedLots[uint32(i)]; !ok {
			slots = append(slots, uint32(i))
		}
	}
	if len(slots) == 0 {
		slots = append(slots, 0)
	}
	return slots
}