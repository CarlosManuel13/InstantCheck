package main

import (
    "fmt"
	"net/http"
	//"os"
	//"strings"
)

type Url struct {
    Links []string
}

func CheckServers(link string, c chan string) {


        res, err := http.Get(link)
        
		if err != nil {
            c <- fmt.Sprintf("❌ %s ha fallado\n", link)
			return
        }

        // El 'else' debe "abrazar" las llaves
        if res.StatusCode == http.StatusOK {
            c <- fmt.Sprintf("✅ %s está OK!\n", link)
        } else { // <--- ¡Esto debe ir así, en la misma línea!
            c <- fmt.Sprintf("⚠️ %s devolvió código %d\n", link, res.StatusCode)
        }
        
        defer res.Body.Close()
}


func main() {

    urls := Url{
        Links: []string{
            "www.google.com", 
           "https://www.youtube.com", 
           "https://discord.com",
			"https://www.google.com", 
			"1.1.1.1", //Cloudflare DNS
            "http://github.com", // Redirige de http a https (301)
            "https://httpbin.org/status/404", // Forzamos un 404
            "https://httpbin.org/status/500", // Forzamos un error de servidor
            "https://jigsaw.w3.org/HTTP/300/301.html",
        },
	}
	c := make(chan string)
    for _, link := range urls.Links {
		go CheckServers(link, c)
	}

	for i := 0; i < len(urls.Links); i++ {
		fmt.Println(<- c)
	}
    
}



    // 1. Field name first: "Links:" 
    // 2. Type second: "[]string{...}"


//**********************************************
//****ENV VARIABLES METHOD**********************
//	varServers := os.Getenv("LINKS")
//	CheckServers(strings.Split(varServers, ","))
//**********************************************


//*******************************************
//********FILE METHOD************************
//data, err := os.ReadFile("links.txt")
//	if err != nil {
//		fmt.Println("❌ Error al leer el archivo:", err)
//		return
//	}


	// 2. Convertir bytes a string y separar por saltos de línea
//	content := string(data)
//	rawLinks := strings.Split(content, "\n")

	// 3. Limpiar espacios y preparar lista final
//	var links []string
//	for _, l := range rawLinks {
//		l = strings.TrimSpace(l)
//		if l != "" { // Evitamos líneas vacías
//			links = append(links, l)
//		}
//	}
//******************************************************************
