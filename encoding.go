package uuid

func (uuid UUID) MarshalText() ([]byte, error) {
	return []byte(uuid.String()), nil
}

func (uuid *UUID) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		*uuid = UUID{}
		return nil
	}
	u, err := Parse(string(text))
	if err != nil {
		return err
	}
	*uuid = u
	return err
}
