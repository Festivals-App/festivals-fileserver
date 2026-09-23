package manipulate

import "testing"

func TestValidImageDimensions(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
		want   bool
	}{
		// beide 0 → erlaubt (keine Dimension angegeben, Originalgröße)
		{"both zero", 0, 0, true},
		// nur eine Dimension angegeben
		{"only width valid", 100, 0, true},
		{"only height valid", 0, 100, true},
		{"only width at min", 30, 0, true},
		{"only height at min", 0, 30, true},
		{"only width below min", 29, 0, false},
		{"only height below min", 0, 29, false},
		{"only width above max", 3001, 0, false},
		{"only height above max", 0, 3001, false},
		// beide Dimensionen angegeben
		{"both valid", 100, 200, true},
		{"both at min", 30, 30, true},
		{"width below min", 29, 100, false},
		{"height below min", 100, 29, false},
		{"both below min", 29, 29, false},
		{"width above max", 3001, 100, false},
		{"height above max", 100, 3001, false},
		{"both above max", 3001, 3001, false},
		{"width at max", 3000, 100, true},
		{"height at max", 100, 3000, true},
		// negative Werte
		{"negative width", -1, 100, false},
		{"negative height", 100, -1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validImageDimensions(tt.width, tt.height); got != tt.want {
				t.Errorf("validImageDimensions(%d, %d) = %v, want %v", tt.width, tt.height, got, tt.want)
			}
		})
	}
}
