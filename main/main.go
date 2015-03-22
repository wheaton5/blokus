

package main

// Imports
import ( 
    //"github.com/ajstarks/svgo" 
     "web"
    //"os" 
    //"io" 
    //"xml"
    //"fmt" 
) 

//type Board struct{
    //squares [][]Square
//}

type Square struct{
    //width   int `xml:"attr"` 
    //height  int `xml:"attr"`
    //name    string `xml:"attr"` 
    //color     string `xml:"attr"` 
    //text      string `xml:"chardata"`
    x_index   int
    y_index   int
    x   int
    y   int
}

var(
    //g = svg.New(os.Stdout)
    width = 800
    height = 600
)

func hello(val string) string { 
    return "hello " + val 
}

func main() {
    //g.Start(width, height)
    //g.Rect(0,0,width, height, "fill:blue")
        web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")   
}

