package skeleton

import (
	"bytes"
	"embed"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"text/template"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
	"github.com/MueR/adventofcode.go/util"
	"github.com/ericchiang/css"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/html"
	_ "openticket.tech/log/v2"
)

//go:embed tmpls/*.go
var fs embed.FS

type puzzle struct {
	Samples      []string
	Input        []byte
	Descriptions [][]byte
}

// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatal().Msgf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatal().Msgf("year is before 2015: %d", year)
	}

	ts, err := template.ParseFS(fs, "tmpls/*.go")
	if err != nil {
		log.Fatal().Msgf("parsing tmpls directory: %s", err)
	}

	mainFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/main.go", year, day))
	testFilename := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/main_test.go", year, day))
	//inputFileName := filepath.Join(util.Dirname(), "../../", fmt.Sprintf("%d/day%02d/input.txt", year, day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatal().Msgf("making directory: %s", err)
	}

	//ensureNotOverwriting(mainFilename)
	//ensureNotOverwriting(testFilename)
	//ensureNotOverwriting(inputFileName)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatal().Msgf("creating main.go file: %v", err)
	}
	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatal().Msgf("creating main_test.go file: %v", err)
	}

	p, err := getPuzzle(filepath.Join(util.Dirname(), "../../"), day, year)
	if err != nil {
		log.Fatal().Msgf("getting puzzle: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	ts.ExecuteTemplate(testFile, "main_test.go", p)

	log.Info().Msgf("templates made for %d-day%d\n", year, day)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatal().Msgf("File already exists: %s", filename)
	}
}

func (p *puzzle) getInput(day, year int) (err error) {
	token, ok := os.LookupEnv("AOC_SESSION_TOKEN")
	if !ok {
		log.Fatal().Msgf("AOC_SESSION_TOKEN environment variable not set")
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)
	if err != nil {
		log.Fatal().Msgf("creating request: %v", err)
		return
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: token})
	resp, err := http.DefaultClient.Do(req)
	log.Err(err).Msgf("getting input")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal().Msgf("getting input: %s", resp.Status)
	}

	p.Input, err = io.ReadAll(resp.Body)
	return
}

func (p *puzzle) getDescription(day, year int) (err error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d", year, day), nil)
	log.Info().Err(err).Msgf("creating request to https://adventofcode.com/%d/day/%d", year, day)
	if err != nil {
		log.Err(err).Msg("creating request")
		return
	}
	resp, err := http.DefaultClient.Do(req)
	log.Err(err).Msg("getting description")
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Info().Msgf("getting description: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	sel, err := css.Parse("article.day-desc")
	if err != nil {
		log.Err(err).Msg("parsing selector")
	}
	node, err := html.Parse(bytes.NewReader(body))
	if err != nil {
		log.Err(err).Msg("parsing html")
	}
	el := sel.Select(node)
	if len(el) == 0 {
		log.Error().Msg("no description found")
		return
	}
	desc := []byte{}
	w := bytes.NewBuffer(desc)
	for _, em := range el {
		err = html.Render(w, em)
		if err != nil {
			log.Error().Msg("could not render description")
			return
		}
		p.Descriptions = append(p.Descriptions, w.Bytes())
	}

	sel, err = css.Parse("article.day-desc pre code")
	if err != nil {
		log.Err(err).Msg("failed build selector parsing")
	}
	for _, n := range sel.Select(node) {
		var b bytes.Buffer
		html.Render(&b, n)
		s := b.String()
		p.Samples = append(p.Samples, s[6:len(s)-7])
	}
	return
}

func getPuzzle(basePath string, day, year int) (p puzzle, err error) {
	p = puzzle{}
	err = p.getInput(day, year)
	if err != nil {
		return
	}
	n := filepath.Join(basePath, fmt.Sprintf("%d/day%02d/input.txt", year, day))
	err = os.WriteFile(n, p.Input, 0644)
	log.Info().Err(err).Msg("writing input")
	if err != nil {
		return
	}

	err = p.getDescription(day, year)
	if err != nil {
		return
	}
	md := ""
	for _, d := range p.Descriptions {
		pmd, e := htmltomarkdown.ConvertString(string(d))
		log.Info().Err(err).Msg("converting html to markdown")
		if e != nil {
			return
		}
		md += pmd + "\n\n"
	}

	n = filepath.Join(basePath, fmt.Sprintf("%d/day%02d/README.md", year, day))
	err = os.WriteFile(n, []byte(md), 0644)
	log.Info().Err(err).Msg("writing description to readme")
	if err != nil {
		return
	}

	return
}
