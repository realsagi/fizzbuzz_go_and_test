package fizzbuzz // ต้องใส่ package  เสมอ

import "strconv" // import ของที่ต้องการให้ convert ค่าระหว่าง int ไปเป็น string

// การใช่ parameter ให้ใช้ชื่อตัวแปรไว้ด้านหน้า และ type ไว้ด้านหลัง
// หลังจากชื่อ ฟังก์ชั่น แล้วให้ ใส่ type อยาก return กลับไปในที่นี้คือ  string
func fizzbuzz(number int) string{
	// strconv.Itoa(number) คือการแปลงค่าจาก int ไปเป็น string
	// แล้วเก็บค่าไว้ในตัวแปรชื่อ  anwser
	anwser := strconv.Itoa(number)

	// การใช้ if ก็คล้ายๆกับภาษาอื่น
	if number % 15 == 0{
		return "FizzBuzz"
	}
	if number % 3 == 0 {
		return "Fizz"
	}
	if number % 5 == 0{
		return "Buzz"
	}
	return anwser
}