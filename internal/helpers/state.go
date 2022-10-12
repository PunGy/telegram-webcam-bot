package helpers

import "sync"

type UserState struct {
	Path string
	Role string
}
type BotState struct {
	users map[int64]*UserState
}

/**
* Never returns nil reference. If the user is not present in the state - it creates one and returns
 */
func (s *BotState) GetUser(id int64) *UserState {
	user := s.users[id]

	if user == nil {
		user = &UserState{
			Path: "/",
			Role: "listener",
		}
		s.users[id] = user
	}

	return user
}

var onceState sync.Once

var state BotState

func GetState() *BotState {
	onceState.Do(func() {
		state.users = make(map[int64]*UserState)
	})

	return &state
}
