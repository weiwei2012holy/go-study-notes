package code

import (
    "fmt"
    "testing"
)

func TestForWithPointer(t *testing.T) {

    s := []int{1, 2, 3, 4}
    t.Run("for1", func(t *testing.T) {
        var res []*int
        for _, v := range s {
            res = append(res, &v)
        }
        //此处保存的地址永远是同一个，同理取的值也是同一个，因为v变量在循环开始时就已经完成定义，每次循环并不会重新定义变量
        t.Log(s, res) //[1 2 3 4] [0xc0000ac1b8 0xc0000ac1b8 0xc0000ac1b8 0xc0000ac1b8]

        for _, v2 := range res {
            t.Log(*v2) // 4 4 4 4
        }
    })
    t.Run("for2", func(t *testing.T) {
        var res []*int
        for _, v := range s {
            v := v //在循环开始时对v就行重新赋值，则后续对v的操作都是全新的变量
            res = append(res, &v)
        }
        t.Log(s, res) //[1 2 3 4] [0xc0000ac288 0xc0000ac288 0xc0000ac288 0xc0000ac288]
    })

}

func TestForWithClosure(t *testing.T) {
    s := []int{1, 2, 3, 4}

    t.Run("error", func(t *testing.T) {
        var prints []func()
        for _, v := range s {
            prints = append(prints, func() {
                fmt.Println(v)
            })
        }
        for _, f := range prints {
            f()
        }
    })
    t.Run("success", func(t *testing.T) {
        var prints []func()
        for _, v := range s {
            v := v
            prints = append(prints, func() {
                fmt.Println(v)
            })
        }
        for _, f := range prints {
            f()
        }
    })

}
