package main

import (
	"bug-log/flatbuffers/example/MyGame"
	"fmt"

	flatbuffers "github.com/google/flatbuffers/go"
)

func main() {
	builder := flatbuffers.NewBuilder(0)
	str := builder.CreateString("cruis")
	MyGame.MonsterStart(builder)
	MyGame.MonsterAddPos(builder, MyGame.CreateVec3(builder, 1.0, 2.0, 3.0))
	MyGame.MonsterAddMana(builder, 10)
	MyGame.MonsterAddHp(builder, 80)
	MyGame.MonsterAddName(builder, str)
	MyGame.MonsterAddTestType(builder, 1)
	m := MyGame.MonsterEnd(builder)
	builder.Finish(m)
	//  --- --- --- ---
	monster := MyGame.GetRootAsMonster(builder.FinishedBytes(), 0)
	fmt.Println(monster.Hp())
	return
}
