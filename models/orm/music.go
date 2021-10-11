package orm

type Music struct {
	Link
}

func (m Music) GetMusicJson() interface{} {
	return nil
}
