package fizzbuzz // ใน go ต้องประกาศ package เสมอ

import "testing" // import lib testing ของ go มาเขียน test ได้เลย

// function test ของ go ต้องขึ้นด้วย Test เสมอ
// และตามด้วย A-Z เท่านั้น หรือใช้ _ ได้ และใน parameter ให้ชี้ pointer ไปที่ *testing.T เสมอ
func Test_Input_1_Should_be_return_1(t *testing.T) {
	// นี่คือการ new class เข้ามา เพื่อส่งค่าไปที่ file fizzbuzz
	// function fizzbuzz โดยส่ง 1 ไป เก็บไว้ในตัวแปร anwser
	anwser := fizzbuzz(1)
	// นำค่าที่ถูก retrun กลับมา มาเช็คค่าที่เรา expect ไว้ ในนี้จะต่างจาก ใน ภาษาอื่นๆ
	// เพราะเราจะต้องเขียน เช็คคำตอบเอง ไม่มี assert ให้ใช้แต่อย่างใด
	// ในที่นี้จะเช็ค anwser != "0" จริงหรือไ่ม่
	if anwser != "1" {
		// ถ้าไม่ใช่ ก็ให้ใช้ ตัวแปร t ที่ชี้ pointer *testing.T ทำการสร้าง error ขึ้นมา
		// จากนั้นก็เขียน error เองซะเลย
		t.Error("Expected 1 Actual ", anwser)
	}
	// ต่อไปนี้ก็เรื่อมเขียน คล้ายๆกัน ไป
}

func Test_Input_2_Should_be_return_2(t *testing.T){
	anwser := fizzbuzz(2)
	if anwser != "2" {
		t.Error("Expected 2 Actual ", anwser)
	}
}

func Test_Input_3_Should_be_return_Fizz(t *testing.T){
	anwser := fizzbuzz(3)
	if anwser != "Fizz" {
		t.Error("Expected Fizz Actual ", anwser)
	}
}

func Test_Input_5_Should_be_return_Buzz(t *testing.T){
	anwser := fizzbuzz(5)
	if anwser != "Buzz" {
		t.Error("Expected Buzz Actual ", anwser)
	}
}

func Test_Input_6_Should_be_return_Fizz(t *testing.T){
	anwser := fizzbuzz(6)
	if anwser != "Fizz" {
		t.Error("Expected Fizz Actual ", anwser)
	}
}

func Test_Input_9_Should_be_return_Fizz(t *testing.T){
	anwser := fizzbuzz(9)
	if anwser != "Fizz" {
		t.Error("Expected Fizz Actual ", anwser)
	}
}

func Test_Input_10_Should_be_return_Buzz(t *testing.T){
	anwser := fizzbuzz(10)
	if anwser != "Buzz" {
		t.Error("Expected Buzz Actual ", anwser)
	}
}

func Test_Input_15_Should_be_return_FizzBuzz(t *testing.T){
	anwser := fizzbuzz(15)
	if anwser != "FizzBuzz" {
		t.Error("Expected FizzBuzz Actual ", anwser)
	}
}