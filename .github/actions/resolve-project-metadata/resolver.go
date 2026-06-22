package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type dockerImage struct {
	Name       string `json:"name,omitempty"`
	Registry   string `json:"registry,omitempty"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
}

type metadata struct {
	Version          string        `json:"version"`
	JavaImages       []dockerImage `json:"java_images"`
	ThirdPartyImages []dockerImage `json:"third_party_images"`
}

func main() {
	versionFile := flag.String("version-file", "version.properties", "Relative path to version.properties")
	githubOutputFile := flag.String("github-output-file", "", "Path to GITHUB_OUTPUT")
	flag.Parse()

	version, err := computeVersion(*versionFile)
	// TODO remove hardcoded value
	fmt.Println("WARNING: Using hardcoded version value for offline build")
	version = "latest"
	fatalIfErr(err)

	out := metadata{
		Version: version,
		JavaImages: []dockerImage{
			{
				Name:       "iot-collector",
				Registry:   "docker.io",
				Repository: "dgs19/iot-collector",
				Tag:        version,
			},
			//{
			//	Name:       "iot-collector-ui",
			//	Registry:   "docker.io",
			//	Repository: "dgs19/iot-collector-ui",
			//	Tag:        version,
			//},
		},
		ThirdPartyImages: []dockerImage{
			//{
			//	Name:       "traefik",
			//	Registry:   "docker.io",
			//	Repository: "traefik",
			//	Tag:        "v3.7.5",
			//},
		},
	}

	if strings.TrimSpace(*githubOutputFile) != "" {
		fatalIfErr(writeGithubOutputs(*githubOutputFile, out))
		return
	}

	jsonBytes, err := json.MarshalIndent(out, "", "  ")
	fatalIfErr(err)
	fmt.Println(string(jsonBytes))
}

func fatalIfErr(err error) {
	if err == nil {
		return
	}
	_, _ = fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func parsePropertiesFile(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open properties file %s: %w", path, err)
	}

	props := make(map[string]string)
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := strings.TrimSpace(s.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		idx := strings.Index(line, "=")
		if idx < 0 {
			continue
		}
		k := strings.TrimSpace(line[:idx])
		v := strings.TrimSpace(line[idx+1:])
		props[k] = v
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("scan properties file %s: %w", path, err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("close properties file %s: %w", path, err)
	}
	return props, nil
}

func computeVersion(versionPath string) (string, error) {
	props, err := parsePropertiesFile(versionPath)
	if err != nil {
		return "", err
	}

	major := props["MAJOR_VERSION"]
	minor := props["MINOR_VERSION"]
	patch := props["PATCH_VERSION"]
	build := props["BUILD_NUMBER"]
	if major == "" || minor == "" || patch == "" || build == "" {
		return "", errors.New("version.properties is missing MAJOR_VERSION, MINOR_VERSION, PATCH_VERSION, or BUILD_NUMBER")
	}
	return fmt.Sprintf("%s.%s.%s-B%s", major, minor, patch, build), nil
}

func writeGithubOutputs(outputPath string, m metadata) (err error) {
	javaJSON, err := json.Marshal(m.JavaImages)
	if err != nil {
		return fmt.Errorf("marshal java_images: %w", err)
	}
	thirdPartyJSON, err := json.Marshal(m.ThirdPartyImages)
	if err != nil {
		return fmt.Errorf("marshal third_party_images: %w", err)
	}

	f, err := os.OpenFile(outputPath, os.O_APPEND|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open github output file %s: %w", outputPath, err)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil && err == nil {
			err = fmt.Errorf("close github output file %s: %w", outputPath, closeErr)
		}
	}()

	if _, err = fmt.Fprintf(f, "version=%s\n", m.Version); err != nil {
		return err
	}
	if _, err = fmt.Fprintf(f, "java_images=%s\n", string(javaJSON)); err != nil {
		return err
	}
	if _, err = fmt.Fprintf(f, "third_party_images=%s\n", string(thirdPartyJSON)); err != nil {
		return err
	}
	return nil
}
