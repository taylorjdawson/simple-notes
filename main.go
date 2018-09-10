package main

import (
	"fmt"
	"os"

	flag "github.com/ogier/pflag"
)

// flags
var (
	note_name string
	fg        string
	bool_flg  bool
	//new_note     string
	//scratch_note string
	//trash_note   string
)

// "main" is the entry point of our CLI app
func main() {

	flag.Parse()
	flag.Visit(func(f *flag.Flag) { fg = f.Shorthand })

	switch fg {
	case "n":
		createNote(note_name)
		fmt.Println("created")
	case "s":
		scratchNote()
		fmt.Println("s")
	case "d":
		tossNote(note_name)
		fmt.Println("d")
	case "D":
		deleteNote(note_name)
		fmt.Println("Deleted")
	case "r":
		fmt.Println("r")
	case "o":
		openNote(note_name)
	case "l":
		listNotes()
	default:
		listNotes()
	}
}

func init() {
	flag.StringVarP(&note_name, "new", "n", "", "Creates new note")
	flag.BoolVarP(&bool_flg, "scratch", "s", true, "Create a scratch note (automatically trashed when closed")
	flag.StringVarP(&note_name, "toss", "d", "", "Crumples note and tosses into trash")
	flag.StringVarP(&note_name, "delete", "D", "", "Permenantly deletes note")
	flag.StringVarP(&note_name, "recover", "r", "", "Recovers deleted note from trash bin")
	flag.StringVarP(&note_name, "open", "o", "", "Opens note with default text editor")
	flag.BoolVarP(&bool_flg, "list", "l", true, "Lists all notes")
}

// If no param supplied
func emptyInput(param string) {
	if param == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// printUsage is a custom function we created to print usage for our CLI app
func printUsage() {
	fmt.Printf("Usage: %s [options]\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(1)
}
