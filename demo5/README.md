# elastic - golang 的 elasticsearch 客户端


## 小结
1. matchQuery与termQuery区别：
    + matchQuery：会将搜索词分词，再与目标查询字段进行匹配，若分词中的任意一个词与目标字段匹配上，则可查询到。
    + termQuery：不会对搜索词进行分词处理，而是作为一个整体与目标字段进行匹配，若完全匹配，则可查询到。
2. 快速搭建一个elasticsearch单节点
    + docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.9.0
## 备注
- github.com/olivere/elastic/v7 -> 对应elasticsearch 7.x版本，不同版本还是有很多差异的


