package Repeat

type Value string

func (v *Value) Repeat(count int) string {

	s := *v

	for i := 0; i < count; i++ {

		*v += s
	}
	return string(*v)

}

func (v *Value) Count() int {

	s := string(*v)
	lenght := len(s)
	r := []rune(s)

	var rep int = 1

	for i := 2; i <= lenght/2; i++ {

		flag := true
		for j := 0; j < lenght-i; j++ {
			if r[j] != r[j+i] {
				flag = false
			}

		}
		if flag {
			rep = i
		}

	}
	return lenght / rep
}
