package pool

var PageSize uint16 = 4096

// Page is the unit of the Bufferpool
type Page struct {
	Id    uint64     // 8 byte
	prev  *Page      // 8 byte
	next  *Page      // 8 byte
	Dirty bool       // 1 byte
	Data  [4071]byte // 4071 byte (4096 - 8 - 8 - 8 - 1)
}

func NewPage(id uint64) *Page {
	return &Page{Id: id, Dirty: true, prev: nil, next: nil, Data: [4071]byte{}}
}
