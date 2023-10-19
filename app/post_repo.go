package app

// Pinging the database for the desired output and DB interactions.

// Storing in DB.
func (s *Storage) persistPost(post *Post) error { //Storage is injected to method
	res := s.db.Create(post)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

// Fetching from DB.
func (s *Storage) getPost(id int) (*Post, error) {
	post := new(Post)

	res := s.db.Find(post, id) // s.db.Get(id)
	if res.Error != nil {
		return nil, res.Error
	}

	//Can further use res.Rows() to check if 0 rows returned matching that passed id.
	return post, nil
}
