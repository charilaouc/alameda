#!/usr/bin/env bash

set -eu

function annotate_crd() {
  script='/^metadata:/a\
\ \ annotations:\
\ \ \ \ exclude.release.openshift.io/internal-openshift-hosted: "true"'
  input="${1}"
  output="${2}"
  sed -e "${script}" "${input}" > "${output}"
}

echo "Building controller-gen tool..."
go build -o bin/controller-gen github.com/openshift/machine-api-operator/vendor/sigs.k8s.io/controller-tools/cmd/controller-gen

dir=$(mktemp -d -t XXXXXXXX)
echo $dir
mkdir -p $dir/src/github.com/openshift/machine-api-operator/pkg/apis

cp -r pkg/apis/machine $dir/src/github.com/openshift/machine-api-operator/pkg/apis
# Some dependencies need to be copied as well. Othwerwise, controller-gen will complain about non-existing kind Unsupported
cp -r vendor $dir/src/github.com/openshift/machine-api-operator/
cp go.mod go.sum $dir/src/github.com/openshift/machine-api-operator/

cwd=$(pwd)
pushd $dir/src/github.com/openshift/machine-api-operator
GOPATH=$dir ${cwd}/bin/controller-gen crd \
    paths=$dir/src/github.com/openshift/machine-api-operator/pkg/apis/... \
    output:crd:dir=$dir/src/github.com/openshift/machine-api-operator/config/crds/

#${cwd}/bin/controller-gen crd paths=$dir/src/github.com/openshift/machine-api-operator/pkg/apis/... output:crd:dir=$dir/src/github.com/openshift/machine-api-operator/config/crds/
popd

echo "Copying and patching generated CRDs"
annotate_crd $dir/src/github.com/openshift/machine-api-operator/config/crds/machine.openshift.io_machinehealthchecks.yaml install/0000_30_machine-api-operator_07_machinehealthcheck.crd.yaml
annotate_crd $dir/src/github.com/openshift/machine-api-operator/config/crds/machine.openshift.io_machinesets.yaml install/0000_30_machine-api-operator_03_machineset.crd.yaml
annotate_crd $dir/src/github.com/openshift/machine-api-operator/config/crds/machine.openshift.io_machines.yaml install/0000_30_machine-api-operator_02_machine.crd.yaml

rm -rf $dir
