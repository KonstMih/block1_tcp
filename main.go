package main

import (
	"block_2/devices"
	"block_2/modbus"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "192.168.127.254:951")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	f, err := os.OpenFile("data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	for {

		var data = make(map[string]interface{})

		t, d := modbus.Today()

		data["T_1time"] = t
		data["Day"] = d

		for _, i := range devices.Devices {

			if _, err = conn.Write(i.Request); err != nil {
				fmt.Println(err)
				return
			}

			buff := make([]byte, 100)

			t := time.NewTimer(1000 * time.Millisecond)
			<-t.C

			conn.SetReadDeadline(time.Now().Add(time.Millisecond * 1000))

			_, err := conn.Read(buff)
			if err != nil {
				fmt.Println(err)
				continue
			}

			var temp []byte
			var mass_temp []float32

			switch i.Name {
			case ("tm5104d"):
				temp = modbus.Rtu_data_tm5104d(buff, i.Qan_chanels)
				mass_temp = modbus.Rtu_temp(temp)

			case ("trm138"):
				temp = modbus.Rtu_data_trm138(buff, i.Qan_chanels)
				mass_temp = modbus.Rtu_temp(temp)

			case ("irt5920n"):
				temp = modbus.Ascii_data_itr5920n(buff, i.Id)
				mass_temp = modbus.Ascii_temp(temp)

			case ("f1772"):
				temp = modbus.Rtu_data_f1772(buff, i.Qan_chanels)
				mass_temp = modbus.Rtu_temp(temp)

			}

			fmt.Println(mass_temp) // убрать при окончании настройки

			for j := 0; j < int(i.Qan_chanels); j++ {
				if mass_temp[j] == float32(math.Inf(1)) {
					data[i.Chanels[j]] = "Обрыв"
					continue
				}
				data[i.Chanels[j]] = mass_temp[j]
			}

		}

		data_js, _ := json.Marshal(data)

		f.WriteString(string(data_js) + "\n")

	}

}
