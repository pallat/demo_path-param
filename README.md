# Demo API with path parameters

โดยปกติการรับ path parameters ในภาษา Go จะไม่มี function ใน standard library มาช่วย
แต่เราสามารถจัดการเองได้ ด้วยการเอา URI มาใช้เทคนิคการตัดคำ
หรือใช้ตัวช่วยจากภายนอก ซึ่งในที่นี้เราจะใช้ gorilla/mux เข้ามาช่วย

## ขั้นตอนการสร้างโปรเจค

1. go mod init {ใส่ชื่อโปรเจค เช่น `github.com/pallat/demoapi` _หมายเหตุ_ **pallat** คือชื่อ account ของผู้สอน}
2. เขียน code ตามตัวอย่างในไฟล์นี้
3. เนื่องจากเราใช้ library จากภายนอก จึงจะต้องนำเข้ามาด้วยคำสั่ง `go get github.com/gorilla/mux` หรือใช้คำสั่ง `go mod tidy` เลยก็ได้
4. รันโปรแกรมด้วยคำสั่ง `go run main.go`
5. ทดสอบด้วย test.http หรือเปิด web browser ไปที่ `http://localhost:8080/todos/100`

**เลข 100 สามารถทดลองเปลี่ยนเป็นอะไรก็ได้**

## คำอธิบายเพิ่มเติม

```go
func main() {
    r := mux.NewRouter()
    r.HandleFunc("/todos/{id}", todoHandler)
    log.Fatal(http.ListenAndServe(":8080", r))
}
```

สังเกตที่ parameter ตัวที่ 2 ของ `http.ListenAndServe(":8080", r)` ที่เราใส่ต้วแปร `r` เข้าไป
ก่อนหน้านี้เราเคยเขียน API โดยไม่ต้องใช้ gorilla/mux มาช่วย เราจะเห็นการใช้ค่า `nil` แทนที่จะเป็น `r`
เนื่องจาก ความจริงค่านี้ เราจะต้องใส่ router ลงไป โดย router นี้เป็น interface ชื่อว่า Handler
แต่ก่อนนี้เราใส่ค่า nil ลงไปได้ เพราะว่าตัว standard library จะสร้าง default router เอาไว้ล่วงหน้า
เมื่อใส่ค่า nil ลงไป มันจะไปใช้ default router มาทำงานให้ทันที

แต่เมื่อเราใช้ gorilla/mux มาสร้าง router ให้ เราจะจำเป็นต้องใส่ `r` ซึ่งสร้างโดย gorilla/mux ลงไป
เพื่อให้ตัว ListenAndServe รู้ว่าจะไม่ต้องไปใช้ default router แต่ให้ใช้ตัวนี้แทนได้เลย

```go
vars := mux.Vars(r)
id := vars["id"]
```

function `mux.Vars(r)` สังเกตุว่า มันรับเอาตัว *http.Request เข้าไป เนื่องจากในนั้นมีค่า request URI
มันแค่ต้องการเอา URI ไปตัดคำที่เราระบุตำแหน่งไว้ใน `/todos/{id}` มาสร้างเป็น key/value และคืนค่าออกมา
ให้เราเป็น `map[string]string`
เวลานำไปใช้ เราก็แค่ระบุ key ให้ตรงกับที่ระบุใน path-param ซึ่งในที่นี้คือ `id` นั่นเอง
