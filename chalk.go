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

// Underline places an underline under your string
func Underline(s string) string {
	return "\u001b[4m" + s + "\u001b[24m"
}

// BgBlack colours the background of your string black
func BgBlack(s string) string {
	return "\u001b[40m" + s + "\u001b[49m"
}

// BgRed colours the background of your string red
func BgRed(s string) string {
	return "\u001b[41m" + s + "\u001b[49m"
}

// BgGreen colours the background of your strinYellowreen
func BgGreen(s string) string {
	return "\u001b[42m" + s + "\u001b[49m"
}

// BgYellow colours the background of your string yellow
func BgYellow(s string) string {
	return "\u001b[43m" + s + "\u001b[49m"
}

// BgBlue colours the background of your string blue
func BgBlue(s string) string {
	return "\u001b[44m" + s + "\u001b[49m"
}

// BgMagenta colours the background of your string magenta
func BgMagenta(s string) string {
	return "\u001b[45m" + s + "\u001b[49m"
}

// BgCyan colours the background of your string cyan
func BgCyan(s string) string {
	return "\u001b[46m" + s + "\u001b[49m"
}

// BgWhite colours the background of your string white
func BgWhite(s string) string {
	return "\u001b[47m" + s + "\u001b[49m"
}
