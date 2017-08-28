// Copyright (c) ThoughtWorks Inc. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for details.

package main

import (
	"archive/zip"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"text/template"

	gh "github.com/google/go-github/github"
)

const (
	getgaugeOrg        = "getgauge"
	gaugeRepo          = "gauge"
	linuxAssetName     = "linux.x86_64.zip"
	tmpFolder          = "tmp"
	gaugeBinaryName    = "gauge"
	downloadedFileName = "gauge.zip"
	dockerfileTemplate = "Dockerfile.template"
	dockerfile         = "Dockerfile"
)

func main() {
	defer clean()
	downloadLatestGauge()
	var p map[string]map[string]string
	err := json.Unmarshal(platformImages(), &p)
	check(err)
	images := make(map[string]*labels, 0)
	gv := latestVersion(gaugeRepo)

	for platform, v := range p {
		for lang, iTag := range v {
			images[filepath.Join("Dockerfiles", lang, platform)] = &labels{
				GaugeVersion:  gv,
				PluginVersion: latestVersion(fmt.Sprintf("gauge-%s", lang)),
				ImageTag:      iTag}
		}
	}
	buildImages(images)
}

func platformImages() []byte {
	if runtime.GOOS == "windows" {
		return []byte(`
		{
			"linux":{
				"csharp":"getgauge/gauge-mono48-centos7",
				"java":"getgauge/gauge-jdk8-centos7",
				"ruby":"getgauge/gauge-ruby23-centos7"
			},
			"windows":{
				"csharp":"",
				"java":"",
				"ruby":""		
			}
		}
		`)
	}
	return []byte(`
	{
		"linux":{
			"csharp":"getgauge/gauge-mono48-centos7",
			"java":"getgauge/gauge-jdk8-centos7",
			"ruby":"getgauge/gauge-ruby23-centos7"
		}
	}
	`)
}

func downloadLatestGauge() {
	err := os.MkdirAll(tmpFolder, 0755)
	check(err)
	rel := latestRelease(gaugeRepo)
	downloadedFile := filepath.Join(tmpFolder, downloadedFileName)
	for _, a := range rel.Assets {
		if strings.HasSuffix(a.GetName(), linuxAssetName) {
			download(a.GetBrowserDownloadURL(), downloadedFile)
		}
	}
	r, err := zip.OpenReader(downloadedFile)
	check(err)
	defer r.Close()

	for _, f := range r.File {
		if strings.HasSuffix(f.Name, gaugeBinaryName) {
			log.Print("extracting gauge")
			af, err := f.Open()
			defer af.Close()
			check(err)
			dest, err := os.OpenFile(filepath.Join(tmpFolder, gaugeBinaryName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			defer dest.Close()
			check(err)
			io.Copy(dest, af)
		}
	}
}

func latestVersion(repo string) string {
	tagname := latestRelease(repo).GetTagName()
	return tagname[1:len(tagname)]
}

func latestRelease(repo string) *gh.RepositoryRelease {
	c := gh.NewClient(nil)
	rel, _, err := c.Repositories.GetLatestRelease(context.Background(), getgaugeOrg, repo)
	check(err)
	return rel
}

func download(url, filename string) {
	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	io.Copy(f, resp.Body)
}

type labels struct {
	GaugeVersion  string
	PluginVersion string
	ImageTag      string
}

func buildImages(m map[string]*labels) {
	wg := &sync.WaitGroup{}
	for k, v := range m {
		t, err := template.ParseFiles(filepath.Join(k, dockerfileTemplate))
		if err != nil {
			log.Fatal(err)
		}
		df := filepath.Join(k, dockerfile)
		f, err := os.OpenFile(df, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)

		if err != nil {
			log.Fatal(err)
		}
		err = t.Execute(f, v)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
		defer func() {
			err := os.Remove(df)
			check(err)
		}()
		wg.Add(1)
		go buildImage(k, v.ImageTag, wg)
	}
	wg.Wait()
}

type labelledWriter struct {
	label string
}

func (l *labelledWriter) Write(p []byte) (n int, err error) {
	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	return f.Write([]byte(fmt.Sprintf("[%s] %s", l.label, p)))
}

func buildImage(p, t string, wg *sync.WaitGroup) {
	// invoke docker command instead of using the Go SDK
	// the below code is less complex than the docker API.
	// https://github.com/moby/moby/issues/27186#issuecomment-252264177
	// imgPlatform, err := filepath.Rel("Dockerfiles", p)
	// check(err)
	// logger := &labelledWriter{label: imgPlatform}
	cmd := exec.Command("docker", "build", "-t", t, "-f", filepath.Join(p, dockerfile), ".", "--force-rm", "--no-cache")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Building : %s", t)
	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Build finished with error: %v", err)
	}
	log.Print("Build completed")
	wg.Done()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func clean() {
	err := os.RemoveAll(tmpFolder)
	check(err)
}
