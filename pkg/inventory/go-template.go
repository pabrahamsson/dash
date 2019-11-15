package inventory

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/template"
)


type GoTemplate struct {
	Template string               `yaml:"template"`
	Params map[string]interface{} `yaml:"params"`
}

func (gt *GoTemplate) Process(ns *string, r *Resource) error {

	// Pre-process to convert to absolute pathing
	abs, err := filepath.Abs(filepath.Join(r.Prefix, gt.Template))
	if err != nil {
		return err
	}
	gt.Template = abs

	fmt.Printf("Found template %s and one set of params\n", gt.Template)
	ext := filepath.Ext(gt.Template)
	basename := baseName(gt.Template, ext)
	err = processOneGoTemplate(gt.Template, gt.Params, r, basename + ".yaml")
	if err != nil {
		return err
	}

	return nil
}

func processOneGoTemplate(tpl string, ps map[string]interface{}, r *Resource, outName string) error {
	if outName == "" {
		outName = filepath.Base(tpl) + ".yaml"
	}

	t, err := template.ParseFiles(tpl)
	if err != nil {
		log.Fatal(err)
	}

	// write resulting resource to file
	outputDir := filepath.Join(r.Output, r.Action)
	out, err := os.Create(filepath.Join(outputDir, outName))
	if err != nil {
		return err
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()
	log.Printf("wrote %s\n", out.Name())
	    err = t.Execute(out, ps)
	if err != nil {
		return err
	}

	return nil
}

func baseName(filename string, extension string) string {
	basename := filepath.Base(filename[0:len(filename)-len(extension)])
	return basename
}
