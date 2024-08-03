package util_test

import (
	"testing"

	"github.com/vldcreation/simple_bank/util"
)

func TestHashPassword(t *testing.T) {
	validPassword := util.RandString(6)
	hashedPassword, err := util.HashPassword(validPassword)
	if err != nil {
		t.Errorf("HashPassword() error = %v", err)
	}

	tests := []struct {
		name     string
		password string
		valid    bool
	}{
		{
			name:     "invalid password",
			password: "test@123",
			valid:    false,
		},
		{
			name:     "valid password",
			password: validPassword,
			valid:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := util.ComparePassword(hashedPassword, tt.password)
			if tt.valid && err != nil {
				t.Errorf("test#%s: ComparePassword() error = %v, want nil", tt.name, err)
			}
		})
	}
}
