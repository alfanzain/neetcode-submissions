type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var jsonData, _ = json.Marshal(strs)

	return string(jsonData)
}

func (s *Solution) Decode(encoded string) []string {
	var decoded []string
	json.Unmarshal([]byte(encoded), &decoded)

	return decoded
}