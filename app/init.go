package app

import (
	"github.com/revel/revel"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/koreset/homef/app/models"
	_ "github.com/go-sql-driver/mysql"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string
)
var DB *gorm.DB

// init db
func InitDB() {
	var err error
	author := models.Author{}
	photo := models.Photo{}
	content := models.Content{}
	category := models.Category{}

	// open db
	driver := revel.Config.StringDefault("db.driver", "mysql")
	conn_string := revel.Config.StringDefault("db.connect", "root:wordpass15@/homefdb?parseTime=True&loc=Local&charset=utf8")
	log.Println("Connection String:::::: ", conn_string)
	db, err := gorm.Open(driver, conn_string)
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&content,
		&author,
		&photo,
		&category)
	db.Model(&content).Related(&photo)
	db.LogMode(true) // Print SQL statements

	DB = db

	if err != nil {
		log.Println("FATAL", err)
		panic(err)
	}
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
