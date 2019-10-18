package GoFolder

import (
	"fmt"
	"time"
	"os"
	"net"
	"io"	
)

type ByteStruct struct {
	asd byte
	height byte
}

const ENQ byte = 05
const ACK byte = 06
const STX byte = 02
const ETX byte = 03
const EOT byte = 04
const NAK byte = 21
const DLE byte = 16

const Ip_And_Port strign = "127.0.0.1:4545"
const type_connect string = "tcp"

//
//// timeouts
//
const T1 int16 = 500
const T2 int16 = 2000
const T3 int16 = 500
const T4 int16 = 500
const T5 int16 = 10000
const T6 int16 = 500
const T7 int16 = 500
const T8 int16 = 1000

var _map_machine = map[string]int{
	"АТОЛ 11Ф": 48,
	"АТОЛ 15Ф": 48,	 
	"АТОЛ 20Ф": 62,
	"АТОЛ FPrint-22ПТК": 62,
	"АТОЛ 25Ф": 57,
	"АТОЛ 30Ф": 48,
	"АТОЛ 50Ф": 50,
	"АТОЛ 52Ф": 66,
	"АТОЛ 55Ф": 50,
	"АТОЛ 60Ф": 48,
	"АТОЛ 77Ф": 57,
	"АТОЛ 90Ф": 48,
	"АТОЛ 42ФС": 62,
	"АТОЛ 91Ф": 46,
	"АТОЛ 92Ф": 46,
	"АТОЛ 150Ф": 50,
   }




// return false this is not signal
//if return 0 is error
//if return 1 is ACK
//if return 2 is ENQ(RC>1)
func ActiveSide(MainData) int {
	var FRC int32 = 0
	var RC int32 = 0
	for FRC := 1; ; FRC++ {
		RS = 0;
		for ; RC == 100 ; RS++ {		
			if RC <= 5 {

				// Send ENQ
				strB := GET_BYTES_WITH_CRC(EOT)
				SendDataNet(type_connect, Ip_And_Port, strB)

				// listen byte with command
				var listenData byte
				listenData, er = ListenAnswer(type_connect, Ip_And_Port, T1)
				switch listenData {

				// NAK
				case NAK:
					time.Sleep(T1)
					continue

				// not byte
				case nil:
					continue

				// !(ACK||ENQ||NAK)
				case !(ACK||ENQ||NAK):
					if IsFRC(FRC) {
						RC = 0
						FRC = FRC + 1
						continue
					} else {
						sendEOT()
						return 0
					}
				// ENQ
				case ENQ:
					time.Sleep(T7)
					if IsFRC(FRC) {
						RC = 0
						FRC = FRC + 1
						continue
					} else {
						sendEOT()
						return 0
					}
					

				// ACK
				case ACK:
					break
				}
			} else {
				sendEOT()
			}
		}
	}
	for RC = 1; ; RC++ {
		if RC <= N {
			strB := GET_BYTES_WITH_CRC(MainData)
			SendDataNet(type_connect, Ip_And_Port, strB)
			
			listenData, er := ListenAnswer(type_connect, Ip_And_Port, T3)		
			
			switch listenData[0] {

			//not byte
			case nil:
				continue
			// (!ACK)||(ENQ&&RC==1)
			case (!ACK)||(ENQ && RC==1):
				continue

			//ACK
			case ACK:
				//send EOT
				sendEOT()
				return 1

			//ENQ&&(RC>1)
			case ENQ&&(RC>1):
				return 2
				//
				
			}
		} else {
			sendEOT()
		}
	}
}


func ActiveReceiver(command int) int {
	var RC int = 0
	if command == 1 {
		for RC = 1; ; i++ {	
		
			if RC <= N1 {
				
				listenData, er := ListenAnswer(type_connect, Ip_And_Port, T5)

				switch listenData[0] {
				case nil:
					fmt.Println("Нет связи")	
					return 0
				case !(ENQ):
					continue
				case ENQ:
					ActiveReceiver1()
				}
			}
		}
	}
		
}
const N int32 = 10
const N1 int32 = 100
func IsFRC(int32 FRC) bool {
	if FRC <= N1 {
		return true
	}	
	return false
}







func ActiveReceiver1()  {
	var byteG byte
	step1:
		//send ACK
		strB := GET_BYTES_WITH_CRC(ACK)
		SendDataNet(type_connect, Ip_And_Port, strB)
		step2:
		for RC := 1; RC < count; RC++ {
			if RC <= N1 {
				//get bite
				listenData, er := ListenAnswer(type_connect, Ip_And_Port, T2)		
				switch listenData[0] {
				case nil:
					fmt.Println("Нет связи")		
				case STX:
					break
				case ENQ:
					FRC++
					if FRC <= N {
						goto step1
					} else {
						fmt.Println("Нет связи")		
					}
				case !(STX||ENQ):
					continue					
				}
			} else {
				fmt.Println("Нет связи")	
			}
		}

		step3:
		//clear buffer
		BC := 0
		DLE_Flag := 0

		step4:
		if BC <= BMax {
			//get bite
			listenData, er := ListenAnswer(type_connect, Ip_And_Port, T2)		
			switch listenData[0] {
			case nil:
				if FRC <= N {
					goto step2
				} else {
					fmt.Println("Нет связи")	
				}
			case !(nil):
				if DLE_Flag ==1 {
					DLE_Flag = 0
					//записать байт в буфер
					BC++
					goto step4
				} else {
					if listenData[0] == DLE {
						DLE_Flag = 1
						//записать байт в буфер
						BC++
						goto step4
					} else {
						if listenData[0] != ETX {
							//записать байт в буфер
							BC++
							goto step4
						}
					}
				}
			}
		} else {
			FRC++
			if FRC <= N {
				goto strp2
			} else {
				fmt.Println("Нет связи")	
			}			
		}

		
		//get bite
		listenData, er := ListenAnswer(type_connect, Ip_And_Port, T2)	










		if byteG == ENQ && FRC <= N {
			continue
		} else {
			fmt.Println("Нет связи")	
		}
		

	
}





func AR_StepOne()  {
	
}

func sendEOT()  {
	strB := GET_BYTES_WITH_CRC(EOT)
	SendDataNet(type_connect, Ip_And_Port, strB)
	fmt.Println("Нет связи")	
}

func IsErrorData(data1 []byte) (bool) {	
	var STX_Pos int
	var ETX_Pos int	
	for i := 0; i < len(data1); i++ {
		if data1 == STX {
			STX_Pos = i
		}
		if data1 == ETX {
			ETX_Pos = i
		}
	}

	var crc byte = data1[ETX_Pos:ETX_Pos+1]

	var one_xor_byte byte
	var data2 []byte = data1[STX_Pos:ETX_Pos]
	one_xor_byte = data2[0]
	for i := 1; i < len(byteACat2); i++ {
		one_xor_byte = one_xor_byte ^ data2[i]
	}
	if one_xor_byte == crc {
		return true
	}
	return false
}


func GET_BYTES_WITH_CRC(arrayByte []byte) ([]byte) {		
	var LvalInd []int
	LvalInd = append(LvalInd, 0)
	var byteACat1 []byte
	for i := 0; i < len(arrayByte); i++ {			
		fmt.Println(i, " : ", arrayByte[i] , " : ", len(arrayByte))	
		if arrayByte[i] == DLE || arrayByte[i] == ETX {									
			byteACat1 = append(byteACat1, DLE) 
			byteACat1 = append(byteACat1, arrayByte[i])									
		} else {
			byteACat1 = append(byteACat1, arrayByte[i])
		}
		
	}

	byteACat1 = append(byteACat1, ETX)

	var one_xor_byte byte
	one_xor_byte = byteACat1[0]
	for i := 1; i < len(byteACat1); i++ {
		one_xor_byte = one_xor_byte ^ byteACat1[i]
	}

	var byteACat2 []byte
	byteACat2 = append(byteACat2, STX)
	byteACat2 = append(byteACat2, byteACat1...)

	byteACat2 = append(byteACat2, one_xor_byte)
	
	return byteACat2
}


func SendDataNet(typeC strign, IpAndPort strign, source []byte) (bool) {
	conn, err := net.Dial(typeC, IpAndPort) 
    if err != nil { 
        fmt.Println(err) 
        return false
	}
	defer conn.Close()
	
	// отправляем сообщение серверу
	if n, err := conn.Write([]byte(source));
	n == 0 || err != nil { 
		fmt.Println(err) 
		return false
	}	
	defer conn.Close()
	return true
}


func ListenAnswer(typeC strign, IpAndPort strign, tS int) ([]byte, bool) {
	var buff []byte
	conn, err := net.Dial(typeC, IpAndPort) 
	if err != nil { 
        fmt.Println(err) 
        return buff, false
	}
	defer conn.Close()

	
	conn.SetReadDeadline(time.Now().Add(time.Second * tS))
        for{
            buff = make([]byte, 1024)
            n, err := conn.Read(buff)
            if err != nil { break}
            fmt.Print(string(buff[0:n]))
            conn.SetReadDeadline(time.Now().Add(time.Millisecond * 700))
		}
		
	defer conn.Close()

	return buff, true
}
func IsCountByte(string name_machine, ArrayByte *byte) bool {
	val_map, ok_map := _map_machine[name_machine]
	if !ok_map {
		return false
	}
	if val >= len(*ArrayByte) {
		return true
	} 
	return false;	
}
func commandAndResponseBlock(str string) ([]byte) {
	var arrayByte []byte
	arrayByte[0] = STX;
	var BytesStr []byte = []byte(str)
	arrayByte = append(arrayByte, BytesStr) 
	arrayByte[len(arrayByte)+1] = ETX

}