package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	// TODO

	zahleneinsbisneunzehn := []string{"eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf", "dreizehn", "vierzehn", "fünfzehn", "sechzehn", "siebzehn", "achtzehn", "neunzehn"}
	zahlnull := "null"

	if n == 0 {
		return zahlnull
	}
	if n > 0 && n < 20 {
		return zahleneinsbisneunzehn[n]
	}

	result := []int{}
	for n != 0 {
		r := n % 10
		n = n / 10

		result = append(result, r)
	}

	//return zahlen[n]
}
