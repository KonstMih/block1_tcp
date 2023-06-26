package modbus

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"time"
)

func Ieee_754(slovo int32) float64 {

	var s int32
	var m int32
	var m1 float64

	if (slovo >> 31) == 0 {
		s = 1
	} else {
		s = -1
	}

	e := (slovo >> 23) & 0xff

	if e != 0 {
		m = ((slovo & 0x7fffff) | 0x800000)
	} else {
		m = (slovo & 0x7fffff) << 1
	}

	m1 = float64(m) * math.Pow(2, -23)

	return float64(s) * m1 * math.Pow(2, float64(e)-127)
}

//---------------------------------------------------------------------------------------------

func Rtu_data_f1772(response []byte, data_len int32) []byte {
	response = response[3 : data_len*4+3]
	for i := 0; i < int(data_len); i++ {
		response = append(response[:], response[2:4]...)
		response = append(response[4:], response[0:2]...)
	}
	return response

}

func Rtu_data_tm5104d(response []byte, data_len int32) []byte {
	response = response[3 : data_len*4+3]
	return response

}

func Rtu_data_trm138(response []byte, data_len int32) []byte {
	response = response[3 : data_len*10+3]
	for i := 0; i < int(data_len); i++ {
		response = append(response[:4*i], response[4*i+6:]...)
	}
	return response
}

func Ascii_data_itr5920n(response []byte, id int) []byte {
	if id < 10 {
		response = response[4:10]
	} else {
		response = response[5:11]
	}
	return response
}

func Rtu_data_kp1m(response []byte) []byte {

    response = response[7:29]		
	
	response[8], 	response[9], response[10], response[11], response[12], response[13] = 
	response[4], response[5], response[2], response[3], response[0], response[1]
	
	response = response[6:]
	
	ascii_code := map[byte]byte{48:0x0, 49:0x1, 50:0x2, 51:0x3, 52:0x4, 53:0x5, 54:0x6, 55:0x7, 
	56:0x8, 57:0x9, 65:0xA, 66:0xB, 67:0xC, 68:0xD, 69:0xE, 70:0xF}
    
    response[8], response[9], response[10], response[11], response[12], response[13], response[14], response[15] = 
    ascii_code[response[0]], ascii_code[response[1]], ascii_code[response[2]], ascii_code[response[3]], ascii_code[response[4]], 
    ascii_code[response[5]], ascii_code[response[6]], ascii_code[response[7]]
	
	response = response[8:]    

	return response
}

//---------------------------------------------------------------------------------------------

func Ascii_temp(response []byte) []float32 {

	var temp float64
	var mass_temp []float32

	temp, _ = strconv.ParseFloat(string(response), 32)
	mass_temp = append(mass_temp, float32(temp))

	return mass_temp

}

func Rtu_temp(response []byte) []float32 {

	var mass_temp []float32

	for i := 0; i < len(response); i += 4 {
		var m []byte = response[i : i+4]
		slovo := int32(binary.BigEndian.Uint32(m))
		var temp float64 = Ieee_754(int32(slovo))
		mass_temp = append(mass_temp, float32(temp))

	}
	return mass_temp
}

func RtuAscii_temp(response []byte) []float32 {

	var mass_temp []float32

	var temp uint32 = (uint32(response[0]) << 28)|(uint32(response[1]) << 24)|
	(uint32(response[2]) << 20)|(uint32(response[3]) << 16)|(uint32(response[4]) << 12)|
	(uint32(response[5]) << 8)|(uint32(response[6]) << 4)|(uint32(response[7]))

	var value float64 = Ieee_754(int32(temp))
	mass_temp = append(mass_temp, float32(value))

	return mass_temp
}

//-----------------------------------------------------------------------------------

func Today() (string, string) {
	today := time.Now()
	hours, min, sec := today.Clock()
	year, month, day := today.Date()

	t := fmt.Sprintf("%d:%d:%d", hours, min, sec)
	d := fmt.Sprintf("%d:%d:%d", year, int(month), day)

	return t, d
}
