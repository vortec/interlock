package version

var (
	Version = "1.0.1"

	// Build will be overwritten automatically by the build system
	Build = "-dev"

	// GitCommit will be overwritten automatically by the build system
	GitCommit = "HEAD"
)

func FullVersion() string {
	return Version + Build + " (" + GitCommit + ")"
}
