package server
import(
	"fmt"
	"net/http"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	//"io/ioutil"
)
func Start(port string){
	mar := martini.Classic()
	mar.Use(render.Renderer())

	  mar.Get("/", func(r render.Render) {
	    r.HTML(200, "hello", "")
	  })
	mar.Post("/",  func(w http.ResponseWriter,  re *http.Request,r render.Render) {
	re.ParseForm()
	fmt.Println(len(re.Form["username"])-1)
	for _,t := range re.Form["username"]{
		//fmt.Println(t)		
		r.HTML(200, "main",t)
	}
	fmt.Printf("asdasd:  %T",re.Form["username"])
	//res := re.Form["username"][1:len(re.Form["username"])-1]
        
	})
	mar.RunOnAddr(":"+port)  
}
