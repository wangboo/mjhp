# mjhp
通用带癞子的麻将胡牌算法，接入到kafka

算法部分，无赖子牌型生成特征码，用到了 http://blog.csdn.net/panshiqu/article/details/58610958
有癞子，暴力生成带癞子的胡牌组合，然后按照癞子数量，暴力拆解。拆解后的牌型按照6bit一张，拼接成int64的特征码存放。
