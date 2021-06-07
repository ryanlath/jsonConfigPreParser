package jsonConfigPreParser

/*
For using JSON in config files
Parses out comments (/* // #), white space and trailing commas in JSON

Use as a pre-parser before json.UnMarshal

Taken from: https://gist.github.com/kurokikaze/1254281
*/

func Parse(p []byte) ([]byte, error) {
	minify := true
	state := "0"
	buf := make([]byte, 0, len(p))

	for _, charcode := range p {
		character := string(charcode)

		switch state {
		case "0": // Start
			switch character {
			case "\"":
				state = "S"
				buf = append(buf, charcode)

			case "/":
				state = "2"

			case "#":
				state = "C"

			case "}", "]":
				// remove trailing commas (needs minify)
				if buf[len(buf)-1] == 44 { // ","
					buf[len(buf)-1] = charcode
				} else {
					buf = append(buf, charcode)
				}

			default:
				//anything outside quote strings/comments/etc. eg. true 8080 , { [ : \n \t \r SPACE
				if !minify || (character != " " && character != "\n" && character != "\t" && character != "\r") {
					buf = append(buf, charcode)
				}
			}

		case "1": // escape symbol inside string
			state = "S"
			buf = append(buf, charcode)

		case "2": // first slash met
			switch character {
			case "/":
				state = "C"
			case "*":
				state = "MC"
			default:
				state = "0"
				buf = append(buf, charcode)
			}

		case "3": // star met inside multiline comment
			if character == "/" {
				state = "0"
			} else {
				state = "MC"
			}

		case "S": // Inside string quote
			switch character {
			case "\"":
				state = "0"
			case "/":
				state = "1"
			}
			buf = append(buf, charcode)

		case "C": // Comment
			if character == "\n" {
				state = "0"
				if !minify {
					buf = append(buf, charcode)
				}
			}

		case "MC": // Multiline comment
			if character == "*" {
				state = "3"
			}

		}
	}

	return buf, nil
}
