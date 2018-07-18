package main

import (
    "fmt"
    "log"
    "net/http"
    "io"
  //  "bytes"
    "strings"
    "os"
    "./quickstart"
    "time"
)

func loadFile(w http.ResponseWriter, r *http.Request) {
	//var Buf bytes.Buffer
    // in your case file would be fileupload
    file, header, err := r.FormFile("file_field")
    if err != nil {
        panic(err)
    }
    defer file.Close()
    name := strings.Split(header.Filename, ".")
    fmt.Printf("File name %s\n", name[0])
    
    fmt.Printf("Creating new file\n")
 	out, err := os.Create("./teste.jpg")
 	if err != nil {
 		fmt.Printf("Unable to create the file for writing. Check your write access privilege\n")
 		return
 	}

 	defer out.Close()

 	// write the content from POST to the file
 	byCp, errr := io.Copy(out, file)
 	fmt.Printf("Copied " +  string(byCp) + "\n");
 	if errr != nil {
 		fmt.Fprintln(w, errr)
 	}

 	fmt.Printf("File uploaded successfully : \n")
 	fmt.Printf(header.Filename + "\n")
    return
}
func loadHandler(w http.ResponseWriter, r *http.Request) {
    loadFile(w , r)

    base, data := LoadImage()
    //fmt.Println(GetStateQueue(base, data))
    currentTime := time.Now()
    quickstart.WriteString(currentTime.Format("02/01/2006"), currentTime.Format("15:04:05"), GetStateQueue(base, data));
}

func main() {
    http.HandleFunc("/", loadHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}