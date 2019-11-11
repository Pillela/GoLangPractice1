package searchingfuncs

//Intlinearsearch will be exported
func Intlinearsearch(intarray [9]int, searchvar int) string {
	for i := 0; i < len(intarray); i++ {
		if intarray[i] == searchvar {
			return "Found"
		}
	}
	return "Not Found"

}
