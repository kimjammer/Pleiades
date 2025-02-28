package main

// TODO: Encode/decode JWT

// Create a token for the given user ID
func makeToken(userId string) string {
	return userId
}

// Validate the token for the given user ID, returns an error if the token is invalid.
func verifyToken(token string) (string, error) {
	return token, nil
}
