# Helm Update Version
This project gives an easy solution to bump up the helm chart version using simple binary.

## Usage
```
Usage:
  helm-update-version [flags]

Flags:
  -c, --changetype string   major/minor/hotfix
  -d, --directory string    Directory containing the Helm chart
  -h, --help                help for helm-update-version
      --verbose             detailed output
  -v, --version string      Specific version to update to

```

### when helm chart has standard version [major.minor.hotfix]
```
helm-update-version -d . -c major
```

### when you want your version to be placed in chart
```
helm-update-version -v <YOUR VERSION>
```

### when you want your version (standard format) to be placed in chart and updated
```
helm-update-version -v <YOUR VERSION> -c major
```