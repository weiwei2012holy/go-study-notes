package code

import (
    "context"
    "fmt"
    "testing"
    "time"
)

func TestContext(t *testing.T) {
    ctx := context.Background()
    //通过取消信号实现对GO routine的管理控制
    t.Run("WithCancel", func(t *testing.T) {
        ctx, cancel := context.WithCancel(ctx)
        defer cancel()
        go testContentSeeder(ctx)

        time.Sleep(time.Second * 2)
        t.Log("cancel start...")
        cancel() //触发取消信号，testContentSeeder将会执行到取消操作
        t.Log("cancel done...")
        time.Sleep(time.Second * 2)

    })
    //超时控制，区别在于一个传递的是具体时间，一个传递的是时间长度
    t.Run("timeout", func(t *testing.T) {
        t.Run("WithTimeout", func(t *testing.T) {
            ctx, cancel := context.WithTimeout(ctx, time.Second*2)
            defer cancel()
            for i := 0; i < 5; i++ {
                time.Sleep(time.Second)
                testContent(ctx, i)
            }
        })
        t.Run("WithDeadline", func(t *testing.T) {
            ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*2))
            defer cancel()
            for i := 0; i < 5; i++ {
                time.Sleep(time.Second)
                testContent(ctx, i)
            }
        })
    })
    //携带数据，可以是任何类型，但是建议用内置类型，并且避免并发问题
    t.Run("WithValue", func(t *testing.T) {
        ctx := context.WithValue(ctx, "name", "user")
        v, ok := ctx.Value("name").(string) //Value方法：获取设置的key对应的值
        if ok {
            t.Log("name:", v)
        }
        _, ok = ctx.Value("name2").(string)
        if !ok {
            t.Log("name2 not found")
        }
    })

}

func testContent(ctx context.Context, v interface{}) {
    d, ok := ctx.Deadline()
    if ok {
        fmt.Println("ctx has set Deadline...", d)
    }
    err := ctx.Err() //Err方法：当Context被取消或者关闭后，返回context取消的原因
    if err != nil {
        fmt.Println("ctx error...", err)
    }
    select {
    case <-ctx.Done(): //Done方法：当Context被取消或者到了deadline返回一个被关闭的channel
        fmt.Println("ctx done...")

    default:
        fmt.Println("v=", v)
    }
}

func testContentSeeder(ctx context.Context) {
    for range time.Tick(time.Second / 2) {
        select {
        case <-ctx.Done():

            fmt.Println("ctx canceled...", ctx.Err())
            return
        default:
            fmt.Println("seeder running...")
        }
    }

}
