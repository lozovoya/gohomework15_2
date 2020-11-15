package core

// "Ядро" информационной системы. Бизнес-логика.
//
// Пакет хранит информаци о музыкальных группах и их членам,
// а также методы для обработки этой информации.

// Band - музыкальная группа
type Band struct {
	ID        int
	Name      string
	Est       int
	MemberIDs []int `json:"-"`
	Members   []Musician
}

type Musician struct {
	ID   int
	Name string
}

var musicians = []Musician{
	{
		ID:   1,
		Name: "James Hetfield",
	},
	{
		ID:   2,
		Name: "Lars Ulrich",
	},
	{
		ID:   3,
		Name: "Kirk Hammett",
	},
	{
		ID:   4,
		Name: "Robert Trujillo",
	},
}

var bands = []Band{
	{
		ID:        1,
		Name:      "Metallica",
		Est:       1981,
		MemberIDs: []int{1, 2, 3, 4},
	},
}

// BandByName возвращает группу по названию.
func BandByName(name string) Band {
	for _, b := range bands {
		if b.Name == name {
			return b
		}
	}

	return Band{}
}

// BandByMember возвращает группу по имени музыканта.
func BandByMember(name string) Band {

	for _, m := range musicians {
		if m.Name == name {
			for _, b := range bands {
				for _, id := range b.MemberIDs {
					if id == m.ID {
						return b
					}
				}
			}
		}
	}

	return Band{}
}

// Bands возвращает все группы.
func Bands() []Band {
	var result []Band

	for _, b := range bands {
		for _, id := range b.MemberIDs {
			for _, m := range musicians {
				if id == m.ID {
					b.Members = append(b.Members, m)
				}
			}
		}

		result = append(result, b)
	}

	return result
}

// NewBand добавляет группу в список.
func NewBand(band Band) {
	bands = append(bands, band)
}
