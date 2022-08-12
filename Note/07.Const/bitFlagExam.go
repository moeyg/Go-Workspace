package main

import "fmt"

const (
	MainRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	Storage
)

// Bit Flag ON
func SetLight(rooms, room uint8) uint8 {
	// 0000 0000 | 0000 0001(MainRoom)
	// 0000 0000 | 0000 0100(BathRoom)
	// 0000 0000 | 0000 1000(Storage)
	// ∴ 0000 0000 | 0000 1101(MainRoom + BathRoom + Storage)
	return rooms | room // 0000 1101
}

// Bit Flag OFF
func ResetLight(rooms, room uint8) uint8 {
	// 0000 1101(rooms)
	// ^0000 1000(room = Storage) -> 1111 0111(^Storage)
	// 0000 1101 & 1111 0111
	return rooms &^ room // 0000 0101
}

// Bit Flag Check
func IsLightOn(rooms, room uint8) bool {
	// 0000 0101 & 0000 0001(MainRoom) => 0000 0001 == 0000 0001 (true)
	// 0000 0101 & 0000 0100(BathRoom) => 0000 0100 == 0000 0100 (true)
	return rooms & room == room
}

// 0000 0101
func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MainRoom) {
		fmt.Println("Turn on the MainRoom")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("Turn on the LivingRoom")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("Turn on the BathRoom")
	}
	if IsLightOn(rooms, Storage) {
		fmt.Println("Turn on the Storage")
	}
}

func main() {
	var rooms uint8
	fmt.Printf("MainRoom : %08b\n", MainRoom) // MainRoom : 0000 0001
	fmt.Printf("LivingRoom : %08b\n", LivingRoom) // LivingRoom : 0000 0010
	fmt.Printf("BathRoom : %08b\n", BathRoom) // BathRoom : 0000 0100
	fmt.Printf("Storage : %08b\n", Storage) // Storage : 0000 1000

	// Bit Flag ON
	rooms = SetLight(rooms, MainRoom)
	rooms = SetLight(rooms, BathRoom)
	rooms = SetLight(rooms, Storage)
	// MainRoom + BathRoom + Storage
	fmt.Printf("%08b\n", rooms) // 0000 1101

	// Storage Bit Flag OFF
	rooms = ResetLight(rooms, Storage)
	fmt.Printf("%08b\n", rooms) // 0000 0101

	// Bit Flag ON 부분만 출력
	TurnLights(rooms)
}