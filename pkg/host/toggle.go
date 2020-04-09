package host

// Toggle alternates between enable and disable status of a profile.

func Toggle(dst, profile string) error {
	h, err := getHostData(dst, profile)
	if err != nil {
		return err
	}

	status := getProfileStatus(h, profile)

	switch status {
	case Enabled:
		disableProfile(h, profile)
	case Disabled:
		enableProfile(h, profile)
	default:
		return UnknownProfileError
	}

	return writeHostData(dst, h)
}

func getProfileStatus(h *hostFile, profile string) ProfileStatus {
	pData, ok := h.profiles[profile]
	if !ok {
		return ""
	}

	for _, l := range pData {
		if !IsHostLine(l) {
			continue
		}
		if IsDisabled(pData[0]) {
			return Disabled
		}
		return Enabled
	}

	return ""
}