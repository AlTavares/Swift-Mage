package mage

func SetVersion(version string) {
	Run("agvtool new-marketing-version", version)
}

func SetBuild(version string) {
	Run("agvtool new-version -all", version)
}

func IncrementBuildNumber() {
	Run("agvtool next-version -all")
}
