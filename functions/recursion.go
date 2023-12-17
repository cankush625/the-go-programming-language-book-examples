package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func findLinksInHTML() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findLinksInHTML: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// Note:
// Many programming language implementation use fixed-size
// function call stack; sizes from 64KB to 2MB are typical.
// Fixed size stacks impose a limit on the depth of recursion,
// so one must be careful to avoid stack overflow when traversing
// large data structures recursively.
// In contrast, typical Go implementation use variable-size stacks
// that start small and grow as needed upto a limit on the order of
// a Gigabyte. This lets us use recursion safely and without worrying
// about overflow.

func main() {
	findLinksInHTML()
}

// Input:
// curl https://go.dev/ | go run recursion.go
// Output:
// /
// #main-content
// #
// /solutions/case-studies
// /solutions/use-cases
// /security/
// /learn/
// #
// /doc/effective_go
// /doc
// https://pkg.go.dev/std
// /doc/devel/release
// https://pkg.go.dev
// #
// /talks/
// https://www.meetup.com/pro/go
// https://github.com/golang/go/wiki/Conferences
// /blog
// /help
// https://groups.google.com/g/golang-nuts
// https://github.com/golang
// https://twitter.com/golang
// https://www.reddit.com/r/golang/
// https://invite.slack.golangbridge.org/
// https://stackoverflow.com/tags/go
// /
// #
// #
// /solutions/case-studies
// /solutions/use-cases
// /security/
// /learn/
// #
// #
// /doc/effective_go
// /doc
// https://pkg.go.dev/std
// /doc/devel/release
// https://pkg.go.dev
// #
// #
// /talks/
// https://www.meetup.com/pro/go
// https://github.com/golang/go/wiki/Conferences
// /blog
// /help
// https://groups.google.com/g/golang-nuts
// https://github.com/golang
// https://twitter.com/golang
// https://www.reddit.com/r/golang/
// https://invite.slack.golangbridge.org/
// https://stackoverflow.com/tags/go
// /learn/
// /dl
// /dl/
// /dl
// /solutions/
// /solutions/google/
// /solutions/paypal
// /solutions/americanexpress
// /solutions/mercadolibre
// https://bitly.com/blog/why-we-write-everything-in-go/?utm_source=go-dev&utm_medium=referral&utm_campaign=go-dev&utm_content=case-study
// https://medium.com/capital-one-tech/a-serverless-and-go-journey-credit-offers-api-74ef1f9fde7f
// https://www.cockroachlabs.com/blog/why-go-was-the-right-choice-for-cockroachdb/
// https://blogs.dropbox.com/tech/2014/07/open-sourcing-our-go-libraries/
// https://blog.cloudflare.com/graceful-upgrades-in-go/
// https://entgo.io/blog/2019/10/03/introducing-ent/
// https://cloudblogs.microsoft.com/opensource/2018/02/21/go-lang-brian-ketelsen-explains-fast-growth/
// https://medium.com/tech-at-wildlife-studios/pitaya-wildlifes-golang-go-af57865f7a11
// https://medium.com/netflix-techblog/application-data-caching-using-ssds-5bf25df851ef
// https://technology.riotgames.com/news/leveraging-golang-game-development-and-operations
//  https://www.zdnet.com/article/salesforce-why-we-ditched-python-for-googles-go-language-in-einstein-analytics/
// https://blog.twitch.tv/en/2016/07/05/gos-march-to-low-latency-gc-a6fa96f06eb7/
// https://blog.twitter.com/engineering/en_us/a/2015/handling-five-billion-sessions-a-day-in-real-time.html
// https://eng.uber.com/aresdb/
// /tour/
// https://cloud.google.com/go/home
// https://aws.amazon.com/sdk-for-go/
// https://github.com/Azure/azure-sdk-for-go
// /solutions/cloud/
// https://github.com/spf13/cobra
// https://github.com/spf13/viper
// https://github.com/urfave/cli
// https://github.com/go-delve/delve
// https://github.com/chzyer/readline
// /solutions/clis/
// /pkg/net/http/
// /pkg/html/template/
// https://github.com/flosch/pongo2
// /pkg/database/sql/
// https://github.com/elastic/go-elasticsearch
// /solutions/webdev/
// https://github.com/open-telemetry/opentelemetry-go
// https://github.com/istio/istio
// https://github.com/urfave/cli
// /solutions/devops/
// /solutions/use-cases
// /learn/
// /doc/install/
// /learn#guided-learning-journeys
// /learn#online-learning
// /learn#featured-books
// /learn#self-paced-labs
// https://www.ardanlabs.com/
// https://www.gopherguides.com/
// https://bosssauce.it/services/training
// https://github.com/shijuvar/gokit/tree/master/training
// /solutions/
// /solutions/use-cases
// /solutions/case-studies
// /learn/
// /play
// /tour/
// https://stackoverflow.com/questions/tagged/go?tab=Newest
// /help/
// https://pkg.go.dev
// /pkg/
// https://pkg.go.dev/about
// /project
// /dl/
// /blog/
// https://github.com/golang/go/issues
// /doc/devel/release
// /brand
// /conduct
// https://www.twitter.com/golang
// https://www.twitter.com/golang
// https://github.com/golang
// https://invite.slack.golangbridge.org/
// https://reddit.com/r/golang
// https://www.meetup.com/pro/go
// https://golangweekly.com/
// /copyright
// /tos
// http://www.google.com/intl/en/policies/privacy/
// /s/website-issue
// https://google.com
// https://policies.google.com/technologies/cookies
