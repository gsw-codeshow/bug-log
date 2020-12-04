package main

import "fmt"

type Player struct {
	userid   int
	username string
}

//传入 Player 对象参数
func print_obj(player Player) {
	//player.username = "new"　　//修改并不会影响传入的对象本身
	fmt.Println("userid:", player.userid)
}

//传入 Player 对象指针参数
func print_ptr(player *Player) {
	player.username = "new01"
	fmt.Printf("userid is %d\n", player.userid)
	fmt.Printf("username is %s\n", player.username)
}

//接收者为 Player 对象的方法，方法接收者的变量，按照 GO 语言的习惯一般不用 this/self ，而是使用接收者类型的第一个小写字母，可以看标准库中的代码风格。
func (p Player) m_print_obj() {
	p.username = "new02" // 修改并不会影响传入的对象本身
	fmt.Printf("username is %s\n", p.username)
}

//接收者为 Player 对象指针的方法
func (p *Player) m_print_ptr() {
	p.username = "new03"
	fmt.Printf("userid  is  %d\n", p.userid)
	fmt.Printf("username is %s\n", p.username)
}

func main() {
	pp := &Player{12, "newname"}

	pp.m_print_ptr()
	pp.m_print_obj()
	fmt.Println(pp)
	return
}
