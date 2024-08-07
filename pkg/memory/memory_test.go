package memory

import (
	"testing"

	"github.com/lukepeterson/go8080cpu/pkg/types"
)

func TestReadByteAt(t *testing.T) {
	memory := Memory{
		Data: []byte{0xAA, 0xBB, 0xCC},
	}
	tests := []struct {
		address types.Word
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
				t.Errorf("expected an error for address 0x%04X, but got none", test.address)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for address 0x%04X, but got: %v", test.address, err)
			}
			if result != test.want {
				t.Errorf("expected byte 0x%02X for address 0x%04X, but got 0x%02X", test.want, test.address, result)
			}
		}
	}
}

func TestWriteByteAt(t *testing.T) {
	memory := Memory{
		Data: make([]byte, 4),
	}
	tests := []struct {
		address types.Word
		value   byte
		want    byte
		wantErr bool
	}{
		{address: 0, value: 0xAA, want: 0xAA, wantErr: false},
		{address: 1, value: 0xBB, want: 0xBB, wantErr: false},
		{address: 2, value: 0xCC, want: 0xCC, wantErr: false},
		{address: 4, want: 0, wantErr: true}, // Out-of-bounds
	}

	for _, test := range tests {
		memory.WriteByteAt(test.address, test.value)
		result, err := memory.ReadByteAt(test.address)
		if test.wantErr {
			if err == nil {
				t.Errorf("expected an error for address 0x%04X, but got none", test.address)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for address 0x%04X, but got: %v", test.address, err)
			}
			if result != test.want {
				t.Errorf("expected byte 0x%02X for address 0x%04X, but got 0x%02X", test.want, test.address, result)
			}
		}
	}
}
