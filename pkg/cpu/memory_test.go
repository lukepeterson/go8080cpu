package cpu

import "testing"

func TestReadByteAt(t *testing.T) {
	memory := Memory{
		Data: []byte{0xAA, 0xBB, 0xCC},
	}
	tests := []struct {
		address word
		want    byte
		wantErr bool
	}{
		{address: 0, want: 0xAA, wantErr: false},
		{address: 1, want: 0xBB, wantErr: false},
		{address: 2, want: 0xCC, wantErr: false},
		{address: 4, want: 0, wantErr: true}, // Out-of-bounds
	}

	for _, test := range tests {
		result, err := memory.ReadByteAt(test.address)
		if test.wantErr {
			if err == nil {
				t.Errorf("expected an error for address %d, but got none", test.address)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for address %d, but got: %v", test.address, err)
			}
			if result != test.want {
				t.Errorf("expected byte %x for address %d, but got %x", test.want, test.address, result)
			}
		}
	}
}
