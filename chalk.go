// Package chalk lets you colour you
// terminal string styles
package chalk

// Black colours your string black
func Black(s string) string {
	return "\033[30m" + s + "\033[0m"
}

// Red colours your string red
func Red(s string) string {
	return "\033[31m" + s + "\033[0m"
}

// Green colours your string green
func Green(s string) string {
	return "\033[32m" + s + "\033[0m"
}

// Yellow colours your string yellow
func Yellow(s string) string {
	return "\033[33m" + s + "\033[0m"
}

// Blue colours your string blue
func Blue(s string) string {
	return "\033[34m" + s + "\033[0m"
}

// Magenta colours your string magenta
func Magenta(s string) string {
	return "\033[35m" + s + "\033[0m"
}

// Cyan colours your string cyan
func Cyan(s string) string {
	return "\033[36m" + s + "\033[0m"
}

// White colours your string white
func White(s string) string {
	return "\033[37m" + s + "\033[0m"
}

func Underline(s string) string {
	return "\u001b[4m" + s + "\u001b[24m"
}
