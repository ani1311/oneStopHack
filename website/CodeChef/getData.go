package CodeChef

import (
	"fmt"

	"../../models"
	"../../utils"
	"../../websiteData"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func GetChallanges() []models.Challenge {
	htm, _ := html.Parse(utils.GetPage("https://www.codechef.com/contests"))
	challenges := make([]models.Challenge, 0)
	challengesNode := utils.GetChallengeNodeUsingAtom(htm, atom.H3, "Present Contests")

	t := challengesNode.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling
	for t.DataAtom == atom.Tr {
		chal, err := GetChallenge(t)
		if err != nil {
			fmt.Println(err)
		} else {
			challenges = append(challenges, chal)
		}
		t = t.NextSibling
		if t.NextSibling != nil {
			t = t.NextSibling
		}
	}
	// for _, ch := range challenges {
	// 	fmt.Println(ch.Name + " : " + ch.Link + " | ")
	// }

	return challenges
}

func GetChallenge(n *html.Node) (models.Challenge, error) {
	var chal models.Challenge
	chal.Website = websiteData.CodeChef
	chal.Link = "https://www.codechef.com" + n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.Attr[0].Val
	chal.Name = n.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.NextSibling.FirstChild.Data

	return chal, nil
}
