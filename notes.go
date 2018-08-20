package main

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"os"
	"os/exec"
	"os/user"
)

var (
	default_dir = getUserHome() + "/notes"
	scratch_dir = default_dir + "/scratch"
	trash_dir   = default_dir + "/trash"
)

func scratchNote() {
	// Create note
	name := createNote("")

	// Open note
	openNote(name)

	// Toss note
	tossNote(name)
}

func createNote(name string) string {
	//TODO:  Check if note exists
	dir := default_dir

	// Name not supplied so create one
	if name == "" {
		name = genName()
		//dir = scratch_dir
	}

	f, err := os.Create(dir + "/" + name)
	check(err)

	defer f.Close()
	return name
}

// Generates a random name
func genName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	hasher := md5.New()
	hasher.Write(randBytes)
	return hex.EncodeToString(hasher.Sum(nil))
}

func openNote(name string) {
	editor := os.Getenv("EDITOR")

	cmd := exec.Command(editor, default_dir+"/"+name)

	//TODO:Print errors from command
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	check(err)

	// TODO: Deal with error
}

func tossNote(name string) {
	// Move file
	err := os.Rename(default_dir+"/"+name, trash_dir+"/"+name)
	if err != nil {
		panic(err)
	}
}

func deleteNote(name string) {
	err := os.Remove(default_dir + "/" + name)
	check(err)
}

func recoverNote(name string) {

}

func listNotes() {
	cmd := exec.Command("ls", default_dir)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	check(err)

}

//-----Helper Functions -------

func getUserHome() string {
	usr, err := user.Current()
	check(err)

	return usr.HomeDir
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
