Product Crawler Server
- instal gvm
- gvm pkgset create "name"
- gvm pkgset use metis-crawler
- do:

export GVM_ROOT; GVM_ROOT="/home/rudy/.gvm"
export gvm_go_name; gvm_go_name="go1.8.3"
export gvm_pkgset_name; gvm_pkgset_name="metis-crawler"
export GOROOT; GOROOT="$GVM_ROOT/gos/go1.8.3"
export GOPATH; GOPATH="$GVM_ROOT/pkgsets/go1.8.3/metis-crawler"
export GVM_OVERLAY_PREFIX; GVM_OVERLAY_PREFIX="${GVM_ROOT}/pkgsets/go1.8.3/metis-crawler/overlay"
export PATH;
PATH="${GVM_ROOT}/pkgsets/go1.8.3/metis-crawler/bin:${GVM_ROOT}/gos/go1.8.3/bin:${GVM_OVERLAY_PREFIX}/bin:${GVM_ROOT}/bin:${PATH}"
export LD_LIBRARY_PATH; LD_LIBRARY_PATH="${GVM_OVERLAY_PREFIX}/lib:${LD_LIBRARY_PATH}"
export DYLD_LIBRARY_PATH; DYLD_LIBRARY_PATH="${GVM_OVERLAY_PREFIX}/lib:${DYLD_LIBRARY_PATH}"
export PKG_CONFIG_PATH; PKG_CONFIG_PATH="${GVM_OVERLAY_PREFIX}/lib/pkgconfig:${PKG_CONFIG_PATH}"
export gvm_pkgset_name="metis-crawler"
export GOPATH; GOPATH="$HOME/workspaces/go/metis-crawler:$GOPATH"
export PATH; PATH="/home/rudy/.gvm/pkgsets/go1.8.3/metis-crawler/bin:$PATH"
# Package Set-Specific Overrides
export GVM_OVERLAY_PREFIX; GVM_OVERLAY_PREFIX="${GVM_ROOT}/pkgsets/go1.8.3/metis-crawler/overlay"
export PATH; PATH="/home/rudy/.gvm/pkgsets/go1.8.3/metis-crawler/bin:${GVM_OVERLAY_PREFIX}/bin:$HOME/workspaces/go/metis-crawler/bin:${PATH}"
export LD_LIBRARY_PATH; LD_LIBRARY_PATH="${GVM_OVERLAY_PREFIX}/lib:${LD_LIBRARY_PATH}"
export DYLD_LIBRARY_PATH; DYLD_LIBRARY_PATH="${GVM_OVERLAY_PREFIX}/lib:${DYLD_LIBRARY_PATH}"
export PKG_CONFIG_PATH; PKG_CONFIG_PATH="${GVM_OVERLAY_PREFIX}/lib/pkgconfig:${PKG_CONFIG_PATH}"
