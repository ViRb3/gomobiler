// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"go/build"
	"path"
	"path/filepath"
	"fmt"
)

func goAndroidBuild(pkg *build.Package, androidArchs []string) (map[string]bool, error) {
	if ndkRoot == "" {
		return nil, errors.New("no Android NDK path is set. Please run gomobile init with the ndk-bundle installed through the Android SDK manager or with the -ndk flag set.")
	}
	appName := path.Base(pkg.ImportPath)
	libName := appName

	var libFiles []string
	nmpkgs := make(map[string]map[string]bool) // map: arch -> extractPkgs' output

	for _, arch := range androidArchs {
		env := androidEnv[arch]
		toolchain := ndk.Toolchain(arch)
		libPath := fmt.Sprintf("%s-android-%s", libName, toolchain.abi)
		libAbsPath := filepath.Join(buildO, libPath)
		if err := mkdir(filepath.Dir(libAbsPath)); err != nil {
			return nil, err
		}
		err := goBuild(
			pkg.ImportPath,
			env,
			"-buildmode=pie",
			"-o", libAbsPath,
		)
		if err != nil {
			return nil, err
		}
		nmpkgs[arch], err = extractPkgs(toolchain.Path("nm"), libAbsPath)
		if err != nil {
			return nil, err
		}
		libFiles = append(libFiles, libPath)
	}


	// TODO: return nmpkgs
	return nmpkgs[androidArchs[0]], nil
}