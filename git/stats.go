package git

func Grass() map[string]int {
	commits := Log()

	grass := make(map[string]int)

	for _, commit := range commits {
		t := commit.AuthorDate
		dateStr := t.Format("2006-01-02")

		_, ok := grass[dateStr]
		if !ok {
			grass[dateStr] = 0
		}

		grass[dateStr] += 1
	}

	return grass
}
