/*
* File Name:	type17_memory_device.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-08-19
 */
package godmi

import (
	"fmt"
	"strconv"
)

type MemoryDeviceFormFactor byte

func (m MemoryDeviceFormFactor) String() string {
	factors := [...]string{
		"Other",
		"Unknown",
		"SIMM",
		"SIP",
		"Chip",
		"DIP",
		"ZIP",
		"Proprietary Card",
		"DIMM",
		"TSOP",
		"Row of chips",
		"RIMM",
		"SODIMM",
		"SRIMM",
		"FB-DIMM",
	}
	return factors[m-1]
}

type MemoryDeviceType byte

func (m MemoryDeviceType) String() string {
	types := [...]string{
		"Other",
		"Unknown",
		"DRAM",
		"EDRAM",
		"VRAM",
		"SRAM",
		"RAM",
		"ROM",
		"FLASH",
		"EEPROM",
		"FEPROM",
		"EPROM",
		"CDRAM",
		"3DRAM",
		"SDRAM",
		"SGRAM",
		"RDRAM",
		"DDR",
		"DDR2",
		"DDR2 FB-DIMM",
		"Reserved",
		"Reserved",
		"Reserved",
		"DDR3",
		"FBD2",
		"DDR4",
		"LPDDR",
		"LPDDR2",
		"LPDDR3",
		"LPDDR4",
	}
	return types[m-1]
}

type MemoryDeviceTypeDetail uint16

func (m MemoryDeviceTypeDetail) String() string {
	//details := [...]string{
	//	"Reserved",
	//	"Other",
	//	"Unknown",
	//	"Fast-paged",
	//	"Static column",
	//	"Pseudo-static",
	//	"RAMBUS",
	//	"Synchronous",
	//	"CMOS",
	//	"EDO",
	//	"Window DRAM",
	//	"Cache DRAM",
	//	"Non-volatile",
	//	"Registered (Buffered)",
	//	"Unbuffered (Unregistered)",
	//	"LRDIMM",
	//}
	var ret string
	if CheckBit(uint64(m), 1) {
		ret += "Other"
	}
	if CheckBit(uint64(m), 2) {
		ret += "Unknown"
	}
	if CheckBit(uint64(m), 3) {
		ret += "Fast-paged"
	}
	if CheckBit(uint64(m), 4) {
		ret += "Static column"
	}
	if CheckBit(uint64(m), 5) {
		ret += "otheeudo-static"
	}
	if CheckBit(uint64(m), 6) {
		ret += "RAMBUS"
	}
	if CheckBit(uint64(m), 7) {
		ret += "Synchronous"
	}
	if CheckBit(uint64(m), 8) {
		ret += "CMOS"
	}
	if CheckBit(uint64(m), 9) {
		ret += "EDO"
	}
	if CheckBit(uint64(m), 10) {
		ret += "Window DRAM"
	}
	if CheckBit(uint64(m), 11) {
		ret += "Cache DRAM"
	}
	if CheckBit(uint64(m), 12) {
		ret += "Non-volatile"
	}
	if CheckBit(uint64(m), 13) {
		ret += "Registered (Buffered)"
	}
	if CheckBit(uint64(m), 14) {
		ret += "Unbuffered (Unregistered)"
	}
	if CheckBit(uint64(m), 15) {
		ret += "LRDIMM"
	}
	return ret
}

type MemorySizeType uint16

type MemoryDeviceSetType byte

func (s MemoryDeviceSetType) String() string {
	if s == 0x00 {
		return "None"
	} else if s == 0xff {
		return "Unknown"
	}
	return strconv.Itoa(int(s))
}

type MemorySpeedType uint16

func (s MemorySpeedType) String() string {
	if s == 0 {
		return "Unknown"
	} else if s == 0xffff {
		return "Reserved"
	}
	return strconv.Itoa(int(s))
}

type MemoryDevice struct {
	infoCommon
	PhysicalMemoryArrayHandle  uint16
	ErrorInformationHandle     uint16
	TotalWidth                 uint16
	DataWidth                  uint16
	Size                       uint16
	FormFactor                 MemoryDeviceFormFactor
	DeviceSet                  MemoryDeviceSetType
	DeviceLocator              string
	BankLocator                string
	Type                       MemoryDeviceType
	TypeDetail                 MemoryDeviceTypeDetail
	Speed                      MemorySpeedType
	Manufacturer               string
	SerialNumber               string
	AssetTag                   string
	PartNumber                 string
	Attributes                 byte
	ExtendedSize               uint32
	ConfiguredMemoryClockSpeed uint16
	MinimumVoltage             uint16
	MaximumVoltage             uint16
	ConfiguredVoltage          uint16
}

func (m MemoryDevice) String() string {
	return fmt.Sprintf("Memory Device\n"+
		"\tPhysical Memory Array Handle: %0#X\n"+
		"\tMEMORY ERROR INFORMATION HANDLE: %0#X\n"+
		"\tTotal Width: %dbits\n"+
		"\tData Width: %dbits\n"+
		"\tSize: %d\n"+
		"\tForm Factor: %s\n"+
		"\tDevice Set: %s\n"+
		"\tDevice Locator: %s\n"+
		"\tBank Locator: %s\n"+
		"\tMemory Type: %s\n"+
		"\tType Detail: %s\n"+
		"\tSpeed: %s MHz\n"+
		"\tManufacturer: %s\n"+
		"\tSerial Number: %s\n"+
		"\tAsset Tag: %s\n"+
		"\tPart Number: %s\n"+
		"\tAttributes: rank %d\n"+
		"\tExtended Size: %d MB\n"+
		"\tConfigured Memory Clock Speed: %d MHz\n"+
		"\tMinimum voltage: %d\n"+
		"\tMaximum voltage: %d\n"+
		"\tConfigured voltage: %d ",
		m.PhysicalMemoryArrayHandle,
		m.ErrorInformationHandle,
		m.TotalWidth,
		m.DataWidth,
		m.Size,
		m.FormFactor,
		m.DeviceSet,
		m.DeviceLocator,
		m.BankLocator,
		m.Type,
		m.TypeDetail,
		m.Speed,
		m.Manufacturer,
		m.SerialNumber,
		m.AssetTag,
		m.PartNumber,
		m.Attributes,
		m.ExtendedSize,
		m.ConfiguredMemoryClockSpeed,
		m.MinimumVoltage,
		m.MaximumVoltage,
		m.ConfiguredVoltage,
	)
}

func newMemoryDevice(h dmiHeader) dmiTyper {
	data := h.data
	m := &MemoryDevice{
		PhysicalMemoryArrayHandle:  u16(data[0x04:0x06]),
		ErrorInformationHandle:     u16(data[0x06:0x08]),
		TotalWidth:                 u16(data[0x08:0x0A]),
		DataWidth:                  u16(data[0x0A:0x0C]),
		Size:                       u16(data[0x0C:0x0e]),
		FormFactor:                 MemoryDeviceFormFactor(data[0x0E]),
		DeviceSet:                  MemoryDeviceSetType(data[0x0F]),
		DeviceLocator:              h.FieldString(int(data[0x10])),
		BankLocator:                h.FieldString(int(data[0x11])),
		Type:                       MemoryDeviceType(data[0x12]),
		TypeDetail:                 MemoryDeviceTypeDetail(u16(data[0x13:0x15])),
		Speed:                      MemorySpeedType(u16(data[0x15:0x17])),
		Manufacturer:               h.FieldString(int(data[0x17])),
		SerialNumber:               h.FieldString(int(data[0x18])),
		AssetTag:                   h.FieldString(int(data[0x19])),
		PartNumber:                 h.FieldString(int(data[0x1A])),
		Attributes:                 data[0x1B],
		ExtendedSize:               u32(data[0x1C:0x20]),
		ConfiguredMemoryClockSpeed: u16(data[0x20:0x22]),
		MinimumVoltage:             u16(data[0x22:0x24]),
		MaximumVoltage:             u16(data[0x24:0x26]),
		ConfiguredVoltage:          u16(data[0x26:0x28]),
	}
	MemoryDevices = append(MemoryDevices, m)
	return m
}

var MemoryDevices []*MemoryDevice

func GetMemoryDevice() string {
	var ret string
	for i, v := range MemoryDevices {
		ret += "\nMemoryDevices index:" + strconv.Itoa(i) + "\n" + v.String()
	}
	return ret
}

func init() {
	addTypeFunc(SMBIOSStructureTypeMemoryDevice, newMemoryDevice)
}
