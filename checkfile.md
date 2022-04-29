
client := github.NewClient(nil)

opt := &github.RepositoryListByOrgOptions{
	ListOptions: github.ListOptions{PerPage: 10},
}
// get all pages of results
var allRepos []*github.Repository
for {
	repos, resp, err := client.Repositories.ListByOrg(ctx, "github", opt)
	if err != nil {
		return err
	}
	allRepos = append(allRepos, repos...)
	if resp.NextPage == 0 {
		break
	}
	opt.Page = resp.NextPage
}
