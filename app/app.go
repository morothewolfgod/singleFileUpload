package app

import(
	"log"
    //"fmt"
	"net/http"
	//"html/template"
	// "io/ioutil"
)

type Server struct{
	port string
}

func NewServer() Server{
	return Server{}
}

func (s *Server) Init(port string){
	log.Println("Initializing Server...")
	s.port= ":" + port
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	/*fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		//parse input
		r.ParseMultipartForm(10 << 20)
		
		//retrieving file
		file, handler, err := r.FormFile("myFile")
    	if err != nil {
        fmt.Println(err)
        	return
    	}
    	defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		
		// put temporary file on server
		tempFile, err := ioutil.TempFile("images", "upload-*.png")
    	if err != nil {
        	fmt.Println(err)
    	}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
    	if err != nil {
        	fmt.Println(err)
    	}
		tempFile.Write(fileBytes)
		
		// return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	 }*/
 }

func (s *Server) Start(){
	log.Println("Starting Server")
	//http.HandleFunc("/", uploadFile)
	http.ListenAndServe(s.port, nil)

}

func main() {
	log.Println("Hello World!")
	
}