package types

import "testing"

func TestStorage_Commit(t *testing.T) {
	type Data struct {
		n int
	}
	data := Data{
		n: 1,
	}
	s := Storage[Data]{
		States: &data,
		Mutations: map[string]func(*Data, ...any){
			"add": func(data *Data, args ...any) {
				data.n += args[0].(int)
			},
		},
	}
	s.Commit("add", 2)
	if data.n != 3 {
		t.Errorf("Expected 3, got %d", data.n)
	}
}
