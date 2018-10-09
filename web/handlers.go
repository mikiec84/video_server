package main

import (
	"html/template"
	"net/http"
	"net/url"
	"net/http/httputil"
	"encoding/json"
	"log"
	"io"
	"io/ioutil"
	"github.com/julienschmidt/httprouter"
)

type HomePage struct {
	Name string 
}

type UserPage struct {
	Name string 
	VideoUrl string 
	HostUrl string 
}

type UserHomePage struct {

}

type LoginPage struct {
	HostUrl string 
}

func loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//cname, err1 := r.Cookie("username")
	//sid, err2 := r.Cookie("session")

	t, e := template.ParseFiles("./bin/template/login.html")
	if e != nil {
		log.Printf("Parsing template login.html error: %s", e)
		return 
	}

	p := &LoginPage{HostUrl: "192.168.189.134:8080"}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	t.Execute(w, p)
}

func homeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	cname, err1 := r.Cookie("username")
	sid, err2 := r.Cookie("session")

	if err1 != nil || err2 != nil {
		p := &HomePage{Name: "King"}
		t, e := template.ParseFiles("./bin/template/login.html")

		if e != nil {
			log.Printf("Parsing template home.html error: %s", e)
			return 
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		t.Execute(w, p)

		return 
	}

	log.Println("cname: ", cname , " sid: ", sid)
	if len(cname.Value) != 0 && len(sid.Value) != 0 {
		http.Redirect(w, r, "/userhome", http.StatusFound)
		return 
	}
}

func userHomeHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	cname, err1 := r.Cookie("username")
	cooike, err2 := r.Cookie("session")

	log.Printf("userHomeHandler name:%s, cookie:%s", cname, cooike);

	if err1 != nil || err2 != nil {
		http.Redirect(w, r, "/", http.StatusFound)
		return 
	}

	fname := r.FormValue("username")

	videourl := "http://192.168.189.134:8080/statics/video/11111.mp4"
	hosturl := "192.168.189.134:8080"
	var p *UserPage
	if len(cname.Value) != 0 {
		p = &UserPage{Name: cname.Value, VideoUrl: videourl, HostUrl: hosturl}
	} else if len(fname) != 0 {
		p = &UserPage{Name: fname, HostUrl: hosturl}
	}

	t, e := template.ParseFiles("./bin/template/userhome.html")
	if e != nil {
		log.Printf("Parseing userhome.html error: %s", e)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	t.Execute(w, p)
}

// visitor
// user

func apiHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Println("Request Method ", r.Method)
	if r.Method != http.MethodPost {
		re, _ := json.Marshal(ErrorRequestNotRecognized)
		io.WriteString(w, string(re))
		return 
	}

	res, _ := ioutil.ReadAll(r.Body)
	apibody := &ApiBody{}

	//log.Println(string(res))
	if err := json.Unmarshal(res, apibody); err != nil {
		re, _ := json.Marshal(ErrorRequestBodyParseFailed)
		io.WriteString(w, string(re))
		return 
	}

	request(apibody, w, r)

	defer r.Body.Close()
}

func proxyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u, _ := url.Parse("http://127.0.0.1:9000/")
	proxy := httputil.NewSingleHostReverseProxy(u)
	proxy.ServeHTTP(w, r)
}
