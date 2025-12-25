package eggy

import pp "github.com/engmtcdrm/go-prettyprint"

// DefaultAnswerFunc styles the answer text in yellow.
var DefaultAnswerFunc = func(answer string) string {
	return pp.Yellow(answer)
}

// DefaultCursorFunc styles the cursor text in yellow.
var DefaultCursorFunc = func(cursor string) string {
	return pp.Yellow(cursor)
}

// DefaultSelectFunc styles the selected option text in green.
var DefaultSelectFunc = func(s string) string {
	return pp.Green(s)
}

// DefaultIconFunc styles the icon text in cyan.
var DefaultIconFunc = func(icon string) string {
	return pp.Cyan(icon)
}
