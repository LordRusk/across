package main

// there are 28 supported system types
const NUMSYS = 28

// supported operating systems
// with their support architectures.
var systems = []system{
	{name: "android", archs: []string{
		"arm",
	}},
	{name: "darwin", archs: []string{
		"amd64",
		"arm64",
	}},
	{name: "dragonfly", archs: []string{
		"amd64",
	}},
	{name: "freebsd", archs: []string{
		"386",
		"amd64",
		"arm",
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
	}},
	{name: "netbsd", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
	{name: "openbsd", archs: []string{
		"386",
		"amd64",
		"arm",
	}},
	{name: "plan9", archs: []string{
		"386",
		"amd64",
	}},
	{name: "solaris", archs: []string{
		"amd64",
	}},
	{name: "windows", archs: []string{
		"386",
		"amd64",
	}},
}
