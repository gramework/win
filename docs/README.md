# gramework.win

- [Directory structure](#directory-structure)
	- [`/` (root)](#-root)
	- [`{domain root}/{project shortcut}/v{version number}/`](#domain-rootproject-shortcutvversion-number)
	- [`{domain root}/{project shortcut}/frontend/`](#domain-rootproject-shortcutfrontend)
	- [`/build`](#build)
	- [`/ops`](#ops)

## Directory structure

### `/` (root)

Project root.

**Data rules**: any private keys, services' passwords or other sensetive data **MUST NEVER PUBLISHED HERE**.

**Allowed root-level files:**
- Links to files project's subfolder (like Vagrantfile -> /ops/_Vagrantfile)
- .git* files and .github dir
- Files, if documented below

**Links:**
- Vagrantfile -> /ops/_Vagrantfile, ansible.cfg -> /ops/ansible.cfg: deployment done from project root, not `/ops` folder.
  Also, vagrant by default copying only current working directory.
- CONTRIGUTING.md -> /docs/CONTRIBUTING.md: used by github

**Files:**
- `run_*.g{o,l}`: project runners
- `sf.go`: `/securityFeedback` register
- `win.go`: `/win` register
- `share.go`: share register
- `init.go`: initializer
- `README.md`: project's welcome/documentation part
- `LICENSE`: license information
- `win`: `go build` compiled binary default path, `.gitignore`d
- `win.img`: built bare-metal image, `.gitignore`d
- `.infrastructure`: infrastructure information, if compiled with goodlang, `.gitignore`d
- `profile.out`, `cpu.out`, `mem.out`: profiler output
- `tls-gramecache.dev`, `tls-gramecache`: gramework tls cache


### `{domain root}/{project shortcut}/v{version number}/`

Domain's server-side code.

**Directory naming rules:** directories should be called exactly the same as microservice that will be stored in it, except:
- `/shared`: shared microservices dependencies
- `/sharedMicroservices`: shared microservices

**Examples:**

Directory path for `infrastructure` microservice from project root should be `/shared/infrastructure/v{version number}`, because we want to host shared
infrastructure service instead of hosting 1 microservice * 3 default replication * domains count.

Directory path for `sfRegistry` microservice from project root should be `/{domain root}/`

### `{domain root}/{project shortcut}/frontend/`

Domain's frontends' code.

**Directory naming rules:** directories should be called exactly the same as frontend in it.

**Examples:**

- Directory path for public frontend from project root should be `/{domain root}/{project shortcut}/frontend/public`
- Directory path for backoffice frontend from project root should be `/{domain root}/{project shortcut}/frontend/backoffice`

### `/build`

Built project files

**Directory naming rules:**  directories should be called exactly the same as the directory with sources, but with `/build/` prefix.

**Examples:**

- Directory path for built public frontend from project root should be `/build/{domain root}/{project shortcut}/frontend/public`
- Directory path for built backoffice frontend from project root should be `/build/{domain root}/{project shortcut}/frontend/public`
- Directory path for `infrastructure` microservice from project root should be `/build/shared/infrastructure/v{version number}`

### `/ops`

Deployment automatisation files.

**Data rules**: any private keys, services' passwords or other sensetive data **MUST NEVER PUBLISHED HERE**.

**Directory naming rules:**: directory structure here is a standard ansible project:
- `/roles` - machine roles
- `/group_vars` - group variables
- `/host_vars` - host-level variables
- `/vagrant_inventory` - ansible config for vagrant deployment
- `/Vagrantfile` - vagrant config
- `/vagrant.sh` - vagrant provision
- `/ansible.cfg` - ansible config
