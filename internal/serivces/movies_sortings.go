package serivces

import "github.com/moviesalreadytaken/recommendation_service/internal/models/db"

type ByRate []db.Movie

func (m ByRate) Len() int           { return len(m) }
func (m ByRate) Less(i, j int) bool { return m[i].Rate < m[j].Rate }
func (m ByRate) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }


type ByRelaseDate []db.Movie

func (m ByRelaseDate) Len() int           { return len(m) }
func (m ByRelaseDate) Less(i, j int) bool { return m[i].ReleaseDate.Before(m[j].ReleaseDate) }
func (m ByRelaseDate) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
