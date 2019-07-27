os=$(go env GOOS)
arch=$(go env GOARCH)

echo "Found: os: $os arch: $arch"

# download kubebuilder and extract it to tmp
curl -sL https://go.kubebuilder.io/dl/2.0.0-beta.0/${os}/${arch} | tar -xz -C /tmp/

# move to a long-term location and put it on your path
# (you'll need to set the KUBEBUILDER_ASSETS env var if you put it somewhere else)
dst=~/dev/tools/k8s/kubebuilder
mkdir -p "$dst"

kube_bin="kubebuilder_2.0.0-beta.0_${os}_${arch}"
mv "/tmp/$kube_bin" "$dst/"
cd "$dst"
rm -f latest
ln -s "$kube_bin" latest
