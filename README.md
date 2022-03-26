# Helm Update Version
This project gives an easy solution to bump up the helm chart version using simple binary.

## Runtime parameters
Binary accepts chart directory and changetype (major/minor/hotfix) as an user input.
By default changetype is none and it will not update the chart version.

- ( -d ) : helm chart directory
- ( -b ) : major/minor/hostfix. default is nil

## Usage

### Read chart version
```
helm-update-version -d <HELM-CHART-DIR>
```

### Update chart version
```
helm-update-version -d <HELM-CHART-DIR> -b <major/minor/hotfix>
```