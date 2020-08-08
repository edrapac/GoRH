package main
import (
	"fmt"
	"net/http"
	// "io/ioutil"
	"golang.org/x/net/html"
	//"bufio"
	//"os"
	//"strings"
)

func main() {
	base_url := "https://robinhood.com/stocks/AAPL" // base landing page for RH's stock information
	//fmt.Println("Enter a Stock's ticker symbol here. Try AAPL")
	//reader := bufio.NewReader(os.Stdin)
	//fragment, _ := reader.ReadString('\n')
	// base_url += strings.ToUpper(strings.TrimSpace(fragment)) // kind of an ugly way to do it but it gets the job done!
	

	response, err := http.Get(base_url)

	if err != nil {
		panic(err)
	}

	z := html.NewTokenizer(response.Body)
	previousStartTokenTest := z.Token()
	loopOver:
		for {
			tt := z.Next()
			switch {
			case tt == html.ErrorToken:
			// EOF
				break loopOver
			case tt == html.StartTagToken:
				previousStartTokenTest = z.Token()
			case tt == html.TextToken:
				if previousStartTokenTest.Data == "span" {
					continue
				}
				TxtContent := string(z.Text())
				i := 0
				if len(TxtContent) > 0 {
					i += 1
					fmt.Println(TxtContent,i)
					
				}
				
				defer response.Body.Close()
				break
			}
		}
	// bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	
	// fmt.Printf(string(bytes))
}