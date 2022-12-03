package test

import (
	"fmt"
	"io"
	_ "oldme-api/internal/logic/file"
	"oldme-api/utility/ufile"
	"os"
	"testing"
)

func TestA(t *testing.T) {
	f, _ := os.Open("C:\\Users\\Wry\\Pictures\\ec585c35b6f208cf7f9efb98756ee151.jpeg")
	b, _ := io.ReadAll(f)
	fmt.Printf("%p", b)
	ta := ufile.Ext(b)
	t.Log(ta)
	//service.File().MoveTmp(ctx, "img", "cor5s0e9miocwmjbr4.png")
}
