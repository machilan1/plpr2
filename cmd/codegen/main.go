package main

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"runtime/debug"
	"strings"
	"text/template"
)

//go:embed templates
var templates embed.FS

type config struct {
	withPagination bool
}

func main() {
	flag.Parse()

	name := flag.Arg(0)
	if name == "" {
		fmt.Println("missing domain name")
		os.Exit(1)
	}

	abbr := flag.Arg(1)
	if abbr == "" {
		fmt.Println("missing abbr name")
		os.Exit(1)
	}

	plural := flag.Arg(2)
	if plural == "" {
		fmt.Println("missing plural name")
		os.Exit(1)
	}

	pagArg := flag.Arg(3)
	pag := true
	if pagArg == "n" || pagArg == "N" {
		pag = false
	}

	cfg := config{
		withPagination: pag,
	}

	if err := run(name, abbr, plural, cfg); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(name string, abbr string, plural string, cfg config) error {
	info, _ := debug.ReadBuildInfo()
	mod := info.Main.Path

	if err := addAppLayer(mod, name, abbr, plural, cfg); err != nil {
		return fmt.Errorf("adding app layer files: %w", err)
	}

	if err := addBusinessLayer(mod, name, abbr, plural, cfg); err != nil {
		return fmt.Errorf("adding bus layer files: %w", err)
	}

	if err := addStorageLayer(mod, name, abbr, plural, cfg); err != nil {
		return fmt.Errorf("adding sto layer files: %w", err)
	}

	if err := fmtCode(); err != nil {
		return fmt.Errorf("formatting code: %w", err)
	}

	fmt.Println("Done")
	return nil
}

func addAppLayer(mod, domain string, abbr string, plural string, cfg config) error {
	const basePath = "internal/app/domain"

	app, err := fs.Sub(templates, "templates/app")
	if err != nil {
		return fmt.Errorf("switching to template/app folder: %w", err)
	}

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		return walkWork(mod, domain, abbr, plural, basePath, app, fileName, dirEntry, err, cfg)
	}

	fmt.Println("=======================================================")
	fmt.Println("APP LAYER CODE")

	if err := fs.WalkDir(app, ".", fn); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}

func addBusinessLayer(mod, domain string, abbr string, plural string, cfg config) error {
	const basePath = "internal/business/domain"

	app, err := fs.Sub(templates, "templates/business")
	if err != nil {
		return fmt.Errorf("switching to template/business folder: %w", err)
	}

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		return walkWork(mod, domain, abbr, plural, basePath, app, fileName, dirEntry, err, cfg)
	}

	fmt.Println("=======================================================")
	fmt.Println("BUSINESS LAYER CODE")

	if err := fs.WalkDir(app, ".", fn); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}

func addStorageLayer(mod, domain string, abbr string, plural string, cfg config) error {
	basePath := fmt.Sprintf("internal/business/domain/%s/stores", domain)

	app, err := fs.Sub(templates, "templates/storage")
	if err != nil {
		return fmt.Errorf("switching to template/storage folder: %w", err)
	}

	fn := func(fileName string, dirEntry fs.DirEntry, err error) error {
		return walkWork(mod, domain, abbr, plural, basePath, app, fileName, dirEntry, err, cfg)
	}

	fmt.Println("=======================================================")
	fmt.Println("STORAGE LAYER CODE")

	if err := fs.WalkDir(app, ".", fn); err != nil {
		return fmt.Errorf("walking directory: %w", err)
	}

	return nil
}

func walkWork(mod string, domain string, abbr string, plural string, basePath string, app fs.FS, fileName string, dirEntry fs.DirEntry, err error, cfg config) error {
	if err != nil {
		return fmt.Errorf("walkdir failure: %w", err)
	}

	if dirEntry.IsDir() {
		return nil
	}

	f, err := app.Open(fileName)
	if err != nil {
		return fmt.Errorf("opening %s: %w", fileName, err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("reading %s: %w", fileName, err)
	}

	tmpl := template.Must(template.New("code").Parse(string(data)))

	domainVar := abbr
	domainPlural := plural
	d := struct {
		Module        string
		DomainL       string
		DomainU       string
		DomainVar     string
		DomainVarU    string
		DomainVars    string
		DomainVarsU   string
		DomainNewVar  string
		DomainUpdVar  string
		DomainPlural  string
		DomainPluralU string
		// Options
		WithPagination bool
	}{
		Module:        mod,
		DomainL:       strings.ToLower(domain),
		DomainU:       strings.ToUpper(domain[0:1]) + strings.ToLower(domain[1:]),
		DomainVar:     strings.ToLower(domainVar),
		DomainVarU:    strings.ToUpper(domainVar[0:1]) + strings.ToLower(domainVar[1:]),
		DomainVars:    strings.ToLower(domainVar) + "s",
		DomainVarsU:   strings.ToUpper(domainVar[0:1]) + strings.ToLower(domainVar[1:]) + "s",
		DomainNewVar:  "n" + strings.ToUpper(domainVar[0:1]) + strings.ToLower(domainVar[1:]),
		DomainUpdVar:  "u" + strings.ToUpper(domainVar[0:1]) + strings.ToLower(domainVar[1:]),
		DomainPlural:  strings.ToLower(domainPlural),
		DomainPluralU: strings.ToUpper(domainPlural[0:1]) + strings.ToLower(domainPlural[1:]),
		// Options
		WithPagination: cfg.withPagination,
	}

	var b bytes.Buffer
	if err := tmpl.Execute(&b, d); err != nil {
		return err
	}

	if err := writeFile(basePath, domain, fileName, b); err != nil {
		return fmt.Errorf("writing %s: %w", fileName, err)
	}

	return nil
}

func writeFile(basePath string, domain string, fileName string, b bytes.Buffer) error {
	path := basePath

	parts := strings.SplitN(basePath, "/", 3)
	switch {
	case strings.HasSuffix(basePath, "stores"):
		path = fmt.Sprintf("%s/%sdb", basePath, domain)
	case parts[1] == "app":
		path = fmt.Sprintf("%s/%sapi", basePath, domain)
	case parts[1] == "business":
		path = fmt.Sprintf("%s/%s", basePath, domain)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Creating directory:", path)

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return fmt.Errorf("write app directory: %w", err)
		}
	}

	// Remove the suffix `.tmpl` from the file name.
	path = fmt.Sprintf("%s/%s", path, fileName[:len(fileName)-5])
	path = strings.Replace(path, "new", domain, 1)

	fmt.Println("Add file:", path)
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}
	defer f.Close()

	fmt.Println("Writing code:", path)
	if _, err := f.Write(b.Bytes()); err != nil {
		return fmt.Errorf("writing bytes: %w", err)
	}

	return nil
}

func fmtCode() error {
	fmt.Println("=======================================================")
	fmt.Println("Formatting code...")
	if err := exec.Command("make", "fmt").Run(); err != nil {
		fmt.Println("command: make fmt: ", err)
	}

	return nil
}
