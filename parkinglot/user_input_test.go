package parkinglot

import ("testing")


func TestUserInputCreate(t *testing.T) {
	input1 := Create + " " + "6"
	input2 := Create + " " + "0"
	input3 := Create
	input4 := Create + " " +  "-1"
	input5 := Create + " " + "hello there"


	_, err := Process(input1)

	if err != nil {
		t.Error(err.Error())
	}

	_, err = Process(input2)

	if err == nil {
		t.Error("parking lot with 0 slots")
	}

	_, err = Process(input3)

	if err == nil {
		t.Error("wrong input accepted")
	}

	_, err = Process(input4)

	if err == nil {
		t.Error("parking lot with -1 slots created")
	}

	_, err = Process(input5)

	if err == nil {
		t.Error("non integer input accepted")
	}

}

func TestUserInputPark(t *testing.T) {
	pLot.TotalSpace = 4
	input1 := Park
	input2 := Park + " " + "PlateNumber" + " " + "Color"
	input3 := Park + " " + "Color"

	_, err := Process(input1)

	if err == nil {
		t.Error("input with no plate number and color")
	}

	_, err = Process(input2)

	if err != nil {
		t.Error(err.Error())
	}

	_, err = Process(input3)

	if err == nil {
		t.Error("not enough arguments")
	}

}

func TestUserInputPlotsByColor(t *testing.T) {
	pLot.TotalSpace = 1
	pLot.AllocatedLots[1] = &Lot{
		Plate:"HH-90-PL",
		Colour:"White",
	}
	input := GetPlatesByColor + " " + "White"
	_, err := Process(input)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUserInputSlotsByColor(t *testing.T) {
	pLot.TotalSpace = 1
	pLot.AllocatedLots[1] = &Lot{
		Plate:"HH-90-PL",
		Colour:"White",
	}
	input := GetSlotsByColor + " " + "White"

	_, err:= Process(input)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestUserInputSlotByPlate (t *testing.T) {
	pLot.TotalSpace = 1
	pLot.AllocatedLots[1] = &Lot{
		Plate:"HH-90-PL",
		Colour:"White",
	}
	input := GetSlotByPlate + " " + "HH-90-PL"
	_, err:= Process(input)

	if err != nil {
		t.Error(err.Error())
	}
}