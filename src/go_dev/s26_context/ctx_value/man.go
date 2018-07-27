package main


import(
	"context"
	"fmt"
)

func process(ctx context.Context) {
	// 拿到的是interface需要转一下
	ret, ok:= ctx.Value("trace_id").(int)
	if !ok {
		ret = 1
		fmt.Println("not ok tarce_id")
		return
	}
	fmt.Println("trace_id=", ret)

	session, ok := ctx.Value("session").(string)
	if !ok {
		fmt.Println("not get ct session...")
	}
	fmt.Printf("session=%s\n",session)
}


func main(){
	ctx := context.WithValue(context.Background(), "trace_id", 12345555)    // k/v
	ctx = context.WithValue(ctx, "session", "asdfg123")
	process(ctx)
}