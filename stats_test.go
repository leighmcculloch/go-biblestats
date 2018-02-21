package biblestats

import "testing"

func TestBooks(t *testing.T) {
	if g, w := len(Books()), 66; g == w {
		t.Logf("got %d", g)
	} else {
		t.Errorf("got %d, want %d", g, w)
	}

	if g, w := Books()[65], BookRevelation; g == w {
		t.Logf("got %#v", g)
	} else {
		t.Errorf("got %#v, want %#v", g, w)
	}
}

func TestChaptersCount(t *testing.T) {
	if g, w := ChapterCount(BookGenesis), 50; g == w {
		t.Logf("got %d", g)
	} else {
		t.Errorf("got %d, want %d", g, w)
	}
}

func TestVerseCount(t *testing.T) {
	if g, w := VerseCount(BookGenesis, 1), 31; g == w {
		t.Logf("Gen 1 got %d", g)
	} else {
		t.Errorf("Gen 1 got %d, want %d", g, w)
	}

	if g, w := VerseCount(BookRevelation, 22), 21; g == w {
		t.Logf("Rev 2 got %d", g)
	} else {
		t.Errorf("Rev 2 got %d, want %d", g, w)
	}
}
