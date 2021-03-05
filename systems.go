package main

// supported operating systems
// with their support architectures.
var systems = []system{
	{name: "android", archs: []string{
		"arm",
	}},
	{name: "darwin", archs: []string{
		"amd64",
	}},
	{name: "dragonfly", archs: []string{
		"amd64",
	}},
	{name: "freebsd", archs: []string{
		"386",
		"amd64",
		"arm",
		"arm64",
	}},
	{name: "ios", archs: []string{
		"amd64",
		"arm64",
	}},
	{name: "illumos", archs: []string{
		"amd64",
	}},
	{name: "js", archs: []string{
		"wasm",
	}},
	{name: "linux", archs: []string{
		"386",
		"amd64",
		"arm",
		"arm64",
		"ppc64",
		"ppc64le",
		"mips",
		"mips64",
		"mips64le",
		"mipsle",
		"riscv64",
		"s390x",
	}},
	{name: "netbsd", archs: []string{
		"386",
		"amd64",
		"arm",
		"arm64",
	}},
	{name: "openbsd", archs: []string{
		"386",
		"amd64",
		"arm",
		"arm64",
		"mips64",
	}},
	{name: "plan9", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
	{name: "solaris", archs: []string{
		"amd64",
	}},
	{name: "windows", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
}
