package main 


import (
    "fmt"
    "io/ioutil"
)

/**
 * describe how page data will be stored in memory
 **/
type Page struct {
    Title string 
    Body []byte
}

/**
 * save the Page's body to a text file
 **/
func (p *Page) save() error {
    filename := p.Title + ".txt"
    // WriteFile will returns an error or nil 
    // 0600 indicate this file should be created with 
    // read-write permissions for the current user only.
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    // ReadFile returns []byte and error
    // '_' is used to throw away the error return value
    // body, _ := ioutil.ReadFile(filename)
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{ Title: title, Body: body }, nil
}

func main() {
    p1 := &Page{ Title: "TestPage", Body: []byte("This is a sample page.") }
    p1.save()
    p2, _ := loadPage("TestPage")
    fmt.Println(string(p2.Body))
}


