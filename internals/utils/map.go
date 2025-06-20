package utils

func MapAge(name string) (int, bool){
	ages:=map[string]int{
		"chair":20,
		"tosh":20,
		"alpha":20,
		"dorito":1,
	}

age, found := ages[name]
return age, found
}