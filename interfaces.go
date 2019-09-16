package parsebase

// Interface ensure to get short news
type ShortNewsInterface interface {
	GetShortNews() []*ShortNews
}

type ArticlesInterface interface {
	getArticles() []*Article
}


// Base type of the short news
type ShortNews struct {
	Title       string
	Description string
	Url         string
	Article     []string
}

type Article struct {
	Title string
	Description string
	Url string
	Content []string
}