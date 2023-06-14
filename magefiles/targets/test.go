package targets

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"

	"github.com/bmatcuk/doublestar"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Test mg.Namespace

func (Test) All() error {
	if _, err := os.Stat("testbin/bin"); os.IsNotExist(err) {
		mg.Deps(Test.Bin)
	}
	r, w := io.Pipe()
	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			line := scanner.Text()
			if !strings.Contains(line, "[no test files]") {
				fmt.Println(line)
			}
		}
	}()
	args := []string{"test",
		fmt.Sprintf("-p=%d", runtime.NumCPU()),
		"-race",
	}
	_, coverageEnabled := os.LookupEnv("DISABLE_COVERAGE")
	if coverageEnabled {
		args = append(args,
			"-cover",
			"-coverprofile=cover.out",
			"-coverpkg=./...",
		)
	}
	args = append(args, "./...")
	_, err := sh.Exec(map[string]string{
		"CGO_ENABLED": "1",
	}, w, os.Stderr, mg.GoCmd(), args...)
	if err != nil {
		return err
	}

	if !coverageEnabled {
		return nil
	}

	fmt.Print("processing coverage report... ")
	defer fmt.Println("done.")
	return filterCoverage("cover.out", []string{
		"**/*.pb.go",
		"**/*.pb*.go",
		"**/zz_*.go",
	})
}

func filterCoverage(report string, patterns []string) error {
	f, err := os.Open(report)
	if err != nil {
		return err
	}
	defer f.Close()

	tempFile := fmt.Sprintf(".%s.tmp", report)
	tf, err := os.Create(tempFile)
	if err != nil {
		return err
	}

	patternIndex := 0
	scan := bufio.NewScanner(f)
	scan.Scan() // mode line
	tf.WriteString(scan.Text() + "\n")
LINES:
	for scan.Scan() {
		line := scan.Text()
		filename, _, _ := strings.Cut(line, ":")
		var j int
		for i := patternIndex; j < len(patterns); i = (i + 1) % len(patterns) {
			match, _ := doublestar.Match(patterns[i], filename)
			if match {
				continue LINES
			}
			j++
		}
		tf.WriteString(line + "\n")
	}
	if err := scan.Err(); err != nil {
		return err
	}
	tf.Close()

	return os.Rename(tempFile, report)
}

func (Test) Env() {
	// check if testbin exists
	if _, err := os.Stat("testbin/bin"); os.IsNotExist(err) {
		mg.Deps(Test.Bin)
	}
	cmd := exec.Command("bin/testenv", "--agent-id-seed=0")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	proc := cmd.Process
	go func() {
		<-sigint
		proc.Signal(os.Interrupt)
	}()
	if err := cmd.Wait(); err != nil {
		panic(err)
	}
}

const k8sVersion = "1.26.3"

var testbinConfig = fmt.Sprintf(`
{
	"binaries": [
		{
			"name": "etcd",
			"sourceImage": "bitnami/etcd",
			"path": "/opt/bitnami/etcd/bin/etcd"
		},
		{
			"name": "prometheus",
			"sourceImage": "prom/prometheus",
			"path": "/bin/prometheus"
		},
		{
			"name": "promtool",
			"sourceImage": "prom/prometheus",
			"path": "/bin/promtool"
		},
		{
			"name": "node_exporter",
			"sourceImage": "prom/node-exporter",
			"path": "/bin/node_exporter"
		},
		{
			"name": "alertmanager",
			"sourceImage": "prom/alertmanager",
			"path": "/bin/alertmanager"
		},
		{
			"name": "amtool",
			"sourceImage": "prom/alertmanager",
			"path": "/bin/amtool"
		},
		{
			"name": "nats-server",
			"sourceImage": "nats",
			"path": "/nats-server"
		},
		{
			"name": "otelcol-custom",
			"sourceImage": "ghcr.io/rancher-sandbox/opni-otel-collector",
			"version": "v0.1.2-0.74.0",
			"path": "/otelcol-custom"
		},
		{
			"name": "kube-apiserver",
			"sourceImage": "registry.k8s.io/kube-apiserver",
			"version": "v%[1]s",
			"path": "/usr/local/bin/kube-apiserver"
		},
		{
			"name": "kube-controller-manager",
			"sourceImage": "registry.k8s.io/kube-controller-manager",
			"version": "v%[1]s",
			"path": "/usr/local/bin/kube-controller-manager"
		},
		{
			"name": "kubectl",
			"sourceImage": "bitnami/kubectl",
			"version": "%[1]s",
			"path": "/opt/bitnami/kubectl/bin/kubectl"
		}
	]
}`[1:], k8sVersion)

func (Test) BinConfig() {
	fmt.Println(testbinConfig)
}

func (Test) Bin() error {
	if _, err := os.Stat("testbin/bin"); err == nil {
		os.RemoveAll("testbin/bin")
	}
	return Dagger{}.run(daggerx, "testbin", testbinConfig)
}
