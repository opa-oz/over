# over
Tired of multiple files with your package version inside? It's over!

![version](https://img.shields.io/badge/version-0.1.0-blue)

## Motivation

Sometimes you put too many utilities and config files to your repository. And many of them require you to specify version of your package.

Usually you would rely on CI/CD like this:
- Push Git `tag` with version
- Get it as environment variable
- `sed` inside all files
- Release

What if you can update version across all the places at one? Try `over`

## Example

To start using `over` you just need to add one config file (`.over.yaml`) to your project. Let's see `over`'s config:
```yaml
package:
  name: over
  version: "0.1.0"
  files:
    - name: README.md
      templates:
        - https://img.shields.io/badge/version-__VERSION__-blue
        - 'version: "__VERSION__"'
    - name: cmd/version.go
      templates:
        - fmt.Println("__VERSION__")
```

- `name` is well self-explanatory
- `version` is the only source of truth for your project now
- `files` is the most interesting part. In here you specify all files and patterns which should be replaced with new version.

`files` array contains objects:
- `name`: filename (relative to project's root)
- `templates`: array of lines, containing version

You may find sort of "inception" now. `over`'s config is updating package's version inside README.md (twice).
Part:
```yaml
package:
  name: over
  version: "0.1.0" # <=== It's text in README.md
  ...
```

And:
```yaml
...
- name: README.md
  templates:
    - https://img.shields.io/badge/version-__VERSION__-blue
    - version: "__VERSION__" # <=== This template says "Replace version that looks like this in README.md"
```

**Important:** 
To avoid non-valid files, `over` doesn't require you to put `__VERSION__` to each and every file. 
When you are setting up project the first time - just put the same version everywhere (for example `0.0.1`). 
But when you configure `.over.yaml` - replace this version with `__VERSION__` in a template. 
Therefore, you will get updated version everywhere without need to compromise on your code check


## Usage
[CLI-docs](docs/over.md)

1. Get current version (also available as [Github action](https://github.com/yakubique/over)):
```bash
over get
```

2. Get current version of nested package (mono-repository friendly):
```bash
over --config nested-package/.over.yaml 
```

3. Bump patch (`v0.0.1 => v0.0.2`)
```bash
over up --patch --inplace
```
**Note:** Without `--inspace` it'll just post next version

4. Bump minor (`v0.1.5 => v0.2.0`)
```bash
over up -i --minor
```

5. Bump major (`v0.2.6 => v1.0.0`)
```bash
over up -i --major 
```

6. Shortcuts also available
```bash
# Eq: --patch --inplace
over up -pi

# Eq: --minor --inplace
over up -mi 

# Eq: --major --inplace
over up -Mi
```
