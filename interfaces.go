package parsebase

// Interface ensure to get short news
type ShortNewsInterface interface {
	GetShortNews() ([]*ShortNews, error)
}

type ArticlesInterface interface {
	GetArticles() ([]*Article, error)
}

// Base type of the short news
type ShortNews struct {
	Title       string
	Description string
	Url         string
	Article     []string
}

type Article struct {
	Title       string
	Description string
	Url         string
	Content     []string
}
