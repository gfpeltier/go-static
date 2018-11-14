package internal

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
	"path/filepath"
	"gopkg.in/russross/blackfriday.v2"
)

const BuildPath = "build"
const PublicBuildPath = BuildPath + string(os.PathSeparator) + "public"
const PostSuffix = ".md"

func check(e error) {
	if e != nil {
		print("ERROR: Panicking\n")
		panic(e)
	}
}

func cleanBuild() {
	err := os.RemoveAll(BuildPath)
	check(err)
}

func moveAssets() {
	os.Mkdir(BuildPath, 0644)
	err := CopyDir("public", PublicBuildPath)
	check(err)
}

func buildIndex(theme string) {
	txt, err := ioutil.ReadFile("index.md")
	check(err)
	out := blackfriday.Run(txt)
	outFile := PublicBuildPath + string(os.PathSeparator)  + "index.html"
	err = ioutil.WriteFile(outFile, out, 0644)
	check(err)
}

func getPostPaths() ([]string, error){
	var paths []string
	err := filepath.Walk("posts", 
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, PostSuffix) {
				paths = append(paths, path)
			}
			return nil
	})
	return paths, err
}

func genOutPaths(inPaths []string) map[string]string {
	inOut := make(map[string]string)
	//for _, s := range inPaths {
	//	parts := strings.Split(s, string(os.PathSeparator))
		
	//}
	return inOut
}

func buildPosts(theme string) {
	paths, err := getPostPaths()
	check(err)
	for _, s := range paths {
		fmt.Println(s)
	}
}

func BuildSite(theme string) {
	fmt.Printf("Creating site with theme: %s\n", theme)
	cleanBuild()
	moveAssets()
	buildIndex(theme)
	buildPosts(theme)
}