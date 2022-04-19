module github.com/notional-labs/tinyport

go 1.16

require (
	github.com/99designs/keyring v1.1.6
	github.com/AlecAivazis/survey/v2 v2.1.1
	github.com/blang/semver v3.5.1+incompatible
	github.com/briandowns/spinner v1.11.1
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/charmbracelet/glow v1.4.0
	github.com/containerd/containerd v1.5.8 // indirect
	github.com/cosmos/cosmos-sdk v0.45.3
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/ibc-go/v2 v2.0.3
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/docker/docker v20.10.7+incompatible
	github.com/emicklei/proto v1.9.0
	github.com/fatih/color v1.13.0
	github.com/ghodss/yaml v1.0.0
	github.com/go-git/go-git/v5 v5.1.0
	github.com/gobuffalo/envy v1.9.0 // indirect
	github.com/gobuffalo/genny v0.6.0
	github.com/gobuffalo/logger v1.0.3
	github.com/gobuffalo/packd v0.3.0
	github.com/gobuffalo/plush v3.8.3+incompatible
	github.com/gobuffalo/plushgen v0.1.2
	github.com/goccy/go-yaml v1.9.4
	github.com/gogo/protobuf v1.3.3
	github.com/google/go-cmp v0.5.7 // indirect
	github.com/google/go-github/v37 v37.0.0
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gookit/color v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/rpc v1.2.0
	github.com/iancoleman/strcase v0.2.0
	github.com/imdario/mergo v0.3.12
	github.com/jpillora/chisel v1.7.7
	github.com/mattn/go-zglob v0.0.3
	github.com/moby/sys/mount v0.3.0 // indirect
	github.com/otiai10/copy v1.6.0
	github.com/pelletier/go-toml v1.9.4
	github.com/pkg/errors v0.9.1
	github.com/radovskyb/watcher v1.0.7
	github.com/rogpeppe/go-internal v1.6.2 // indirect
	github.com/rs/cors v1.8.2
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/stretchr/testify v1.7.1
	github.com/takuoki/gocase v1.0.0
	github.com/tendermint/flutter/v2 v2.0.3
	github.com/tendermint/tendermint v0.34.19
	github.com/tendermint/tm-db v0.6.6
	github.com/tendermint/vue v0.3.0
	golang.org/x/mod v0.5.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211
	google.golang.org/genproto v0.0.0-20220302033224-9aa15565e42a // indirect
	google.golang.org/grpc v1.45.0
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
