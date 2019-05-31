# Value receiver vs Pointer receiver in Golang

Suatu dilema ketika  mendefinisikan metode struct.

Khususnya ketika memutuskan reciver di method, haruskah Kita menggunakan Pointer receiver atau Value receiver?

`func (t *Type) Method() {} //pointer receiver`

Vs.

`func (t Type) Method() {} //value receiver`

# Kapan kita harus menggunakan pointer receiver

Jika ingin mengubah nilai yang diterima dalam method, memanipulasi dan sebagainya kita gunakan Pointer receiver

# Kapan kita harus menggunakan value receiver
Jika Kita tidak perlu mengedit nilai gunakan value receiver.

Untuk pemakaian seperti apa bisa dilihat di Project simple-order-golang
[Disini](https://gitlab.warungpintar.co/enrinal/intern-diary/tree/master/Simple-Order-System-Management-Golang)
