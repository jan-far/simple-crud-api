package config 
  
import ( 
	"github.com/jinzhu/gorm" 
	_ "github.com/jinzhu/gorm/dialects/mysql" 
) 
 
var ( 
	db * gorm.DB 
) 
 
// Connect to db
func Connect() { 
	// Please define your user name and password for my sql. 
	d, err := gorm.Open("mysql", "janfar@/book?charset=utf8&parseTime=True&loc=Local") 
	if err != nil{ 
		panic(err) 
	} 
	db = d 
} 
 
func GetDB() *gorm.DB { 
	return db 
} 
