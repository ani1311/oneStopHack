package websiteData

type Website string

const (
	CodeChef    Website = "codeChef"
	HackerEarth Website = "hackerEarth"
)

var Websites = []Website{
	CodeChef,
	HackerEarth,
}
