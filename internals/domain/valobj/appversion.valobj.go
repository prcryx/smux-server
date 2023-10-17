package valobj

type AppVersion string

func GetAppVersion(version string) AppVersion {
	return AppVersion(version)
}

func (appVersion AppVersion) ToString() string {
	return string(appVersion)
}
