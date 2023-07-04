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
	"strconv"
	"io"
)

func main() {

	conn_kettle, err := net.Dial("tcp", "192.168.127.254:951")  // порты 951-котёл, 950-турбина 
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn_kettle.Close()

	conn_turbo, err := net.Dial("tcp", "192.168.127.254:950")  // порты 951-котёл, 950-турбина 
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn_kettle.Close()

	var key_week bool = false;

	
	for {

		data_file, err := os.OpenFile("/home/asutp/code/goprogect/block1_tcp/data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
		}
		defer data_file.Close()

		today := time.Now()
	    weekday := today.Weekday()

	    //int(today.Minute()) == 28 {
	    if int(weekday) == 6 {

	    	fmt.Println("Запсиь")

			if key_week == true {
				_, week := today.ISOWeek()
				weekStr := strconv.Itoa(week)

				data_file.Close()

				new, err := os.Create("/home/asutp/code/goprogect/block1_tcp/" + weekStr + ".json")
	    		if err != nil {
	        		fmt.Println(err)
	            }
	            
	            original, err := os.Open("/home/asutp/code/goprogect/block1_tcp/data.json")
                if err != nil {
                    fmt.Println(err)
                }


	            bytesCopy, err := io.Copy(new, original)
	            if err != nil {
	                fmt.Println(err)
				}

                new.Close()
				original.Close()

				fmt.Printf("Bytes Written: %d\n", bytesCopy)

				os.Remove("/home/asutp/code/goprogect/block1_tcp/data.json")

				data_file, err := os.OpenFile("/home/asutp/code/goprogect/block1_tcp/data.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
				if err != nil {
			        fmt.Println(err)
		        }
		        defer data_file.Close()

				key_week = false
			
		    }   
		} else {key_week = true}

	    for {

			var data = make(map[string]interface{})

			t, d := modbus.Today()

			data["Time"] = t
			data["Day"] = d

			for _, i := range devices.Devices_kettle {

				if _, err = conn_kettle.Write(i.Request); err != nil {
					fmt.Println(err)
					return
				}

				buff := make([]byte, 100)

				t := time.NewTimer(120 * time.Millisecond)
				<-t.C

				conn_kettle.SetReadDeadline(time.Now().Add(time.Millisecond * 120))

				_, err := conn_kettle.Read(buff)
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

				case ("kp1m"):
					temp = modbus.Rtu_data_kp1m(buff)
					mass_temp = modbus.RtuAscii_temp(temp)
				}

				//fmt.Println(mass_temp) // убрать при окончании настройки

				for j := 0; j < int(i.Qan_chanels); j++ {
					if mass_temp[j] == float32(math.Inf(1)) {
						data[i.Chanels[j]] = "Обрыв"
						continue
					}
					data[i.Chanels[j]] = mass_temp[j]
				}

			}

			for _, i := range devices.Devices_turbo {

				if _, err = conn_turbo.Write(i.Request); err != nil {
					fmt.Println(err)
					return
				}

				buff := make([]byte, 100)

				t := time.NewTimer(120 * time.Millisecond)
				<-t.C

				conn_turbo.SetReadDeadline(time.Now().Add(time.Millisecond * 120))

				_, err := conn_turbo.Read(buff)
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

			data_file.WriteString(string(data_js) + "\n")

            today := time.Now()
	        weekday := today.Weekday()

	        //int(today.Minute()) == 28 {
	        if int(weekday) == 6 {
	        	break
	        }		

	}

    }
}

