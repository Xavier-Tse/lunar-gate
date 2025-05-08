package email

import "sync"

type emailStore struct {
	Email string
	Code  string
}

var StoreMap sync.Map

func Set(emailID, email, code string) {
	StoreMap.Store(emailID, emailStore{
		Email: email,
		Code:  code,
	})
}

func Verify(emailID, email, code string) bool {
	info, ok := StoreMap.Load(emailID)
	if !ok {
		return false
	}
	es, ok := info.(emailStore) // 类型断言
	if !ok {
		return false
	}
	return es.Email == email && es.Code == code
}

func Remove(emailID string) {
	StoreMap.Delete(emailID)
}
