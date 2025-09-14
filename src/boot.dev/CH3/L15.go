package main

func reformat(message string, formatter func(string) string) string {
	message1 := formatter(message)
	message2 := formatter(message1)
	message3 := formatter(message2)
	return "TEXTIO: " + message3
}
