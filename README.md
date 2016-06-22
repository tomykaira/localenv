# Localenv

Per-directory environment variables manager.

Feature

- Securitiy: env vars are encrypted using GPG. You can credentials like API keys in piece.
- Portability: Runs on most of POSIX OSes. Validated on Fedura and Mac OS X.

We owe much to [trousseau](https://github.com/oleiade/trousseau) to achieve these feature.

## Setup

For now, localenv cannot run by itself.  Trousseau is required to create secure store file.

```
gpg --gen-key
# Create key pair for <foo@bar.com>
export TROUSSEAU_MASTER_GPG_ID=foo@bar.com
trousseau create foo@bar.com
```

Use keyring:

- For Mac OS, create a password entry with following information
    - Keychain item name: localenv-gpg-password (or anything you like)
    - Account name: your login account name, that is, `$USER`
    - Password: passphrase for gpg key created above

```
export TROUSSEAU_KEYRING_SERVICE=localenv-gpg-password
```

Use gpg-agent:

CAUTION: trousseau has a bug in GPG_AGENT_INFO envvar name.
Check [my PR](https://github.com/oleiade/trousseau/pull/169/files).

```
eval $(gpg-agent --daemon)
```

At last, hook `cd` to update env.

```
echo "source ${PWD}/loader.bash" >> ~/.bashrc
```

## Usage

```
localenv set foo bar
localenv get foo # => bar
```
