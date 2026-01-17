package store

import "sync"

var (
	tokenStore = make(map[string]int)
	mu         sync.Mutex
)

func SaveToken(token string, userID int) {
	mu.Lock()
	defer mu.Unlock()
	tokenStore[token] = userID
}
func GetUserIDByToken(token string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	id, ok := tokenStore[token]
	return id, ok
}
func DeleteToken(token string) {
	mu.Lock()
	defer mu.Unlock()
	delete(tokenStore, token)
}
