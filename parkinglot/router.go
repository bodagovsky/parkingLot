package parkinglot

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Create = `create_parking_lot`
	Park = `park`
	Leave = `leave`
	Status = `status`
	GetPlatesByColor = `registration_numbers_for_cars_with_colour`
	GetSlotsByColor = `slot_numbers_for_cars_with_colour`
	GetSlotByPlate = `slot_number_for_registration_number`
	Exit = `exit`
	Help = `help`
	)
var pLot = &ParkingLot{TotalSpace: 0,
	AllocatedLots:make(map[uint32]*Lot)}


func Process(input string) (string,error) {
	cmd := strings.Split(input, " ")

	switch cmd[0] {

	case Create:
		switch len(cmd) {
		case 2:
			num, err := strconv.Atoi(cmd[1])
			if err != nil{
				return fmt.Sprintf("Invalid input: %v.Try another one", num), err
			} else if num <= 0{
				return "", errors.New("undesirable input")
			}
			pLot.TotalSpace = uint32(num)
			return fmt.Sprintf("Successfully created a parkinglot with %v slots", num), nil
		default:
			err := errors.New("wrong command")
			return "", err
		}

	case Park:
		switch len(cmd) {
		case 3:
			plate, color := cmd[1], cmd[2]

			num := pLot.AllocSlot(plate, color)

			if num != 0 {
				return fmt.Sprintf("Successfully allocated slot number %v", num), nil
			}

			return "", errors.New("sorry, parking lot is full")

		default:
			err := errors.New("wrong command")
			return "", err
		}

	case Leave:
		switch len(cmd) {
		case 2:
			num, err := strconv.Atoi(cmd[1])

			if err != nil || num <=0{
				return "", err
			}
			delete(pLot.AllocatedLots, uint32(num))
			return fmt.Sprintf("Slot number %v is free", num), nil

		default:
			err := errors.New("wrong command")
			return "", err
		}

	case Status:
		switch len(cmd) {
		case 1:
			header := "Slot No.   Registration No   Color\n"
			for key, value := range pLot.AllocatedLots {
				header += fmt.Sprintf("%v  %v  %v\n", key, value.Plate, value.Colour)
			}
			emptySlots := pLot.findEmptySlots()
			switch len(emptySlots) {
			case 1:
				msg := fmt.Sprintf("Empty slot is %v\n", emptySlots[0]) + header
				return msg, nil
			default:
				msg := "Empty slots are "
				for _, item := range emptySlots {
					msg += fmt.Sprintf("%v, ", item)
				}
				return fmt.Sprintf("%v\n%v",msg[:len(msg)-2], header) , nil
			}

		default:
			return "", errors.New("wrong command")
		}

	case GetPlatesByColor:
		switch len(cmd) {
		case 2:
			color := cmd[1]
			msg1 := fmt.Sprintf("Plates with color %v are ", color)
			msg2 := fmt.Sprintf("Plate with color %v is ", color)
			slots := pLot.GetPlatesByColor(color)
			if slots[0] != "No" {
				switch len(slots) {
				case 1:
					msg2 = msg2 + slots[0]
					return msg2, nil
				default:
					for _, slot := range slots {
						msg1 = msg1 + fmt.Sprintf("%v, ", slot)
					}
					return msg1[:len(msg1)-2], nil
				}

			}
			return "", errors.New(fmt.Sprintf("no cars with color %v", color))

		default:
			return "", errors.New("wrong input for getting plate by color")
	}

	case GetSlotsByColor:
		switch len(cmd) {
		case 2:
			color := cmd[1]
			msg1 := fmt.Sprintf("Slots with color %v are ", color)
			msg2 := fmt.Sprintf("Slot with color %v is ", color)
			slots := pLot.GetSlotByColor(color)
			if slots[0] != 0 {
				switch len(slots) {
				case 1:
					msg2 = msg2 + fmt.Sprintf("%v", slots[0])
					return msg2, nil
				default:
					for _, slot := range slots {
						msg1 = msg1 + fmt.Sprintf("%v, ", slot)
					}
					return msg1, nil
				}
			}
			return "", errors.New(fmt.Sprintf("no slots with color %v", color))

		default:
			return "", errors.New("error trying to get slot by color")
		}

	case GetSlotByPlate:
		switch len(cmd) {
		case 2:
			plate := cmd[1]
			msg := fmt.Sprintf("Slot with plate %v is ", plate)
			slot := pLot.GetSlotByPlate(plate)
			if slot != 0 {
				msg = msg + fmt.Sprintf("%v", slot)
				return msg, nil
			}

			return "", errors.New(fmt.Sprintf("no slots with plate %v", plate))

		default:
			return "", errors.New("error trying to get slot by plate")
		}
	case Help:
		helpMsg := fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n", Create,
			Park, Leave, Status, GetPlatesByColor, GetSlotsByColor,GetSlotByPlate, Exit)
		return helpMsg, nil

	case Exit:
		os.Exit(1)
	default:
		return "unknown command", nil
	}
	return "", errors.New("unknown error")
}