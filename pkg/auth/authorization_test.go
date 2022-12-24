package auth

import "testing"

func Benchmark_And_IsSatisfied(b *testing.B) {
	checker := And{[]string{"super_admin", "admin", "user"}}
	xPermissions := "admin user super_admin"

	for i := 0; i < b.N; i++ {
		_ = checker.IsSatisfied(xPermissions)
	}
}

func Benchmark_Or_IsSatisfied(b *testing.B) {
	checker := Or{[]string{"super_admin", "admin", "user"}}
	xPermissions := "admin user super_admin"

	for i := 0; i < b.N; i++ {
		_ = checker.IsSatisfied(xPermissions)
	}
}

func Test_And_IsSatisfied(t *testing.T) {
	tests := []struct {
		name             string
		havePermissions  []string
		haveXPermissions string
		wantIsSatisfied  bool
	}{
		{
			"not satisfied when no permissions were required",
			nil,
			"admin user guest",
			false,
		},
		{
			"not satisfied when no permissions were found in header",
			[]string{"admin", "user", "guest"},
			"",
			false,
		},
		{
			"not satisfied when at least one permission was not found in header",
			[]string{"admin", "user", "guest"},
			"admin guest",
			false,
		},
		{
			"satisfied when all permissions were found in header",
			[]string{"admin", "user", "guest"},
			"user admin guest",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := And{Permissions: tt.havePermissions}

			if ok := checker.IsSatisfied(tt.haveXPermissions); ok != tt.wantIsSatisfied {
				t.Errorf("expected %v got %v", tt.wantIsSatisfied, ok)
			}
		})
	}
}

func Test_Or_IsSatisfied(t *testing.T) {
	tests := []struct {
		name             string
		havePermissions  []string
		haveXPermissions string
		wantIsSatisfied  bool
	}{
		{
			"not satisfied when no permissions were required",
			nil,
			"read write execute",
			false,
		},
		{
			"not satisfied when no permissions were found in header",
			[]string{"read", "write", "execute"},
			"",
			false,
		},
		{
			"satisfied when at least one permission was found in header",
			[]string{"read", "write", "execute"},
			"user admin read guest accounts",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			checker := Or{Permissions: tt.havePermissions}

			if ok := checker.IsSatisfied(tt.haveXPermissions); ok != tt.wantIsSatisfied {
				t.Errorf("expected %v got %v", tt.wantIsSatisfied, ok)
			}
		})
	}
}
