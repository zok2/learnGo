package main

import "fmt"

type USB interface {
	read()
	write()
}
type Sir interface {
	Say (s string)string
}
type Computer struct {
	name string
}
type Mobile struct {
	 name string
}

func (m Mobile)read()  {
	fmt.Println("read message...")
}
func (m Mobile)write()  {
	fmt.Println("write message")
}

func (m Mobile)Say(s string) string  {
	fmt.Println("Hi~ s",s)
	return "我没有听清楚"
}

func (c Computer)read()  {
	fmt.Println(c.name)
	fmt.Println("read...")
}
func (c Computer)write()  {
	fmt.Println("write")
}


func do(usb USB)  {
	usb.write()
	usb.read()
}
func doSir(sir Sir,s string)  {
	str := sir.Say(s)
	fmt.Println(str)
}

func main()  {
	mobile := Mobile{
		name: "Iphone",
	}
	do(mobile)
	doSir(mobile,"hi")
	computer := Computer{
		name: "Dell",
	}
	do(computer)

}


