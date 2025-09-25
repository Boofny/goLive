package middleware
import(
)

func CORS(){

}
func CustomCORS(){

}

//on top of the allowed all origin * cors with all the get put post delete 
//need to add a custom url for header and origin

// e := echo.New()
// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
//   AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
//   AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
// }))
