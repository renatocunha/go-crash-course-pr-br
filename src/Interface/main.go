package main

import "fmt"

type  Ser interface  {
	Falar()
}

type Humano string

func (h Humano) Falar() {
	 fmt.Println("Perdeu Playboy")
}

type Cachorro string

func (c Cachorro) Falar(){
	fmt.Println("Au Au")
}

type Golfinho string

func (g Golfinho) Falar() {
	fmt.Println("Ki Ki Ki")

}

func Falar(s Ser){
	s.Falar()
}

func  main() {
	p1 := Humano("Joao")
	p2 := Cachorro("Toff")
	p3 := Golfinho("Flufy")
	
	fmt.Println(p1 + " diz:")
	Falar(&p1)
	
	fmt.Println(p3 + " diz:")
	Falar(&p2)
	
	fmt.Println(p3 + " diz:")
	Falar(&p3)
}

