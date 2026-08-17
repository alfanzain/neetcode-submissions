type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return "🤣"
	}

	text := ""
	for i, str := range strs {
		if str == "" {
			text += "😶"
		} else {
			text += str
		}

		if i != len(strs)-1 {
			text += "🤣"
		}
	}

	return text
}

func (s *Solution) Decode(encoded string) []string {
	fmt.Println("encoded:", encoded)

	if encoded == "🤣" {
		return []string{}
	}

	decoded := make([]string, 0)

	text := ""
	for i, code := range encoded {
		currentCode := string(code)
		switch currentCode {
		case "😶":
			decoded = append(decoded, text)
			text = ""
		case "🤣":
			if text == "" {
				continue
			} else {
				decoded = append(decoded, text)
				text = ""
			}
		default:
			text += currentCode

			if i == len(encoded)-1 {
				decoded = append(decoded, text)
			}
		}
	}

	return decoded
}
