package Math_Functions

func Half_evenorodd(a int) bool{
	var f_half float64
	f_half = float64(a)/2.0
	i_half := a/2
	if f_half==float64(i_half){
		return true
	}else {
		return false
	}
}
