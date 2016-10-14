package models


type User struct {
	BaseModel				`xorm:"extends"`

	/**********************************个人信息**************************************/
	//姓名
	Name 		string		`xorm:"varchar(20)"`
	//用户名
	Username	string		`xorm:"varchar(20)"`
	//密码
	Password	string		`xorm:"varchar(128)"`
	//电子邮箱
	Email		string		`xorm:"varchar(255)"`
	//电话号码
	Phone		string 		`xorm:"varchar(20)"`

	/**********************************统计信息**************************************/
	//用户星级
	Stars		int32

	ArticleNum	int32


	//是否已删除，软删除
	Deleted		string
}