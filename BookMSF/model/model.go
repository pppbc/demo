package model

//数据模型,封装一些数据的基本操作

//
////初始化
//func InitDB() {
//	//10个用户
//	for i := 0; i < 10; i++ {
//		u := &User{ID: i + 1, Name: "USER" + strconv.Itoa(i+1), Age: 17 + i, Password: "123456", Tel: "18890961231", Desc: "我就是我，不一样的烟火"}
//		users = append(users, u)
//	}
//
//	//2个图书管理员
//	BookM1 := &User{ID: 101, Name: "BM1", Password: "123456", Age: 22, IsM: 1, Tel: "13490783341"}
//	BookM2 := &User{ID: 102, Name: "BM2", Password: "123456", Age: 30, IsM: 1, Tel: "13466669999"}
//	users = append(users, BookM1)
//	users = append(users, BookM2)
//
//	//一个权限管理
//	AuthorM1 := &User{ID: 201, Name: "AM1", Password: "123456", Age: 22, IsM: 2, Tel: "17770809012", Desc: "无敌是多么的寂寞！！"}
//	users = append(users, AuthorM1)
//
//	//10本书
//	b0 := &Book{1901, "西游记", "吴承恩", 12.5, 10, "../book10.jpg", "师徒的故事"}
//	b1 := &Book{1902, "与时间赛跑的大熊", "卢克.xixi", 12.5, 20, "http://img3m5.ddimg.cn/32/16/27925655-1_w_9.jpg", "熊的故事"}
//	b2 := &Book{1903, "水浒传", "施耐庵", 11.5, 20, "http://img3m1.ddimg.cn/34/3/27932191-1_w_3.jpg", "黑多人的故事"}
//	b3 := &Book{1904, "狂人日记", "周树人", 12.5, 1, "http://img3m8.ddimg.cn/17/30/27915938-1_w_12.jpg", ""}
//	b4 := &Book{1905, "红楼梦", "曹雪芹", 12.5, 20, "http://img3m6.ddimg.cn/30/27/27907536-1_w_13.jpg", ""}
//	b5 := &Book{1906, "养育男孩套装", "史蒂夫.比达尔福", 12.5, 40, "http://img3m1.ddimg.cn/40/27/27922891-1_w_5.jpg", ""}
//	b6 := &Book{1907, "嘻嘻", "lisi", 12.5, 20, "http://img3m0.ddimg.cn/37/36/27911800-1_w_3.jpg", "下嘻嘻嘻嘻嘻嘻"}
//	b7 := &Book{1908, "我们都曾受过伤，却", "zhangsan", 12.5, 20, "../book4.jpg", ""}
//	b8 := &Book{1909, "vue从入门到放弃", "zhangsan", 12.5, 20, "http://img3m4.ddimg.cn/89/1/27911654-1_w_3.jpg", "哈哈啊哈"}
//	b9 := &Book{1910, "余秋雨-中国文化课", "余秋雨", 12.5, 20, "http://img3m9.ddimg.cn/84/6/27914619-1_w_5.jpg", ""}
//
//	books = append(books, b0)
//	books = append(books, b1)
//	books = append(books, b2)
//	books = append(books, b3)
//	books = append(books, b4)
//	books = append(books, b5)
//	books = append(books, b6)
//	books = append(books, b7)
//	books = append(books, b8)
//	books = append(books, b9)
//
//	for _, v := range users {
//		v.Password = tools.Md5Encode(v.Password)
//	}
//
//	//两条借书记录
//	l1 := &Lend{"USER1", "狂人日记", "2019-09-01 12:12:01", "已还", "2019-09-10 13:01:01"}
//	l2 := &Lend{"AM1", "vue从入门到放弃", "2019-07-10 09:10:01", "在借", ""}
//
//	lendList = append(lendList, l1)
//	lendList = append(lendList, l2)
//}
