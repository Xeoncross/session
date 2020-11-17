package session

// func TestMemoryStory(t *testing.T) {

// 	store := &MemoryStore{}

// 	session := Session{
// 		UserID:    3,
// 		Token:     GenerateRandomToken(24),
// 		CreatedAt: time.Now(),
// 	}

// 	err := store.Save(session)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	sessionCopy, err := store.Get(session.Token)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	// Failed to save
// 	if session.Token != sessionCopy.Token {
// 		t.Errorf("Failed to load token: %q", session.Token)
// 	}

// }
