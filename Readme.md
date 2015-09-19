# FizzBuzz
## เนื้อหา
* การติดตั้ง go ให้ไปดูที่ [link](https://golang.org/doc/install)
* วิธีการเขียนภาษา go เบื้องต้น นิดหน่อย
* การเขียน test ในภาษา go เบื้องต้น นิดหน่อย
* วิธีการเขียนในแบบต่างๆ (ตอนนี้ขอมีท่า เดียวก่อน)

### note syntax ประกาศ function ของ go
	เมื่องต้องการประกาศ ฟังก์ชั่นต่างๆ เช่น
	function sum(){
	}
	การเขียนคือให้ เขียน { ต่อท้ายของ function ไปเลย
	ห้ามขึ้นบรรทัดใหม่มิเช่นนั้นมันจะ error
	
	ตัวอย่างการเขียนแล้วเกิด error 
	function sum()
	{
	}

การ run test
เข้าไปใน folder ที่เป็น ไฟล์ก่อนจากนั้น run ด้วยคำสั่ง
>$go test

### test coverate
	test coverage คือ การทดสอบว่า code โปรแกรมของเราที่เป็น
	เป็น code production นั้นมีการเขียน test ไว้แล้วหรือยัง
	มันสามารถบอกออกเป็นเป็น % ได้ด้วย ตัวอย่างการแสดงผมดังนี้
	PASS
	coverage: 100.0% of statements
	ok_/Users/apple/Documents/yeddo_team/fizzbuzz/	method1	0.008s

ติดตั้ง test cover ก่อนด้วยคำสั่ง
>$go get golang.org/x/tools/cmd/cover

run คำสั่ง test cover เพื่อตรวจสอบ code coverage
>$go test -cover