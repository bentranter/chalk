package chalk

import "testing"

func TestBlack(t *testing.T) {
	if Black("test") != "\033[30mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestRed(t *testing.T) {
	if Red("test") != "\033[31mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestGreen(t *testing.T) {
	if Green("test") != "\033[32mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestYellow(t *testing.T) {
	if Yellow("test") != "\033[33mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBlue(t *testing.T) {
	if Blue("test") != "\033[34mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestMagenta(t *testing.T) {
	if Magenta("test") != "\033[35mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestCyan(t *testing.T) {
	if Cyan("test") != "\033[36mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestWhite(t *testing.T) {
	if White("test") != "\033[37mtest\033[0m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgBlack(t *testing.T) {
	if BgBlack("test") != "\u001b[40mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgRed(t *testing.T) {
	if BgRed("test") != "\u001b[41mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgGreen(t *testing.T) {
	if BgGreen("test") != "\u001b[42mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgYellow(t *testing.T) {
	if BgYellow("test") != "\u001b[43mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgBlue(t *testing.T) {
	if BgBlue("test") != "\u001b[44mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgMagenta(t *testing.T) {
	if BgMagenta("test") != "\u001b[45mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgCyan(t *testing.T) {
	if BgCyan("test") != "\u001b[46mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestBgWhite(t *testing.T) {
	if BgWhite("test") != "\u001b[47mtest\u001b[49m" {
		t.Errorf("Expected true, got false")
	}
}

func TestUnderline(t *testing.T) {
	if Underline("test") != "\u001b[4mtest\u001b[24m" {
		t.Errorf("Expected true, got false")
	}
}
