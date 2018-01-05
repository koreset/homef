package controllers

import (
	"github.com/revel/revel"
	"github.com/koreset/homef/app"
	"github.com/koreset/homef/app/models"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Populate() revel.Result {

	//var dbAuthors []models.Author
	//
	//app.DB.Find(&dbAuthors)
	//
	//fmt.Println(dbAuthors)
	//
	//contents := []models.Content{
	//	{Title: "Article One is the start of new things",
	//		Body: "Lorem Ipsum has been the standard dummy text ever since the 1500s, " +
	//			"when an unknown printer took a galley of type and scrambled it to make a type specimen book. Ahem!!",
	//		Author: dbAuthors[0], AuthorID: dbAuthors[0].ID},
	//	{Title: "Article Two is an nice follow up to One",
	//		Body: "Lorem Ipsum has been the standard dummy text ever since the 1500s, " +
	//			"when an unknown printer took a galley of type and scrambled it to make a type specimen book. Ahem!!",
	//		Author: dbAuthors[1], AuthorID: dbAuthors[1].ID},
	//	{Title: "Article Three is icing on the cake",
	//		Body: "Lorem Ipsum has been the standard dummy text ever since the 1500s, " +
	//			"when an unknown printer took a galley of type and scrambled it to make a type specimen book. Ahem!!",
	//		Author: dbAuthors[0], AuthorID: dbAuthors[0].ID},
	//}
	//categories := []string{"Sustainable Development", "Home School", "Oil Politics"}
	//authors := []models.Author{
	//	{Name:"Jome Akpoduado", Email:"jome@example.com"},
	//	{Name:"Nnimmo Bassey", Email:"nnimmo@base.com"},
	//}
	//
	//for _, val := range categories {
	//	DB.NewRecord(models.Category{Name:val})
	//	DB.Create(&models.Category{Name:val})
	//	fmt.Println(val)
	//}

	//for _, val := range authors {
	//	app.DB.NewRecord(val)
	//	app.DB.Create(&val)
	//	fmt.Println(val)
	//}

	//for _, val := range contents {
	//	app.DB.NewRecord(val)
	//	app.DB.Create(&val)
	//	fmt.Println(val)
	//}

	fmt.Println(app.DB.HasTable("articles"))
	//app.DB.Model(&models.Content{}).AddForeignKey("author_id", "authors(id)", "RESTRICT", "RESTRICT")
	return c.RenderJSON(`{"message":"Success and victory"}`)
}

func (c App) Index() revel.Result {
	var contents []models.Content

	app.DB.Find(&contents)

	for _, item := range contents {
		fmt.Println(item.Title)
		fmt.Println(item.CreatedAt.Format("Jan 2 2006"))
		fmt.Println(item.CreatedAt)
	}
	greeting := "Welcome to Home of Mother Earth"
	message := "Lorem ipsum would be great to have here..."
	c.Flash.Success("Great work going here my friend")
	return c.Render(greeting, message, contents)
}
