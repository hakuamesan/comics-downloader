package sites

import (
	"testing"

	"github.com/Girbons/comics-downloader/pkg/core"
	"github.com/stretchr/testify/assert"
)

const (
	URL    = "https://www.mangareader.net/naruto/1/"
	SOURCE = "www.mangareader.net"
)

func TestMangaReadGetInfo(t *testing.T) {
	mr := new(Mangareader)
	name, issueNumber := mr.GetInfo(URL)

	assert.Equal(t, "naruto", name)
	assert.Equal(t, "1", issueNumber)
}

func TestRetrieveMangaReaderImageLinks(t *testing.T) {
	mr := new(Mangareader)

	comic := new(core.Comic)
	comic.URLSource = URL
	comic.Name = "naruto"
	comic.IssueNumber = "1"
	comic.Source = SOURCE

	links, err := mr.retrieveImageLinks(comic)

	assert.Equal(t, 53, len(links))
	assert.Nil(t, err)
}

func TestSetupMangaReader(t *testing.T) {
	mr := new(Mangareader)

	comic := new(core.Comic)
	comic.Name = "naruto"
	comic.IssueNumber = "1"
	comic.URLSource = URL
	comic.Source = SOURCE

	err := mr.Initialize(comic)

	assert.Nil(t, err)
	assert.Equal(t, 53, len(comic.Links))
}

func TestMangareaderRetrieveIssueLinks(t *testing.T) {
	mr := new(Mangareader)
	issues, err := mr.RetrieveIssueLinks("https://www.mangareader.net/naruto", false, false)

	assert.Nil(t, err)
	assert.Equal(t, 700, len(issues))
}