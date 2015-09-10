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

func TestUnderline(t *testing.T) {
	if Underline("test") != "\u001b[4mtest\u001b[24m" {
		t.Errorf("Expected true, got false")
	}
}
