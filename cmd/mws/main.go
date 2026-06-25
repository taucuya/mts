package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"mws/internal/profile"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(1)
	}

	switch os.Args[1] {
	case "profile":
		handleProfile(os.Args[2:])
	case "help", "--help", "-h":
		printHelp()
	default:
		fmt.Printf("unknown command: %s\n\n", os.Args[1])
		printHelp()
		os.Exit(1)
	}
}

func handleProfile(args []string) {
	if len(args) < 1 {
		printHelp()
		os.Exit(1)
	}

	switch args[0] {
	case "create":
		handleProfileCreate(args[1:])
	case "get":
		handleProfileGet(args[1:])
	case "list":
		handleProfileList(args[1:])
	case "delete":
		handleProfileDelete(args[1:])
	case "help":
		printHelp()
	default:
		fmt.Printf("unknown profile command: %s\n\n", args[0])
		printHelp()
		os.Exit(1)
	}
}

func handleProfileCreate(args []string) {
	fs := flag.NewFlagSet("profile create", flag.ExitOnError)
	name := fs.String("name", "", "profile name")
	user := fs.String("user", "", "user name")
	project := fs.String("project", "", "project name")
	fs.Parse(args)

	if *name == "" || *user == "" || *project == "" {
		fmt.Println("flags --name, --user, --project are required")
		os.Exit(1)
	}

	if err := profile.Create(*name, *user, *project); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("profile created")
}

func handleProfileGet(args []string) {
	fs := flag.NewFlagSet("profile get", flag.ExitOnError)
	name := fs.String("name", "", "profile name")
	fs.Parse(args)

	if *name == "" {
		fmt.Println("flag --name is required")
		os.Exit(1)
	}

	p, err := profile.Get(*name)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Printf("name: %s\n", *name)
	fmt.Printf("user: %s\n", p.User)
	fmt.Printf("project: %s\n", p.Project)
}

func handleProfileList(args []string) {
	fs := flag.NewFlagSet("profile list", flag.ExitOnError)
	fs.Parse(args)

	list, err := profile.List()
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if len(list) == 0 {
		fmt.Println("no profiles found")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "NAME\tUSER\tPROJECT")
	for _, p := range list {
		fmt.Fprintf(w, "%s\t%s\t%s\n", p.Name, p.User, p.Project)
	}
	w.Flush()
}

func handleProfileDelete(args []string) {
	fs := flag.NewFlagSet("profile delete", flag.ExitOnError)
	name := fs.String("name", "", "profile name")
	fs.Parse(args)

	if *name == "" {
		fmt.Println("flag --name is required")
		os.Exit(1)
	}

	if err := profile.Delete(*name); err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	fmt.Println("profile deleted")
}

func printHelp() {
	fmt.Println(`Usage:
  mws profile create --name=<name> --user=<user> --project=<project>
  mws profile get --name=<name>
  mws profile list
  mws profile delete --name=<name>
  mws help

Commands:
  profile create   Create a profile in current directory
  profile get      Show profile by name
  profile list     List profiles in current directory
  profile delete   Delete profile by name
  help             Show this help`)
}
