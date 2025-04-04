package tools

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"bhaskar": {
		AuthToken: "bhaskar123",
		Username:  "bhaskar",
	},
	"john": {
		AuthToken: "john123",
		Username:  "john",
	},
	"jane": {
		AuthToken: "jane123",
		Username:  "jane",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"bhaskar": {
		Coins:    100,
		Username: "bhaskar",
	},
	"john": {
		Coins:    200,
		Username: "john",
	},
	"jane": {
		Coins:    300,
		Username: "jane",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) (*LoginDetails, error) {
	time
}