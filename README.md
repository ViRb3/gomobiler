# GoMobiler

A stripped-down modified `gomobile` to easily build go executables for Android.

Instead of generating an APK, `gomobiler` simply builds binaries for all applicable architectures and saves them in the format:
```
*app_name*-android-*architecture*
```

## Usage
1. `gomobiler init -ndk "/path/to/sdk/ndk-bundle/"`
2. Go to your project's directory
3. `gomobiler build`

You can specify an output directory using the `-o` flag