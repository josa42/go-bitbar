package bitbar

// https://github.com/johnmccabe/openfaas-bitbar/blob/master/commands/install.go
// func getPluginDir() (string, error) {
// 	cmdName := "/usr/bin/defaults"
// 	cmdArgs := []string{"read", "com.matryer.BitBar", "pluginsDirectory"}

// 	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
// 	if err != nil {
// 		return "", fmt.Errorf("unable to determine pluginsDirectory: %v, %s", err, string(cmdOut))
// 	}

// 	dir := strings.TrimRight(string(cmdOut), "\n")

// 	if !dirExists(dir) {
// 		return "", fmt.Errorf("unable to check if dir exists: %v, %s", err, dir)
// 	}

// 	return dir, nil
// }

// open "bitbar://openPlugin?title=Cycle%20text%20and%20detail%20text&amp;src=https://github.com/matryer/bitbar-plugins/raw/master/Tutorial%2fcycle_text_and_detail.sh"
