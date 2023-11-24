module github.com/parsiya/semgrep_fun

go 1.19

require (
	github.com/parsiya/semgrep_go v0.0.0-20231105051654-c1d0fdd4be55
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

require (
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
)

// Change this for release.
replace github.com/parsiya/semgrep_go => ./semgrep_go
