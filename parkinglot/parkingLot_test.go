package parkinglot

import (
	"testing"
)



func TestAllocSlot(t *testing.T) {
	lot1 := &Lot{
		Plate:"H1-31-MM",
		Colour:"White",
	}
	lot2 := &Lot{
		Plate:"GY-1F-U6",
		Colour:"Black",
	}
	lot3 := &Lot{
		Plate:"P1-R4-LP",
		Colour:"Blue",
	}


	parkingLot := ParkingLot{
		TotalSpace:5,
		AllocatedLots: map[uint32]*Lot{
			1:lot1,
			2:lot2,
			3:lot3,
		},
	}
	slot := parkingLot.AllocSlot("123-HTR-876", "white")
	if slot == 0 {
		t.Error("Problem with allocating new slot while having available slots")
	}

	parkingLot.AllocatedLots[5] = &Lot{
		Plate:"TR-Y5-HU",
		Colour:"Ocean",
	}

	slot = parkingLot.AllocSlot("77-HJ-09", "Red")

	if slot != 0 {
		t.Error("Problem with overloading parkinglot lot")
	}
}

func TestLeavingSlot(t *testing.T) {
	parkingLot := ParkingLot{
		TotalSpace:5,
	}


	lot1 := &Lot{
		Plate:"H1-31-MM",
		Colour:"White",
	}
	lot2 := &Lot{
		Plate:"GY-1F-U6",
		Colour:"Black",
	}
	lot3 := &Lot{
		Plate:"P1-R4-LP",
		Colour:"Blue",
	}

	alocslots:= map[uint32]*Lot{
		1:lot1,
		2:lot2,
		3:lot3,
	}

	parkingLot.AllocatedLots = alocslots

	_ = parkingLot.LeaveSlot(2)

	emptyLot := parkingLot.findEmptySlots()[0]

	switch emptyLot {
	case 0:
		t.Error("failed to find an empty slot")

	case 2:
		return

	default:
		t.Error("wrong place for parking is found")
	}

}

func TestLeavingEmptySlot(t *testing.T) {
	parkingLot := ParkingLot{
		TotalSpace:5,
	}

	err := parkingLot.LeaveSlot(6)

	if err != nil {
		switch err.Error() {
		case SlotIsEmptyError:
			return
		default:
			t.Error("wrong behavior while leaving empty slot")
		}
	}

}

func TestGetSlotByColor (t *testing.T) {
	lot1 := &Lot{
		Plate:"H1-31-MM",
		Colour:"White",
	}
	lot2 := &Lot{
		Plate:"GY-1F-U6",
		Colour:"Black",
	}
	lot3 := &Lot{
		Plate:"P1-R4-LP",
		Colour:"Blue",
	}


	parkingLot := ParkingLot{
		TotalSpace:5,
		AllocatedLots: map[uint32]*Lot{
			1:lot1,
			2:lot2,
			3:lot3,
		},
	}

	blueSlot := parkingLot.GetSlotByColor("Blue")

	blackSlot := parkingLot.GetSlotByColor("Black")

	greenSlot := parkingLot.GetSlotByColor("Green")

	if blueSlot[0] != 3 || blackSlot[0] != 2 || greenSlot[0] != 0 {
		t.Error("problem with getting slot by color")
	}

}

func TestGetPlatesByColor (t *testing.T) {
	lot1 := &Lot{
		Plate:"H1-31-MM",
		Colour:"White",
	}
	lot2 := &Lot{
		Plate:"GY-1F-U6",
		Colour:"Black",
	}
	lot3 := &Lot{
		Plate:"P1-R4-LP",
		Colour:"Blue",
	}


	parkingLot := ParkingLot{
		TotalSpace:5,
		AllocatedLots: map[uint32]*Lot{
			1:lot1,
			2:lot2,
			3:lot3,
		},
	}

	bluePlates := parkingLot.GetPlatesByColor("Blue")

	blackPlates := parkingLot.GetPlatesByColor("Black")

	greenPlates := parkingLot.GetPlatesByColor("Green")

	if bluePlates[0] != "P1-R4-LP" || blackPlates[0] != "GY-1F-U6" || greenPlates[0] != "No" {
		t.Error("problem with getting plates by color")
	}
}

func TestGetSlotByPlate(t *testing.T) {
	lot1 := &Lot{
		Plate:"H1-31-MM",
		Colour:"White",
	}
	lot2 := &Lot{
		Plate:"GY-1F-U6",
		Colour:"Black",
	}
	lot3 := &Lot{
		Plate:"P1-R4-LP",
		Colour:"Blue",
	}


	parkingLot := ParkingLot{
		TotalSpace:5,
		AllocatedLots: map[uint32]*Lot{
			1:lot1,
			2:lot2,
			3:lot3,
		},
	}

	slot1 := parkingLot.GetSlotByPlate("H1-31-MM")
	slot2 := parkingLot.GetSlotByPlate("P1-R4-LP")
	slot3 := parkingLot.GetSlotByPlate("NotExistingPlateNumber")

	if slot1 != 1 || slot2 != 3 || slot3 != 0 {
		t.Error("problem with getting slot by plate number")
	}
}



