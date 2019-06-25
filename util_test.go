package gorose

import (
	"fmt"
	"github.com/gohouse/gocar/varBindValue"
	"testing"
	"time"
)

func TestStructToMap(t *testing.T) {
	user := Users{Uid:1,Name:"gorose"}
	data := StructToMap(user)
	t.Log(data)
}

func TestIf(t *testing.T) {
	closer := func() {
		time.Sleep(1*time.Second)
		fmt.Println("234")
	}
	withRunTimeContext(closer, func(td time.Duration){
		fmt.Println("用时:",td,td.Seconds()>1)
	})
}

func TestStructToMap2(t *testing.T) {
	var u Users
	res := strutForScan(&u)
	fmt.Printf("%#v\n",res)
	for _,item:=range res {
		varBindValue.BindVal(item,1234)
	}
	fmt.Println(u)
}