package server

import(

"github.com/gofiber/fiber/v2"

)


func Run(){
	
		app := fiber.New()

		app.Get("/" , func (ctx *fiber.ctx)error{
		
		return ctx.status(http.statusok).Json(fiber.Map(
			"massage":"Hello I am working",
		))
		
	})
	app.Get(""/relth" ,func(ctx *fiber.ctx)error{
	return ctx.status(http.statusok).Json(fiber.Map{
	"heaithy": true
	"})
})
	app.listen("3000")






		
