package main

//go:generate go run . -o ./pkg/

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir_generate/genutil"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"golang.org/x/sync/semaphore"
)

var (
	output  string
	verbose bool
	listPkg bool
)

func init() {
	flag.StringVar(&output, "o", "", "output directory to mkdir in")
	flag.BoolVar(&verbose, "v", verbose, "log verbosely (debug mode)")
	flag.BoolVar(&listPkg, "l", listPkg, "only list packages and exit")
	flag.Parse()

	if !listPkg && output == "" {
		log.Fatalln("Missing -o output directory.")
	}

	if !verbose {
		verbose = os.Getenv("GIR_VERBOSE") == "1"
	}
}

func main() {
	var repos gir.Repositories

	// Load all of gotk4's packages first.
	genutil.MustAddPackages(&repos, gendata.Packages)
	// Get a map of exteral imports for packages that gotk4 already generates.
	overrides := genutil.LoadExternOverrides(gotk4Module, repos)

	// Add our own packages in.
	genutil.MustAddPackages(&repos, packages)
	// Dump the added packages down.
	genutil.PrintAddedPkgs(repos)

	if listPkg {
		return
	}

	strcases.AddPascalSpecials([]string{
		"Itp",
	})

	gen := girgen.NewGenerator(repos, genutil.ModulePath(webkitgtkModule, overrides))
	gen.Logger = log.New(os.Stderr, "girgen: ", log.Lmsgprefix)
	gen.AddFilters(gendata.Filters)
	gen.AddFilters(filters)
	gen.ApplyPreprocessors(preprocessors)
	gen.ApplyPreprocessors(gendata.Preprocessors)

	if verbose {
		gen.LogLevel = logger.Debug
	}

	if err := genutil.CleanDirectory(output, pkgExceptions); err != nil {
		log.Fatalln("failed to clean output directory:", err)
	}

	if errors := GeneratePackages(gen, output, packages); len(errors) > 0 {
		for _, err := range errors {
			log.Println("generation error:", err)
		}
		os.Exit(1)
	}

	if err := genutil.EnsureDirectory(output, pkgExceptions, pkgGenerated); err != nil {
		log.Fatalln("error verifying generation:", err)
	}
}

func GeneratePackages(gen *girgen.Generator, dst string, pkgs []gendata.Package) []error {
	sema := semaphore.NewWeighted(int64(runtime.GOMAXPROCS(-1)))

	var wg sync.WaitGroup
	defer wg.Wait()

	var errMut sync.Mutex
	var errors []error

	genNamespace := func(namespace *gir.Namespace) {
		ng := gen.UseNamespace(namespace.Name, namespace.Version)
		if ng == nil {
			log.Fatalln("cannot find namespace", namespace.Name, "v"+namespace.Version)
		}

		sema.Acquire(context.Background(), 1)
		wg.Add(1)

		go func() {
			if err := genutil.WriteNamespace(ng, dst); err != nil {
				errMut.Lock()
				errors = append(errors, err)
				errMut.Unlock()
			}

			sema.Release(1)
			wg.Done()
		}()
	}

	repos := gen.Repositories()

	for _, pkg := range pkgs {
		if pkg.Namespaces != nil {
			for _, wantedName := range pkg.Namespaces {
				namespace := repos.FindNamespace(wantedName)
				if namespace == nil {
					return []error{fmt.Errorf("namespace %q not found", wantedName)}
				}

				genNamespace(namespace.Namespace)
			}
		}

		repo := repos.FromPkg(pkg.PkgName)
		if repo == nil {
			return []error{fmt.Errorf("package %q not found", pkg.PkgName)}
		}

		for _, namespace := range repo.Namespaces {
			genNamespace(&namespace)
		}
	}

	wg.Wait()

	return errors
}
