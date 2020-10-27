package sample

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/lifenglsf/grpc_demo/hw/pb"
)

func newKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}
func newCPU() *pb.CPU {
	brand := randomeCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat(2.0, 3.5)
	maxGhz := randomFloat(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

func newGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat(1.0, 1.5)
	maxGhz := randomFloat(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}
func newRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}
	return ram
}
func newSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(120, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}
	return ssd
}

func newHDD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randomInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}
	return ssd
}
func newScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randomFloat32(123, 17),
		Resolution: randomScreenResolution(),
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
	return screen
}
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:       randomID(),
		Name:     name,
		Brand:    brand,
		Cpu:      newCPU(),
		Gpus:     []*pb.GPU{newGPU()},
		Storages: []*pb.Storage{newSSD(), newHDD()},
		Screen:   newScreen(),
		Keyboard: newKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randomFloat(1.0, 3.0),
		},
		PriceUsd:    randomFloat(1500, 3000),
		ReleaseYear: uint32(randomInt(2015, 2020)),
		UpdatedAt:   ptypes.TimestampNow(),
	}
	return laptop
}
