package session

// type MemoryStore struct {
// 	sessions []Session
// 	m        sync.Mutex
// }

// // Get session using token from client
// func (m *MemoryStore) Get(token string) (Session, error) {
// 	m.m.Lock()
// 	defer m.m.Unlock()

// 	for _, session := range m.sessions {
// 		if session.Token == token {
// 			return session, nil
// 		}
// 	}

// 	return Session{}, nil
// }

// // Saves session provided by HTTP router
// func (m *MemoryStore) Save(session Session) error {
// 	m.m.Lock()
// 	m.m.Unlock()

// 	for i, session := range m.sessions {
// 		if session.Token == session.Token {
// 			m.sessions[i] = session
// 			return nil
// 		}
// 	}

// 	m.sessions = append(m.sessions, session)
// 	return nil
// }

// // Delete session with the given token
// func (m *MemoryStore) Delete(token string) error {
// 	m.m.Lock()
// 	defer m.m.Unlock()

// 	for index, session := range m.sessions {
// 		if session.Token == token {
// 			m.sessions = append(m.sessions[:index], m.sessions[:index+1]...)
// 			return nil
// 		}
// 	}

// 	return nil
// }
