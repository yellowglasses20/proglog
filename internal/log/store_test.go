package log

import (
	"os"
	"reflect"
	"testing"
)

func Test_newStore(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name    string
		args    args
		want    *store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newStore(tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("newStore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Append(t *testing.T) {
	type args struct {
		p []byte
	}
	f, _ := os.CreateTemp("","store_append_test");
	s, _ := newStore(f);
	tests := []struct {
		name    string
		s       *store
		args    args
		wantN   uint64
		wantPos uint64
		wantErr bool
	}{
		{
			name: "hello1",
			s: s,
			args: args{[]byte("hello world")},
			wantN: 19,
			wantPos: 0,
			wantErr: false,

		},
		{
			name: "hello2",
			s: s,
			args: args{[]byte("hello world2")},
			wantN: 20,
			wantPos: 19,
			wantErr: false,

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, gotPos, err := tt.s.Append(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.Append() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("store.Append() gotN = %v, want %v", gotN, tt.wantN)
			}
			if gotPos != tt.wantPos {
				t.Errorf("store.Append() gotPos = %v, want %v", gotPos, tt.wantPos)
			}
		})
	}
}

func Test_store_Read(t *testing.T) {
	type args struct {
		pos uint64
	}
	tests := []struct {
		name    string
		s       *store
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Read(tt.args.pos)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.Read() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_ReadAt(t *testing.T) {
	type args struct {
		p   []byte
		off int64
	}
	tests := []struct {
		name    string
		s       *store
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ReadAt(tt.args.p, tt.args.off)
			if (err != nil) != tt.wantErr {
				t.Errorf("store.ReadAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("store.ReadAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_store_Close(t *testing.T) {
	tests := []struct {
		name    string
		s       *store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Close(); (err != nil) != tt.wantErr {
				t.Errorf("store.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
