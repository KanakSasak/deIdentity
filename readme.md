# deidentity
**deidentity** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Additionally, Ignite CLI offers both Vue and React options for frontend scaffolding:

For a Vue frontend, use: `ignite scaffold vue`
For a React frontend, use: `ignite scaffold react`
These commands can be run within your scaffolded blockchain project. 


For more information see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/kanaksasak/deIdentity@latest! | sudo bash
```
`username/deIdentity` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

### Query

Register Identity
```
deIdentityd tx deidentity register-identity "Alice Wonderland" "07/04/1865" NationIDNumber123 --from alice --chain-id deIdentity
```

Approve Identity
```
deIdentityd tx deidentity approve-identity "Robot bro" "07/04/1865" NationIDNumber123 --from zack --chain-id deIdentity 
```

Get List Identity
```
deIdentityd query deidentity list-identity
```

### Output

```
┌──(kanaksasak㉿KanakSasak)-[~/GolandProjects/deIdentity]
└─$ deIdentityd query deidentity list-identity                                                                                      
Identity:
- birthdate: 07/04/1865
  creator: cosmos19j5hr3e3pgpe6ffwsk7ax7jkq7q8dc5jdsdt8g
  name: Alice Wonderland
  nationalId: 8fff406d619482a5e9984cbb85e747d1eeff3057f25c61d23045b161c761e871
- birthdate: 07/04/1865
  creator: cosmos19dzmndpmn9p6zje85k9zwavlwskugwgjxmctvl
  id: "1"
  name: Donald
  nationalId: 8fff406d619482a5e9984cbb85e747d1eeff3057f25c61d23045b161c761e871
- approve: cosmos19dzmndpmn9p6zje85k9zwavlwskugwgjxmctvl
  birthdate: 07/04/1865
  creator: cosmos19j5hr3e3pgpe6ffwsk7ax7jkq7q8dc5jdsdt8g
  id: "2"
  name: Robot bro
  nationalId: 8fff406d619482a5e9984cbb85e747d1eeff3057f25c61d23045b161c761e871
  verified: true
pagination:
  total: "3"

```

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
