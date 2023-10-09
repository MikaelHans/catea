package util

func CheckIfMemberExists(email string) (bool, error){
	rows, err := GetMemberWithEmail(email)

	if (err != nil){
		return false, err
	}

	counter := 0
	for rows.Next() {
	// you can even scan+store the result if you need them later
		counter++
	}

	if counter > 0{
		return true, err
	}
	return false, err
}