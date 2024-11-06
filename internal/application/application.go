package application

import (
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
	"time"
	"ytlog/internal/config"
	"ytlog/internal/persistence"
	"ytlog/internal/ytracker"
)

type Application struct {
	cfg *config.Config
	yt  *ytracker.YTracker
	db  *persistence.Database
}

func NewApplication(cfg *config.Config) *Application {
	db := persistence.NewDatabase(cfg)
	err := db.CreateTables()
	if err != nil {
		panic(err.Error())
	}

	yt := ytracker.CreateYTracker(cfg)
	return &Application{
		cfg: cfg,
		yt:  yt,
		db:  db,
	}
}

func (app *Application) loadUsers() {
	ytUsers := app.yt.GetUsers()
	users := persistence.MapUsers(ytUsers)
	app.db.SaveUsers(users)
}

func (app *Application) loadIssues(lastSync time.Time) *[]persistence.Issue {
	ytIssues := app.yt.GetIssues(lastSync)
	issues := persistence.MapIssues(ytIssues)
	app.db.SaveIssues(issues)
	return issues
}

func (app *Application) loadIssueLog(issueKey string) {
	ytIssueLog := app.yt.GetIssueChangelog(issueKey)
	issueLogRecords := persistence.MapIssueLog(ytIssueLog)
	app.db.SaveIssueLog(issueLogRecords)
}

func (app *Application) Run() {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true))

	bar.Describe("[cyan][1/3][reset] Load [red]users[reset]")
	app.loadUsers()

	bar.Describe("[cyan][2/3][reset] Load [red]issues[reset]")
	lastSync := app.db.GetLastSync()
	issues := app.loadIssues(lastSync.UpdatedAt)
	bar = progressbar.NewOptions(
		len(*issues),
		progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("[cyan][3/3][reset] Load [red]changes log[reset]"),
	)
	for _, issue := range *issues {
		_ = bar.Add(1)
		app.loadIssueLog(issue.Key)
	}
	app.db.UpdateSync()
}
