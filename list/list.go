package list

type List interface {
	New()
	Insert(element int)
	Delete(element int)
	Get(index int) (int, string)
}
