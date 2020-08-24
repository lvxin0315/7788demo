# GoJieba是"结巴"中文分词的Golang语言版本

官方：
- 支持多种分词方式，包括: 最大概率模式, HMM新词发现模式, 搜索引擎模式, 全模式
- 核心算法底层由C++实现，性能高效。
- 字典路径可配置，NewJieba(...string), NewExtractor(...string) 可变形参，当参数为空时使用默认词典(推荐方式)

## demo4
1. demo的场景是为了做产品的搜索引擎，所有其他分词方法没写，详情大家可以看github
2. https://github.com/yanyiwu/gojieba

