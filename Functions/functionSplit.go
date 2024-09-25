package function

func ExtractText(text string) [][]string {
	str := ""
	slc := []string{}
	slc_slc := [][]string{}
	for _, s := range text {
		if s == ' ' {
			if str != "" {
				slc = append(slc, str)
				str = ""
			}
			continue
		}
		if s == '\n' {
			if str != "" {
				slc = append(slc, str)
				str = ""
			}
			if len(slc) > 0 {
				slc_slc = append(slc_slc, slc)
				slc = []string{}
			}
			continue
		}
		str += string(s)
	}
	if str != "" {
		slc = append(slc, str)
		str = ""
	}
	if len(slc) > 0 {
		slc_slc = append(slc_slc, slc)
		slc = []string{}
	}

	return slc_slc

}

func Split(text string) [][]string {
	lines := ExtractText(text)

	for i, line := range lines {
		lines[i] = CheckVowel(line)
	}

	for i, line := range lines {
		lines[i] = Flag(line)
	}

	for i, line := range lines {
		lines[i] = FlagWithParams(line)
	}

	return lines
}
