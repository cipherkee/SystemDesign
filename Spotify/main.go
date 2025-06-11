/*
Requirements:
1. Upload music track
2. Search by title, artist, album, genre
3. Play music track
4. Create playlist
*/
package main

type GenreType string

const (
	GenreTypePop GenreType = "Pop"
)

type Track struct {
	Id     string
	Title  string
	Artist string
	Album  string
	Genre  GenreType
}

type Playlist struct {
	Id     string
	Name   string
	Tracks []Track
}

func (p Playlist) AddTrack(track Track) {
	p.Tracks = append(p.Tracks, track)
}

func (p Playlist) RemoveTrack(trackId string) {
	for i, track := range p.Tracks {
		if track.Id == trackId {
			p.Tracks = append(p.Tracks[:i], p.Tracks[i+1:]...)
			break
		}
	}
}

func (p Playlist) ListTracks() []Track {
	return p.Tracks
}

func (p Playlist) Rename(name string) {
	p.Name = name
}

type MusicPlayer struct {
	tracks     map[string]Track
	titleIndex map[string][]string // Index by track title

	playlists   map[string]Playlist // name is unique, global playlist.... Or keep Id as key. And Allow access to users
	plnameIndex map[string][]string // Index by playlist name
}

func (m *MusicPlayer) UploadTrack(track Track) error {
	// Logic to upload track
	//m.tracks[track.Id] = track
	return nil
}

func (m *MusicPlayer) SearchTracks(query string) ([]Track, error) {
	// Logic to search tracks by title, artist, album, genre
	return []Track{}, nil
}

func (m *MusicPlayer) PlayTrack(trackId string) error {
	// Logic to play track
	return nil
}

func (m *MusicPlayer) CreatePlaylist(name string, tracks []Track) (Playlist, error) {
	// Logic to create a playlist
	playlist := Playlist{
		Id:     "some_unique_id",
		Name:   name,
		Tracks: tracks,
	}
	return playlist, nil
}

func main() {
	// Example usage of the MusicPlayer
	musicPlayer := &MusicPlayer{
		tracks:      make(map[string]Track),
		titleIndex:  make(map[string][]string),
		playlists:   make(map[string]Playlist),
		plnameIndex: make(map[string][]string),
	}

	// Upload a track
	track := Track{
		Id:     "1",
		Title:  "Dua Lipa - Levitating",
		Artist: "Dua Lipa",
		Album:  "Levitating",
		Genre:  GenreTypePop,
	}
	musicPlayer.UploadTrack(track)

	// Create a playlist
	newPlaylist, _ := musicPlayer.CreatePlaylist("Favorites", []Track{track})
	musicPlayer.playlists[newPlaylist.Id] = newPlaylist

	// Add track to playlist
	newPlaylist.AddTrack(track)

	// List tracks in the playlist
	tracksInPlaylist := newPlaylist.ListTracks()
	for _, t := range tracksInPlaylist {
		println(t.Title)
	}
}
