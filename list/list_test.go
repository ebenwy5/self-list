package list

import (
	"reflect"
	"testing"
)

func TestNewList(t *testing.T) {
	tests := []struct {
		name string
		want List
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			want: List{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewList(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Add(t *testing.T) {
	type fields struct {
		NodeMap map[int]*Node
		Len     int
		Cap     int
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   List
	}{
		// TODO: Add test cases.
		{
			name: "Success_from_0",
			fields: fields{
				NodeMap: map[int]*Node{},
				Len:     0,
				Cap:     0,
			},
			args: args{
				input: "input",
			},
			want: List{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input"},
				},
				Len: 1,
				Cap: 2,
			},
		},
		{
			name: "Success_amortize_2_to_4",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
				},
				Len: 2,
				Cap: 2,
			},
			args: args{
				input: "input_2",
			},
			want: List{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				NodeMap: tt.fields.NodeMap,
				Len:     tt.fields.Len,
				Cap:     tt.fields.Cap,
			}
			l.Add(tt.args.input)

			for k, v := range tt.want.NodeMap {
				if l.NodeMap[k].Element != v.Element {
					t.Errorf("List.Add() Element %v = %v, want %v", k, l.NodeMap[k].Element, v.Element)
				}
			}

			if l.Len != tt.want.Len {
				t.Errorf("List.Add() Len = %v, want %v", l, tt.want.Len)
			}

			if l.Cap != tt.want.Cap {
				t.Errorf("List.Add() Cap = %v, want %v", l, tt.want.Cap)
			}
		})
	}
}

func TestList_Get(t *testing.T) {
	type fields struct {
		NodeMap map[int]*Node
		Len     int
		Cap     int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
			args: args{index: 1},
			want: "input_1",
		},
		{
			name: "Fail_index_not_exist",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
			args: args{index: 3},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				NodeMap: tt.fields.NodeMap,
				Len:     tt.fields.Len,
				Cap:     tt.fields.Cap,
			}
			if got := l.Get(tt.args.index); got != tt.want {
				t.Errorf("List.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Size(t *testing.T) {
	type fields struct {
		NodeMap map[int]*Node
		Len     int
		Cap     int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				NodeMap: tt.fields.NodeMap,
				Len:     tt.fields.Len,
				Cap:     tt.fields.Cap,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("List.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestList_Remove(t *testing.T) {
	type fields struct {
		NodeMap map[int]*Node
		Len     int
		Cap     int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     string
		wantList List
	}{
		// TODO: Add test cases.
		{
			name: "Success",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
			args: args{index: 1},
			want: "input_1",
			wantList: List{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_2"},
				},
				Len: 2,
				Cap: 4,
			},
		},
		{
			name: "Success_1_element",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
				},
				Len: 1,
				Cap: 2,
			},
			args: args{index: 0},
			want: "input_0",
			wantList: List{
				NodeMap: map[int]*Node{},
				Len:     0,
				Cap:     2,
			},
		},
		{
			name: "Fail_len_0",
			fields: fields{
				NodeMap: map[int]*Node{},
				Len:     0,
				Cap:     0,
			},
			args: args{index: 1},
			want: "",
			wantList: List{
				NodeMap: map[int]*Node{},
				Len:     0,
				Cap:     0,
			},
		},
		{
			name: "Fail_index_not_exist",
			fields: fields{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
			args: args{index: 3},
			want: "",
			wantList: List{
				NodeMap: map[int]*Node{
					0: &Node{Element: "input_0"},
					1: &Node{Element: "input_1"},
					2: &Node{Element: "input_2"},
				},
				Len: 3,
				Cap: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				NodeMap: tt.fields.NodeMap,
				Len:     tt.fields.Len,
				Cap:     tt.fields.Cap,
			}
			if got := l.Remove(tt.args.index); got != tt.want {
				t.Errorf("List.Remove() = %v, want %v", got, tt.want)
			}

			for k, v := range tt.wantList.NodeMap {
				if l.NodeMap[k].Element != v.Element {
					t.Errorf("List.Remove() Element %v = %v, want %v", k, l.NodeMap[k].Element, v.Element)
				}
			}

			if l.Len != tt.wantList.Len {
				t.Errorf("List.Remove() Len = %v, want %v", l, tt.wantList.Len)
			}

			if l.Cap != tt.wantList.Cap {
				t.Errorf("List.Remove() Cap = %v, want %v", l, tt.wantList.Cap)
			}
		})
	}
}
