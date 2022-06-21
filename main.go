package main

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "time"
)

type Stumbel struct {
  Data UserStumble `json:"User"`
}

type UserStumble struct {
  Name string `json:"Username"`
  Country string `json:"Country"`
  Tropy int `json:"SkillRating"`
  Crowns int `json:Crowns"`
  Banned bool `json:"IsBanned"`
}

func GoStumble(auth string) (Stumbel, error) {
  var result Stumbel
  req, err := http.NewRequest("GET", "http://kitkabackend.eastus.cloudapp.azure.com:5010/round/finishv2/3", nil)
  if err != nil {
    return result,err
  }
  req.Header.Set("Authorization", auth)
  resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result,err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result,err
	}
	if err := json.Unmarshal([]byte(body), &result); err != nil {
	  return result,err
	}
	return result,nil
}

func createBannerAndInput() string {
  fmt.Printf(`
███████╗████████╗██╗   ██╗███╗   ███╗██████╗ ██╗     ███████╗
██╔════╝╚══██╔══╝██║   ██║████╗ ████║██╔══██╗██║     ██╔════╝
███████╗   ██║   ██║   ██║██╔████╔██║██████╔╝██║     █████╗  
╚════██║   ██║   ██║   ██║██║╚██╔╝██║██╔══██╗██║     ██╔══╝  
███████║   ██║   ╚██████╔╝██║ ╚═╝ ██║██████╔╝███████╗███████╗
╚══════╝   ╚═╝    ╚═════╝ ╚═╝     ╚═╝╚═════╝ ╚══════╝╚══════╝
                                                             
Create By : %sVnia%s - %sgithub.com/fckvania%s

`, "\033[31m", "\033[0m", "\033[34m", "\033[0m")
  fmt.Print("[+] Please Enter Token: ")
  var input string
  _, err := fmt.Scanln(&input)
  if err != nil {
    fmt.Fprintln(os.Stderr, err)
    return input
  }
  return input
}

func main() {
  input := createBannerAndInput()
  fmt.Println("\n")
  i := 1
  max := 10
  for i < max {
    go func() {
      res, err := GoStumble(input)
      if err != nil {
        log.Fatal("Cokkie Expired!")
      }
      data := res.Data
      if data.Banned {
        log.Fatal("Your Account Is Banned")
      }
      log.Printf("\033[32mNickname : %s | Country : %s | \033[34mSkillRating : %v | Crowns : %v\033[0m\n", data.Name, data.Country, data.Tropy, data.Crowns)
      i += 1
    }()
  }
  //Sleep
  time.Sleep(5 * time.Second)
}