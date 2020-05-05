package viewmodel

// ArticleListModel 文章列表模型
type ArticleListModel struct {
	ArticleTypes *[]ArticleListTypeModel
}

// ArticleListTypeModel 文章分类
type ArticleListTypeModel struct {
	TypeName   string
	CreateTime int64
	Articles   []*ArticleListItemModel
}

// ArticleListItemModel 文章列表文章信息
type ArticleListItemModel struct {
	Title  string
	PinYin string
	Date   string
}