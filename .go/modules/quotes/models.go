package quotes

/**
 * 语录请求参数
 * @author oddfar
 */
type QuoteRequest struct {
	Type   string `form:"c" json:"c"`
	Encode string `form:"encode" json:"encode"`
}

/**
 * 语录响应（JSON格式）
 * @author oddfar
 */
type QuoteResponse struct {
	Code string    `json:"code"`
	Msg  string    `json:"msg"`
	Type string    `json:"type"`
	Data QuoteData `json:"data"`
	From string    `json:"from"`
}

/**
 * 语录数据
 * @author oddfar
 */
type QuoteData struct {
	Text string `json:"text"`
}

/**
 * 文件语录结构
 * @author oddfar
 */
type FileQuote struct {
	ID       int    `json:"id"`
	UUID     string `json:"uuid"`
	Hitokoto string `json:"hitokoto"`
	Type     string `json:"type"`
	From     string `json:"from"`
	FromWho  string `json:"from_who"`
}

/**
 * 数据库语录结构
 * @author oddfar
 */
type DBQuote struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}
