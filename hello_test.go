package hello

import (
	"testing"
)

func TestSayHello(t *testing.T){
	want := "Hello, test!\n"
	got := Say([] string{"test"})

	if want != got {
		t.Errorf("Wanted %s got %s",want,got)
	}

}