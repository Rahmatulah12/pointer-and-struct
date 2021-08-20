package main

import (
	"fmt"
)

func main() {

	/*
		Pointer
		Pointer adalah referensi atau alamat memory.
		Variabel pointer berarti variabel yang menampung alamat memori suatu nilai.
		Variabel bertipe pointer ditandai dengan adanya tanda asterisk ( * )
		tepat sebelum penulisan tipe data ketika deklarasi.

		Variabel biasa sebenarnya juga bisa diambil nilai pointernya,
		caranya dengan menambahkan tanda ampersand ( & ) tepat sebelum nama variabel.
		Metode ini disebut dengan referencing

		Dan sebaliknya, nilai asli variabel pointer juga bisa diambil,
		dengan cara menambahkan tanda asterisk ( * ) tepat sebelum nama variabel.
		Metode ini disebut dengan dereferencing
	*/

	// variable biasa
	var numberA int = 4

	// pointer
	var numberB *int = &numberA

	fmt.Println("Value Number A :", numberA)
	fmt.Println("Address Number A :", &numberA) //&numberA => mengambil address variable Number A

	fmt.Println("Value Number B :", *numberB) // *numberB => mengambil value variable Number B
	fmt.Println("Address Number A :", numberB)

	fmt.Println("")
	numberA = 5

	/*
		Ketika salah satu variabel pointer di ubah nilainya,
		sedang ada variabel lain yang memiliki referensi memori yang sama,
		maka nilai variabel lain tersebut juga akan berubah.
	*/
	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")
	// pointer sebagai parameter
	var number = 4
	fmt.Println("before :", number) // 4

	change(&number, 10)
	fmt.Println("after :", number) // 10
	fmt.Println("address :", &number)

	// memanggil struct
	var s1 student
	var s2 = student{}
	s1.name = "John wick"
	s1.grade = 10

	s2.name = "Siti Fatimah"
	s2.grade = 3

	var s3 = student{"Daus", 6}

	var s4 = student{grade: 1, name: "Okta"}

	fmt.Println("Name :", s1.name)
	fmt.Println("Grade :", s1.grade)

	fmt.Println("Name :", s2.name)
	fmt.Println("Grade :", s2.grade)

	fmt.Println("Name :", s3.name)
	fmt.Println("Grade :", s3.grade)

	fmt.Println("Name :", s4.name)
	fmt.Println("Grade :", s4.grade)

	/*
		variable objek pointer
		Objek hasil cetakan struct bisa diambil nilai pointer-nya,
		dan bisa disimpan pada variabel objek yang bertipe struct pointer.
	*/
	var s5 = student{name: "Mahasiswa", grade: 9}
	// variable objek yang bertipe struct pointer
	var s6 *student = &s5

	fmt.Println("S5 Name :", s5.name)
	fmt.Println("S6 Name :", s6.name)

	s6.name = "ehtan"
	fmt.Println("S5 Name :", s5.name)
	fmt.Println("S6 Name :", s6.name)
	fmt.Println("Address Memory :", &s6)

	var emp = employee{}
	emp.person.name = "Ma'ruf Amin"
	emp.person.age = 50
	emp.jabatan = "Wakil Persiden"
	fmt.Println(emp)

	var person = person{name: "Person 2", age: 25}
	var emp2 = employee{person: person, jabatan: "IT Manager"}
	fmt.Println(emp2)

	/*
		Anonymous Struct
		struct yang tidak dideklarasikan di awal, melainkan ketika dibutuhkan saja,
		langsung pada saat penciptaan objek. Teknik ini cukup efisien untuk pembuatan variabel
		objek yang struct nya hanya dipakai sekali
	*/
	var anonymousStruct = struct {
		student
		class int
	}{}

	anonymousStruct.student.name = "Test"
	anonymousStruct.student.grade = 10
	anonymousStruct.class = 3
	fmt.Println(anonymousStruct)

	/*
		array anonymous struct
	*/

	var allStudents = []struct {
		student
		class int
	}{
		{student: student{name: "Ragay", grade: 10}, class: 12},
		{student: student{name: "Ragay", grade: 10}, class: 12},
	}

	fmt.Println(allStudents)

	/*
		Deklarasi Struct dengan var
	*/
	var student4 struct {
		name  string
		grade int
	}

	student4.name = "Bahaya"
	student4.grade = 10

	var student5 = struct {
		name  string
		grade int
	}{
		"Rahmatulah Sidik",
		12,
	}

	fmt.Println(student4)
	fmt.Println(student5)

	// memanggil nested struct
	var student9 = student9{}
	student9.person.name = "Aylin Fariha"
	student9.person.age = 8
	student9.person.gender = "Wanita"
	var hobbies2 = []string{"nenen", "maen"}
	student9.hobbies = hobbies2
	fmt.Println(student9)

	var p1 = struct {
		name string
		age  int
	}{name: "Temon", age: 50}
	fmt.Println(p1)
}

// pointer sebagai parameter
func change(original *int, value int) int {
	*original = value
	return *original
}

/*
	Struct
	 kumpulan definisi variabel (atau property) dan
	 atau fungsi (atau method), yang dibungkus dengan nama tertentu.
	 atau bisa disebut struct itu seperti class dalam oop
*/
type student struct {
	name  string
	grade int
}

/*
	Embeded Struct
	penurunan properti dari satu struct ke struct lain,
	sehingga properti struct yang diturunkan bisa digunakan.
*/
type person struct {
	name string
	age  int
}

type employee struct {
	jabatan string
	person
}

/*
	Nested Struct
*/
type student9 struct {
	person struct {
		name   string
		age    int
		gender string
	}
	hobbies []string
}

/*

	Deklarasi dan inisialisasi struct secara horizontal
	tanda semi colon( ; ) digunakan sebagai pembatas deklarasi proptery
	type person2 struct{ name string; age int; hobbies []string }

*/
type person2 struct {
	name    string
	age     int
	hobbies []string
}
