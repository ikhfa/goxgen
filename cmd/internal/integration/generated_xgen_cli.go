// Code generated by github.com/goxgen/goxgen, DO NOT EDIT.

package main

import(
    "github.com/goxgen/goxgen/plugins/cli/app"
    "github.com/goxgen/goxgen/plugins/cli/project"
    "github.com/goxgen/goxgen/cmd/internal/integration/gorm_advanced"
    gorm_advancedServer "github.com/goxgen/goxgen/cmd/internal/integration/gorm_advanced/generated/server"
    "github.com/goxgen/goxgen/cmd/internal/integration/gorm_example"
    gorm_exampleServer "github.com/goxgen/goxgen/cmd/internal/integration/gorm_example/generated/server"
    "github.com/goxgen/goxgen/cmd/internal/integration/myproject"
    myprojectServer "github.com/goxgen/goxgen/cmd/internal/integration/myproject/generated/server"
)

var(
    settings = &app.Settings{
        Projects: project.NewProjectList(
            &project.CliProject{
                Name: "gorm_advanced",
                Server: gorm_advancedServer.Server,
                TestsFS: gorm_advanced.TestsFS,
                TestDir: "tests",
            },
            &project.CliProject{
                Name: "gorm_example",
                Server: gorm_exampleServer.Server,
                TestsFS: gorm_example.TestsFS,
                TestDir: "tests",
            },
            &project.CliProject{
                Name: "myproject",
                Server: myprojectServer.Server,
                TestsFS: myproject.TestsFS,
                TestDir: "tests",
            },
        ),
    }
)

func main() {
    app.Run(settings)
}


