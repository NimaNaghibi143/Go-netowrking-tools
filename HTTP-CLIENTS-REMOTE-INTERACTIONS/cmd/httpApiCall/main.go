package main

r1, err := http.Get("http://www.google.com/robots.txt")
defer r1.Body.Close()

r2, err := http.Head("http://www.google.com/robots.txt")
defer r2.Body.Close()

form := url.Values{}
form.Add("foo", "bar")

r3, err := http.Post(
	"https://www.google.com/robots.txt",
	"application/x-www-form-urlencoded",
	strings.NewReader(form.Encode()w ),
)

// alternative approach:
// form := url.Values{}
// form.Add("foo", "bar")
// r3, err := http.PostForm("https://www.google.com/robots.txt", form)

// Read response body and close.
defer r3.Body.Close()