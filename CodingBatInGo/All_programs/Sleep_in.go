package All_programs


func Sleep_in(weekday bool, vacation bool) bool{
	if weekday==false || vacation == true{
		return true
	}else{
		return false
	}
}