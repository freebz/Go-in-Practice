// Listing 5.20  Gnerative test for pad

func TestPadGenerative(t *testing.T) {
	fn := func(s string, max uint8) bool {
		p := Pad(s, uint(max))
		return len(p) == int(max)
	}

	if err := quick.Check(fn, &quick.Config(MaxCount: 200)); err != nil {
		t.error(err)
	}
}
