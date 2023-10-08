package main

func (s *Storage) persistPost(post *Post) error { //Storage is injected to method
	res := s.db.Create(post)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
