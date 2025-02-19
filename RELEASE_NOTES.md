<!-- markdownlint-disable MD013 -->
<!-- markdownlint-disable MD024 -->
<!-- markdownlint-disable MD040 -->

# Release Notes

Release Procedure is defined in the [CONTRIBUTING](CONTRIBUTING.md#release-procedure) document.

## v6.0.0

Highlights:

- TODO: special assets, new Gov messages
- New [inflation mechanism](./docs/design_docs/012-umee-inflation-v2.md).

### New Inflation Mechanism

The Upgrade Handler sets the following values to the Umee `x/ugov` Inflation Cycle parameters:

- `max_supply = 21e18uumee` (21 billions UMEE)
- `inflation_cycle = time.Hour * 24 * 365 * 2` (2 years)
- `inflation_reduction_rate = 2500 basis points` (25%)

The new Inflation Cycle will start on 2023-10-15 15:00 UTC. This will mark the first inflation reduction from the current rates:

- `inflation_min` 7% → 5.25%
- `inflation_max` 14% → 10.5%

The x/staking Bonded Goal stays the same: 33.00%.

### Validators

#### libwasmvm update

Our dependencies have been updated. Now the binary requires `libwasmvm v1.3.0`. When you build the binary from source on the server machine you probably don't need any change. However when you download a binary from GitHub, or from other source, make sure you update the `/usr/lib/libwasmvm.<cpu_arch>.so`. For example:

- copy from `$GOPATH/pkg/mod/github.com/!cosm!wasm/wasmvm@v1.3.0/internal/api/libwasmvm.$(uname -m).so`
- or download from github `wget https://raw.githubusercontent.com/CosmWasm/wasmvm/v1.3.0/internal/api/libwasmvm.$(uname -m).so -O /lib/libwasmvm.$(uname -m).so`

You don't need to do anything if you are using our Docker image.

#### Min Gas Prices

Since v4.2 release we request all validators set a `minimum-gas-prices` setting (in app `config/app.toml` file, general settings). We recommend `0.1uumee` which is equal the current Keplr _average_ setting:

```
minimum-gas-prices = "0.1uumee"
```

### Upgrade instructions

- Download latest binary or build from source.
- Make sure `libwasmvm.$(uname -m).so` is properly linked
  - Run the binary to make sure it works for you: `umeed version`
- Wait for software upgrade proposal to pass and trigger the chain upgrade.
- Swap binaries.
- Ensure latest Price Feeder (see [compatibility matrix](https://github.com/umee-network/umee/#release-compatibility-matrix)) is running and check your price feeder config is up to date.
- Restart the chain.

You can use Cosmovisor → see [instructions](https://github.com/umee-network/umee/#cosmovisor).

#### Docker

Docker images are available in [ghcr.io umee-network](https://github.com/umee-network/umee/pkgs/container/umeed) repository.
