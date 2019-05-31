# Clean Architecture Golang


> **Independent , Testable, and Clean** 

Dalam konsep clean architecture, setiap komponen yang ada bersifat independen dan tidak bergantung pada library external yang spesifik. Seperti tidak tergantung pada spesifik framework atau tidak bergantung pada spesifik database yang dipakai.

[Uncle Bob Clean Architecture](https://www.youtube.com/watch?v=Nsjsiz2A9mg)

Oleh Uncle Bob beliau menyebutkan 4 layer pada arsitekturnya:
1. Entities
2. Usecase
3. Controller
4. Framework dan Driver


## Entities
Layer ini merupakan layer yang menyimpan model yang dipakai

```
type Customer struct {
	Id   int
	Name string
}

type Item struct {
	Id        int
	Name      string
	Available bool
}
```


## Usecase
Layer ini merupakan layer yang akan bertugas sebagai pengontrol, yakni menangangi business logic pada setiap domain. Layer ini juga bertugas memilih repository apa yang akan digunakan, dan domain ini bisa memiliki lebih dari satu repository layer.

```
func IsPembelian(pembeli User) bool {
	if pembeli.Status == 1 && pembeli.CountPembelian <= 5 && pembeli.CountItem <= 10 {
		return true
	} else if pembeli.Status == 2 && pembeli.CountItem <= 10 {
		return true
	} else {
		return false
	}
}

func (pembeli *User) BeliBarang() bool {
	if IsPembelian(*pembeli) == true {
		pembeli.CountPembelian++
		return true
	} else {
		return false
	}
}
```
Untuk pemakaian seperti apa bisa dilihat di Project simple-order-golang
[Disini](https://gitlab.warungpintar.co/enrinal/intern-diary/tree/master/Simple-Order-System-Management-Golang)