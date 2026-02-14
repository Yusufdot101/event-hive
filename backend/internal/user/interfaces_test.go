package user

import (
	"testing"
)

func TestPasswordFlow(t *testing.T) {
	tests := []struct {
		name           string
		setValue       string
		matchValue     string
		wantSetError   error
		wantMatchError error
	}{
		{
			name:           "correct password matched",
			setValue:       "12345678",
			matchValue:     "12345678",
			wantSetError:   nil,
			wantMatchError: nil,
		},
		{
			name:           "incorrect password matched",
			setValue:       "12345678",
			matchValue:     "abcdefgh",
			wantSetError:   nil,
			wantMatchError: nil,
		},
		{
			name:           "invalid password set",
			setValue:       "123",
			matchValue:     "12345678",
			wantSetError:   errInvalidPassword,
			wantMatchError: nil,
		},
		{
			name:           "invalid password matched",
			setValue:       "12345678",
			matchValue:     "123",
			wantSetError:   nil,
			wantMatchError: errInvalidPassword,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			u := &user{}

			err := u.password.set(test.setValue)
			if err != nil {
				if test.wantSetError == nil {
					t.Fatalf("expected no error, got error=%v", err)
				}
				return
			}

			// expected an error but got none
			if test.wantSetError != nil {
				t.Fatalf("expected error=%v, got no error", test.wantSetError)
			}

			matched, err := u.password.matches(test.matchValue)
			if err != nil {
				if test.wantMatchError == nil {
					t.Fatalf("expected no error, got error=%v", err)
				}
				return
			}
			// expected an error but got none
			if test.wantMatchError != nil {
				t.Fatalf("expected error=%v, got no error", test.wantMatchError)
			}

			if matched != (test.setValue == test.matchValue) {
				t.Fatalf("expected matched=%v, got matched=%v", !matched, matched)
			}
		})
	}
}
