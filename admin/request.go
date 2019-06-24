package admin

type AdminAdd struct {
	Name 			string
	Adminname		string
	Password		string
	AdminGroupId 	int64
	Email 			string
	Token			string
}

type AdminLogin struct {
	Adminname 		string
	Password		string
	Token			string
}

type AdminId struct {
	Id 				int64
	Token			string
}

type AdminToken struct {
	Token			string
}

type AdminAdminname struct {
	Adminname		string
	Token			string
}

type AdminChangePassword struct {
	CurrentPassword	string
	NewPassword	 	string
	Token			string
}

type AdminChangeAdminname struct {
	CurrentPassword	string
	NewAdminname	string
	Token			string
}

type AdminChangeName struct {
	CurrentPassword	string
	NewName		 	string
	Token			string
}

type AdminChangeEmail struct {
	Id				int64
	NewEmail	 	string
	Token			string
}

type CountryAdd struct {
	Name			string
	IsoCode	 		string
	DialCode 		string
	Code 			string
	Token			string
}

type CountryUpdate struct {
	Id 				int64
	Name			string
	IsoCode	 		string
	DialCode 		string
	Code 			string
	Token			string
}

type CurrencyAdd struct {
    Code            string 
    Name            string
    Description     string
    Type            string
    CountryId       int64
	Token			string
}

type CurrencyUpdate struct {
	Id 				int64
    Code            string 
    Name            string
    Description     string
    Type            string
    CountryId       int64
	Token			string
}

type UserSearch struct {
	CountryId 		int64
	Eligible 		string
}

type UserGet struct {
	UserSearch 		UserSearch
	Token  			string
}