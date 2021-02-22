package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"time"

	gogit "github.com/go-git/go-git"
	"github.com/go-git/go-git/plumbing/object"
)

// FileScore : Bug prediction score per file path
type FileScore struct {
	Path  string
	Score float64
}

// FixMessageJudger : Simple struct for testing regexp
type FixMessageJudger struct {
	regexpObj *regexp.Regexp
}

// NewFixMessageJudger : Constructor of FixMessageJudger
func NewFixMessageJudger(regexString string) *FixMessageJudger {
	fmj := new(FixMessageJudger)
	fmj.regexpObj = regexp.MustCompile(regexString)
	return fmj
}

// IsFixMessage : Returns whether the given string is a fix message
func (fmj *FixMessageJudger) IsFixMessage(s string) bool {
	return fmj.regexpObj.MatchString(s)
}

func main() {
	var (
		gitPath  = flag.String("path", "", "A path to Git repository")
		fixRegex = flag.String("regexp", "(?i)(^| )(fi(x|xed|xes)|clos(e|es|ed))", "A regexp specify bug fixes in commit message.(default: (?i)(.*(f|F)i(x|xed|xes).*|.*(c|C)los(e|es|ed)).*)")
		outPath  = flag.String("o", "./bugospots.csv", "An output of csv.")
	)
	flag.Parse()

	repo, err := gogit.PlainOpen(*gitPath)
	if err != nil {
		log.Fatal(err)
	}
	ref, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}
	head, err := repo.CommitObject(ref.Hash())
	if err != nil {
		log.Fatal(err)
	}
	cIter, err := repo.Log(&gogit.LogOptions{From: head.Hash})
	if err != nil {
		log.Fatal(err)
	}

	var bugFixes []object.Commit
	fmj := NewFixMessageJudger(*fixRegex)

	cIter.ForEach(func(c *object.Commit) error {
		if fmj.IsFixMessage(c.Message) {
			bugFixes = append(bugFixes, *c)
		}
		return nil
	})

	if len(bugFixes) == 0 {
		log.Println("No bugfixes")
		return
	}

	sort.Slice(bugFixes, func(i, j int) bool {
		return bugFixes[i].Author.When.Unix() < bugFixes[j].Author.When.Unix()
	})

	oldestFixTime := bugFixes[0].Author.When.Unix()
	currentTime := time.Now().Unix()
	log.Println("oldest bug fix:", bugFixes[0].Author.When.Local())
	log.Println("latest bug fix:", bugFixes[len(bugFixes)-1].Author.When.Local())
	log.Println("current:", time.Now().Local())
	log.Println("bug fixes:", len(bugFixes))
	fileScoreMap := make(map[string]float64)
	for _, b := range bugFixes[1:] {
		prev, err := repo.CommitObject(b.ParentHashes[0])
		if err != nil {
			continue
		}
		bTree, err := b.Tree()
		if err != nil {
			continue
		}
		prevTree, err := prev.Tree()
		if err != nil {
			continue
		}
		patch, err := bTree.Patch(prevTree)
		if err != nil {
			continue
		}
		for _, fileStat := range patch.Stats() {
			t := float64(1) - float64(currentTime-b.Author.When.Unix())/float64(currentTime-oldestFixTime)
			if t < 0 {
				log.Fatal(t)
			}
			fileScoreMap[fileStat.Name] += 1 / (1 + math.Exp(-12*t+12))
		}
	}

	var fileScoreArray []FileScore
	for key, value := range fileScoreMap {
		var fs FileScore
		fs.Path = key
		fs.Score = value
		fileScoreArray = append(fileScoreArray, fs)
	}

	sort.Slice(fileScoreArray, func(i, j int) bool {
		return fileScoreArray[i].Score > fileScoreArray[j].Score
	})

	for i, fs := range fileScoreArray {
		fmt.Println(fmt.Sprint(fs.Score) + "," + fs.Path)
		if i >= 10 {
			break
		}
	}

	outFile, _ := os.Create(*outPath)
	defer outFile.Close()
	w := bufio.NewWriter(outFile)
	for _, fs := range fileScoreArray {
		line := fmt.Sprint(fs.Score) + "," + fs.Path + "\n"
		w.WriteString(line)
		w.Flush()
	}
}
