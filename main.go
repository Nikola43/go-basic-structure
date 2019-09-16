package main




import (
	"github.com/nikola43/go_github/first_package"
	"github.com/nikola43/go_github/second_package"
	"github.com/nikola43/go_github/services"
)

func main() {
	services.CreateGithubRepository()
	services.FindGithubRepository()
	services.UpdateGithubRepository()

	first_package.GetUrl()
	second_package.GetUrl()
}
