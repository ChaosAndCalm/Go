package hello

import (
	"testing"
)

func TestSayHello(t *testing.T){

	subtests := []struct{
		items[] string
		result string
	}{
		{
			result : "Hello, World!\n",
		},{
			items : []string{"Matt"},
			result : "Hello, Matt!\n",
		},{
			items : []string{"Matt","Anne"},
			result : "Hello, Matt Anne!\n",
		},
	}

	for _,st := range subtests {

		if s := Say(st.items); s != st.result {
			t.Errorf("wanted %s (%v) got %s",st.result,st.items,s)
		}
	}

}