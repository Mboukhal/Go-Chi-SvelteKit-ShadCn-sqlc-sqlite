package profile


type UserProfile struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	Picture   string `json:"picture"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	LastSeen  string `json:"last_seen"`
}


func getUserProfileById() *UserProfile {
	return nil
}

func updateUserProfile() bool {	
	return false
}

func deleteUserProfile() bool {
	return false
}

func createUserProfile() bool {
	return false
}