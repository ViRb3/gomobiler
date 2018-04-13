// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	javacTargetVer = "1.7"
	minAndroidAPI  = 15
)


// androidAPIPath returns an android SDK platform directory under ANDROID_HOME.
// If there are multiple platforms that satisfy the minimum version requirement
// androidAPIPath returns the latest one among them.
func androidAPIPath() (string, error) {
	sdk := os.Getenv("ANDROID_HOME")
	if sdk == "" {
		return "", fmt.Errorf("ANDROID_HOME environment var is not set")
	}
	sdkDir, err := os.Open(filepath.Join(sdk, "platforms"))
	if err != nil {
		return "", fmt.Errorf("failed to find android SDK platform: %v", err)
	}
	defer sdkDir.Close()
	fis, err := sdkDir.Readdir(-1)
	if err != nil {
		return "", fmt.Errorf("failed to find android SDK platform (min API level: %d): %v", minAndroidAPI, err)
	}

	var apiPath string
	var apiVer int
	for _, fi := range fis {
		name := fi.Name()
		if !fi.IsDir() || !strings.HasPrefix(name, "android-") {
			continue
		}
		n, err := strconv.Atoi(name[len("android-"):])
		if err != nil || n < minAndroidAPI {
			continue
		}
		p := filepath.Join(sdkDir.Name(), name)
		_, err = os.Stat(filepath.Join(p, "android.jar"))
		if err == nil && apiVer < n {
			apiPath = p
			apiVer = n
		}
	}
	if apiVer == 0 {
		return "", fmt.Errorf("failed to find android SDK platform (min API level: %d) in %s",
			minAndroidAPI, sdkDir.Name())
	}
	return apiPath, nil
}
