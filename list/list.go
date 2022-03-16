package list

type List interface {
	New()
	Insert(element int)
	Delete(element int)
	Find(element int)
}
