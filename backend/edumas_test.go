package edumasbackend

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "edumas123"
	userdata.Password = "edumas321"
	userdata.Role = "user"
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	CreateNewUserRole(mconn, "user", userdata)
}

func TestCreateNewAdminRole(t *testing.T) {
	var admindata Admin
	admindata.Username = "edumasmin"
	admindata.Password = "edumasmin1"
	admindata.Role = "admin"
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	CreateNewAdminRole(mconn, "admin", admindata)
}	

func TestDeleteUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var userdata User
	userdata.Username = "edumas"
	DeleteUser(mconn, "user", userdata)
}

func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.Username = "edumas123"
	userdata.Password = "mantap"
	userdata.Role = "user"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "edumasapk")

	// Call the function to create a user and generate a token
	err := CreateUserAndAddToken("your_private_key_env", mconn, "user", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

func CreateNewAdminToken(t *testing.T) {
	var admindata Admin
	admindata.Username = "admin"
	admindata.Password = "mantap"
	admindata.Role = "admin"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "edumasapk")

	// Call the function to create a admin and generate a token
	err := CreateAdminAndAddToken("your_private_key_env", mconn, "admin", admindata)

	if err != nil {
		t.Errorf("Error creating admin and token: %v", err)
	}
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var userdata User
	userdata.Username = "edumas123"
	userdata.Password = "mantap"
	userdata.Role = "user"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestGFCPostHandlerAdmin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var admindata Admin
	admindata.Username = "edumas123"
	admindata.Password = "mantap"
	admindata.Role = "Admin"
	CreateNewAdminRole(mconn, "admin", admindata)
}

func TestReport(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var reportdata Report
	reportdata.Nik = 12121
	reportdata.Title = "Jalan Rusak"
	reportdata.Description = "Di sarijadi ada jalan bolong rusak"
	reportdata.DateOccurred = "18112002"
	reportdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
	CreateNewReport(mconn, "report", reportdata)
}

func TestAllReport(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	report := GetAllReport(mconn, "report")
	fmt.Println(report)
}

func TestGeneratePasswordHash(t *testing.T) {
	password := "ganteng"
	hash, _ := HashPass(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CompareHashPass(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("wew", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var userdata User
	userdata.Username = "tes123"
	userdata.Password = "tes321"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "user", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPass(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CompareHashPass(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)
}

func TestHashFunctionAdmin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var admindata Admin
	admindata.Username = "tes123"
	admindata.Password = "tes321"

	filter := bson.M{"username": admindata.Username}
	res := atdb.GetOneDoc[Admin](mconn, "admin", filter)
	fmt.Println("Mongo Admin Result: ", res)
	hash, _ := HashPass(admindata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CompareHashPass(admindata.Password, res.Password)
	fmt.Println("Match:   ", match)
}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var userdata User
	userdata.Username = "ganteng"
	userdata.Password = "gelis"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func TestIsPasswordValidAdmin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var admindata Admin
	admindata.Username = "ganteng"
	admindata.Password = "gelis"

	anu := IsPasswordValidAdmin(mconn, "admin", admindata)
	fmt.Println(anu)
}

func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var userdata User
	userdata.Username = "edumas"
	userdata.Password = "edumas0"
	userdata.Role = "user"
	CreateUser(mconn, "user", userdata)
}

func TestAdminFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "edumasapk")
	var admindata Admin
	admindata.Username = "edumas"
	admindata.Password = "edumas0"
	admindata.Role = "admin"
	CreateAdmin(mconn, "admin", admindata)
}

func TestLoginn(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "pasabarapk")
	var userdata User
	userdata.Username = "edumas"
	userdata.Password = "edumas0"
	IsPasswordValid(mconn, "user", userdata)
	fmt.Println(userdata)
}

func TestLoginnAdmin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "pasabarapk")
	var admindata Admin
	admindata.Username = "edumas"
	admindata.Password = "edumas0"
	IsPasswordValidAdmin(mconn, "admin", admindata)
	fmt.Println(admindata)
}
