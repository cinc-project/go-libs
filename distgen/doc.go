/*

A simple generator for creating easily distributable Go packages.

This automation provides an easy way to generate variables that can be configured
globally across multiple go packages, such as trademarks, product names, websites,
etc. This generator should be defined as a `go:generate` comment and run at build
time, using the `go generate` command.

Single Package Example

Our first example involves a simple use within a single `main` package. First, create
a file called `dist_gen.go` with the following command:

  	package main
  	//go:generate go run github.com/chef/go-libs/distgen

The automation will deploy a file called `dist.go` with all the variables defined
inside the JSON file `glob_dist.json` inside this repository. (See a real example
at [example-main/](example-main).)

Multi-package Example

In our second example involving multi-package, create a go package called `dist`
with a file called `gen.go` with the following command:

  	package dist
  	//go:generate go run github.com/chef/go-libs/distgen global.go dist

This usage is for go projects that has multiple packages. By creating a single `dist`
package inside your repository, you can import the generated package in any other
packages. (See a real example at [example-multi-pkg/](example-multi-pkg).)

Custom JSON File Example

To fully customize this automation, a user can provide a URL pointing to a custom JSON
file as a third parameter of the `go:generate` directive. This custom JSON file should contain the
global variables to generate. (See an example of a JSON file at
[glob_dist.json](glob_dist.json).)

  	package dist
  	//go:generate go run github.com/chef/go-libs/distgen global.go dist https://example.com/path/to/glob_dist.json

*/

package main // import "github.com/chef/go-libs/distgen"
