// Package biblestats contains stats about the books of the Bible. It includes:
//
//   - A list of the books of the Bible.
//   - Count of chapters in each book.
//   - Count of verses in each chapter.
package biblestats

// Book is a book of the Bible.
type Book string

var (
	BookGenesis        Book = "Genesis"
	BookExodus         Book = "Exodus"
	BookLeviticus      Book = "Leviticus"
	BookNumbers        Book = "Numbers"
	BookDeuteronomy    Book = "Deuteronomy"
	BookJoshua         Book = "Joshua"
	BookJudges         Book = "Judges"
	BookRuth           Book = "Ruth"
	Book1Samuel        Book = "1 Samuel"
	Book2Samuel        Book = "2 Samuel"
	Book1Kings         Book = "1 Kings"
	Book2Kings         Book = "2 Kings"
	Book1Chronicles    Book = "1 Chronicles"
	Book2Chronicles    Book = "2 Chronicles"
	BookEzra           Book = "Ezra"
	BookNehemiah       Book = "Nehemiah"
	BookEsther         Book = "Esther"
	BookJob            Book = "Job"
	BookPsalms         Book = "Psalms"
	BookProverbs       Book = "Proverbs"
	BookEcclesiastes   Book = "Ecclesiastes"
	BookSongofSolomon  Book = "Song of Solomon"
	BookIsaiah         Book = "Isaiah"
	BookJeremiah       Book = "Jeremiah"
	BookLamentations   Book = "Lamentations"
	BookEzekiel        Book = "Ezekiel"
	BookDaniel         Book = "Daniel"
	BookHosea          Book = "Hosea"
	BookJoel           Book = "Joel"
	BookAmos           Book = "Amos"
	BookObadiah        Book = "Obadiah"
	BookJonah          Book = "Jonah"
	BookMicah          Book = "Micah"
	BookNahum          Book = "Nahum"
	BookHabakkuk       Book = "Habakkuk"
	BookZephaniah      Book = "Zephaniah"
	BookHaggai         Book = "Haggai"
	BookZechariah      Book = "Zechariah"
	BookMalachi        Book = "Malachi"
	BookMatthew        Book = "Matthew"
	BookMark           Book = "Mark"
	BookLuke           Book = "Luke"
	BookJohn           Book = "John"
	BookActs           Book = "Acts"
	BookRomans         Book = "Romans"
	Book1Corinthians   Book = "1 Corinthians"
	Book2Corinthians   Book = "2 Corinthians"
	BookGalatians      Book = "Galatians"
	BookEphesians      Book = "Ephesians"
	BookPhilippians    Book = "Philippians"
	BookColossians     Book = "Colossians"
	Book1Thessalonians Book = "1 Thessalonians"
	Book2Thessalonians Book = "2 Thessalonians"
	Book1Timothy       Book = "1 Timothy"
	Book2Timothy       Book = "2 Timothy"
	BookTitus          Book = "Titus"
	BookPhilemon       Book = "Philemon"
	BookHebrews        Book = "Hebrews"
	BookJames          Book = "James"
	Book1Peter         Book = "1 Peter"
	Book2Peter         Book = "2 Peter"
	Book1John          Book = "1 John"
	Book2John          Book = "2 John"
	Book3John          Book = "3 John"
	BookJude           Book = "Jude"
	BookRevelation     Book = "Revelation"
)

var books = []Book{
	BookGenesis,
	BookExodus,
	BookLeviticus,
	BookNumbers,
	BookDeuteronomy,
	BookJoshua,
	BookJudges,
	BookRuth,
	Book1Samuel,
	Book2Samuel,
	Book1Kings,
	Book2Kings,
	Book1Chronicles,
	Book2Chronicles,
	BookEzra,
	BookNehemiah,
	BookEsther,
	BookJob,
	BookPsalms,
	BookProverbs,
	BookEcclesiastes,
	BookSongofSolomon,
	BookIsaiah,
	BookJeremiah,
	BookLamentations,
	BookEzekiel,
	BookDaniel,
	BookHosea,
	BookJoel,
	BookAmos,
	BookObadiah,
	BookJonah,
	BookMicah,
	BookNahum,
	BookHabakkuk,
	BookZephaniah,
	BookHaggai,
	BookZechariah,
	BookMalachi,
	BookMatthew,
	BookMark,
	BookLuke,
	BookJohn,
	BookActs,
	BookRomans,
	Book1Corinthians,
	Book2Corinthians,
	BookGalatians,
	BookEphesians,
	BookPhilippians,
	BookColossians,
	Book1Thessalonians,
	Book2Thessalonians,
	Book1Timothy,
	Book2Timothy,
	BookTitus,
	BookPhilemon,
	BookHebrews,
	BookJames,
	Book1Peter,
	Book2Peter,
	Book1John,
	Book2John,
	Book3John,
	BookJude,
	BookRevelation,
}

//go:generate sh -c "go run generate/verses_map.go > verses_map.go"

var chapters = func() map[Book]int {
	m := map[Book]int{}
	for b, c := range verses {
		m[Book(b)] = len(c)
	}
	return m
}()

// Books returns a slice of books in the bible.
func Books() []Book {
	return books
}

// VerseCount returns the number of verses in the book and chapter.
// Returns 0 if the book or chapter is unknown.
func VerseCount(b Book, chapter int) int {
	return verses[string(b)][chapter-1]
}

// ChapterCount returns the number of chapters in the book.
// Returns 0 if the book is unknown.
func ChapterCount(b Book) int {
	return chapters[b]
}
